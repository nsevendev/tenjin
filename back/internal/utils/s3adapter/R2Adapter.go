package s3adapter

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
)

var adapterCloudflareR2 *Adapter

type Adapter struct {
	accountID  string
	accessKey  string
	secretKey  string
	bucketName string
	keyPrefix  string
	httpClient *http.Client
	baseURL    string
}

// Options permet de configurer l'Adapter R2 sans dépendre d'ENV dans les tests.
type options struct {
	AccountID string
	AccessKey string
	SecretKey string
	Bucket    string
	KeyPrefix string
}

// AdapterCloudflareR2 retourne l'adapteur Cloudflare R2
// Il est initialisé par CreateAdapteur() qui charge les options depuis l'environnement.
func AdapterCloudflareR2() *Adapter {
	return adapterCloudflareR2
}

// CreateAdapteur crée un Adapter Cloudflare R2 avec HTTP direct
func CreateAdapteur() {
	opts := &options{
		AccountID: env.Get("R2_ACCOUNT_ID"),
		AccessKey: env.Get("R2_ACCESS_KEY_ID"),
		SecretKey: env.Get("R2_SECRET_ACCESS_KEY"),
		Bucket:    env.Get("R2_BUCKET_NAME"),
		KeyPrefix: env.Get("R2_KEY_PREFIX"),
	}

	baseURL := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", opts.AccountID)

	logger.Sf("R2: Chargement config pour %s", opts.AccountID)
	logger.Sf("R2: Chargement bucket %s", opts.Bucket)

	adapterCloudflareR2 = &Adapter{
		accountID:  opts.AccountID,
		accessKey:  opts.AccessKey,
		secretKey:  opts.SecretKey,
		bucketName: opts.Bucket,
		keyPrefix:  opts.KeyPrefix,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    baseURL,
	}
}

