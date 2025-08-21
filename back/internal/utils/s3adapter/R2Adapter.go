package s3adapter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
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
