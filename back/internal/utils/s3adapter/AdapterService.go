package s3adapter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

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