// Upload envoie des octets à la clé donnée via HTTP direct
func (a *Adapter) Upload(ctx context.Context, key string, data []byte) error {
	finalKey := a.prefixed(key)

	fmt.Printf("DEBUG R2 Adapter HTTP:\n")
	fmt.Printf("  - key reçue: %q\n", key)
	fmt.Printf("  - keyPrefix: %q\n", a.keyPrefix)
	fmt.Printf("  - key finale: %q\n", finalKey)
	fmt.Printf("  - bucket: %q\n", a.bucketName)
	fmt.Printf("  - baseURL: %q\n", a.baseURL)

	// Construction de l'URL
	objectURL := fmt.Sprintf("%s/%s/%s", a.baseURL, a.bucketName, finalKey)

	// Création de la requête
	req, err := http.NewRequestWithContext(ctx, "PUT", objectURL, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("r2: création requête: %w", err)
	}

	// Headers de base
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(data)))
	req.Header.Set("Host", fmt.Sprintf("%s.r2.cloudflarestorage.com", a.accountID))

	// Signature AWS v4
	if err := a.signRequest(req, data); err != nil {
		return fmt.Errorf("r2: signature requête: %w", err)
	}

	fmt.Printf("  - URL finale: %q\n", objectURL)
	fmt.Printf("  - Headers: %v\n", req.Header)

	// Exécution de la requête
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("r2: requête HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("r2: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Download lit le contenu stocké à la clé donnée via HTTP direct
func (a *Adapter) Download(ctx context.Context, key string) ([]byte, error) {
	finalKey := a.prefixed(key)
	objectURL := fmt.Sprintf("%s/%s/%s", a.baseURL, a.bucketName, finalKey)

	req, err := http.NewRequestWithContext(ctx, "GET", objectURL, nil)
	if err != nil {
		return nil, fmt.Errorf("r2: création requête: %w", err)
	}

	req.Header.Set("Host", fmt.Sprintf("%s.r2.cloudflarestorage.com", a.accountID))

	if err := a.signRequest(req, nil); err != nil {
		return nil, fmt.Errorf("r2: signature requête: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("r2: requête HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("r2: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

// Delete supprime l'objet à la clé donnée via HTTP direct
func (a *Adapter) Delete(ctx context.Context, key string) error {
	finalKey := a.prefixed(key)
	objectURL := fmt.Sprintf("%s/%s/%s", a.baseURL, a.bucketName, finalKey)

	req, err := http.NewRequestWithContext(ctx, "DELETE", objectURL, nil)
	if err != nil {
		return fmt.Errorf("r2: création requête: %w", err)
	}

	req.Header.Set("Host", fmt.Sprintf("%s.r2.cloudflarestorage.com", a.accountID))

	if err := a.signRequest(req, nil); err != nil {
		return fmt.Errorf("r2: signature requête: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("r2: requête HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("r2: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// prefixed ajoute le préfixe de clé si défini, sinon retourne la clé d'origine.
func (a *Adapter) prefixed(key string) string {
	if a.keyPrefix == "" {
		return key
	}
	return a.keyPrefix + key
}

// signRequest signe la requête avec AWS Signature v4
func (a *Adapter) signRequest(req *http.Request, payload []byte) error {
	now := time.Now().UTC()
	dateStamp := now.Format("20060102")
	amzDate := now.Format("20060102T150405Z")

	// Headers requis
	req.Header.Set("X-Amz-Date", amzDate)

	// Payload hash
	var payloadHash string
	if payload != nil {
		h := sha256.Sum256(payload)
		payloadHash = hex.EncodeToString(h[:])
	} else {
		h := sha256.Sum256([]byte(""))
		payloadHash = hex.EncodeToString(h[:])
	}
	req.Header.Set("X-Amz-Content-Sha256", payloadHash)

	// Canonical request
	canonicalURI := req.URL.Path
	if canonicalURI == "" {
		canonicalURI = "/"
	}

	canonicalQueryString := req.URL.RawQuery
	canonicalHeaders := a.getCanonicalHeaders(req)
	signedHeaders := a.getSignedHeaders(req)

	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		req.Method,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		payloadHash,
	)

	// String to sign
	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := fmt.Sprintf("%s/auto/s3/aws4_request", dateStamp)
	hasher := sha256.Sum256([]byte(canonicalRequest))
	hashedCanonicalRequest := hex.EncodeToString(hasher[:])

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s",
		algorithm,
		amzDate,
		credentialScope,
		hashedCanonicalRequest,
	)

	// Signature
	signature := a.calculateSignature(stringToSign, dateStamp)

	// Authorization header
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		a.accessKey,
		credentialScope,
		signedHeaders,
		signature,
	)

	req.Header.Set("Authorization", authorization)
	return nil
}

func (a *Adapter) getCanonicalHeaders(req *http.Request) string {
	var canonicalHeaders strings.Builder

	// Créer une map avec les noms de headers en minuscules
	headerMap := make(map[string][]string)
	for name, values := range req.Header {
		headerMap[strings.ToLower(name)] = values
	}

	// Trier les noms de headers
	var headerNames []string
	for name := range headerMap {
		headerNames = append(headerNames, name)
	}
	sort.Strings(headerNames)

	// Construire les headers canoniques
	for _, name := range headerNames {
		values := headerMap[name]
		if len(values) > 0 {
			value := strings.Join(values, ",")
			canonicalHeaders.WriteString(fmt.Sprintf("%s:%s\n", name, strings.TrimSpace(value)))
		}
	}
	return canonicalHeaders.String()
}

func (a *Adapter) getSignedHeaders(req *http.Request) string {
	var headerNames []string
	for name := range req.Header {
		headerNames = append(headerNames, strings.ToLower(name))
	}
	sort.Strings(headerNames)
	return strings.Join(headerNames, ";")
}

func (a *Adapter) calculateSignature(stringToSign, dateStamp string) string {
	kDate := a.hmacSHA256([]byte("AWS4"+a.secretKey), dateStamp)
	kRegion := a.hmacSHA256(kDate, "auto")
	kService := a.hmacSHA256(kRegion, "s3")
	kSigning := a.hmacSHA256(kService, "aws4_request")
	signature := a.hmacSHA256(kSigning, stringToSign)
	return hex.EncodeToString(signature)
}

func (a *Adapter) hmacSHA256(key []byte, data string) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}
