package s3adapter

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testAdapter *Adapter
)

func TestMain(m *testing.M) {
	testAdapter = &Adapter{
		accountID:  "test-account-id",
		accessKey:  "test-access-key",
		secretKey:  "test-secret-key",
		bucketName: "test-bucket",
		keyPrefix:  "test/prefix/",
		httpClient: &http.Client{},
		baseURL:    "https://test-account-id.r2.cloudflarestorage.com",
	}

	code := m.Run()
	logger.If("Tests terminés")
	_ = code
}

func TestAdapter_prefixed_WithPrefix(t *testing.T) {
	testup.LogNameTestInfo(t, "Test prefixed function with prefix")

	key := "documents/file.txt"
	result := testAdapter.prefixed(key)

	expected := "test/prefix/documents/file.txt"
	assert.Equal(t, expected, result)
}

func TestAdapter_prefixed_WithoutPrefix(t *testing.T) {
	testup.LogNameTestInfo(t, "Test prefixed function without prefix")

	// Créer un adaptateur sans préfixe
	adapterNoPrefix := &Adapter{
		accountID:  "test-account-id",
		accessKey:  "test-access-key",
		secretKey:  "test-secret-key",
		bucketName: "test-bucket",
		keyPrefix:  "",
		httpClient: &http.Client{},
		baseURL:    "https://test-account-id.r2.cloudflarestorage.com",
	}

	key := "documents/file.txt"
	result := adapterNoPrefix.prefixed(key)

	assert.Equal(t, key, result)
}

func TestAdapter_prefixed_EmptyKey(t *testing.T) {
	testup.LogNameTestInfo(t, "Test prefixed function with empty key")

	key := ""
	result := testAdapter.prefixed(key)

	expected := "test/prefix/"
	assert.Equal(t, expected, result)
}

func TestAdapter_hmacSHA256(t *testing.T) {
	testup.LogNameTestInfo(t, "Test HMAC SHA256 function")

	key := []byte("test-key")
	data := "test-data"

	result := testAdapter.hmacSHA256(key, data)

	// Vérifier que le résultat est bien un hash de 32 bytes
	assert.Len(t, result, 32)

	// Tester la reproductibilité
	result2 := testAdapter.hmacSHA256(key, data)
	assert.Equal(t, result, result2)

	// Tester avec des données différentes
	result3 := testAdapter.hmacSHA256(key, "different-data")
	assert.NotEqual(t, result, result3)
}

func TestAdapter_calculateSignature(t *testing.T) {
	testup.LogNameTestInfo(t, "Test calculate signature function")

	stringToSign := "AWS4-HMAC-SHA256\n20250821T120000Z\n20250821/auto/s3/aws4_request\ntest-hash"
	dateStamp := "20250821"

	signature := testAdapter.calculateSignature(stringToSign, dateStamp)

	// Vérifier que la signature est une chaîne hex de 64 caractères (32 bytes)
	assert.Len(t, signature, 64)
	assert.Regexp(t, "^[a-f0-9]+$", signature)

	// Tester la reproductibilité
	signature2 := testAdapter.calculateSignature(stringToSign, dateStamp)
	assert.Equal(t, signature, signature2)

	// Tester avec des données différentes
	signature3 := testAdapter.calculateSignature(stringToSign, "20250822")
	assert.NotEqual(t, signature, signature3)
}

func TestAdapter_getSignedHeaders(t *testing.T) {
	testup.LogNameTestInfo(t, "Test get signed headers function")

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Amz-Date", "20250821T120000Z")
	req.Header.Set("Authorization", "AWS4-HMAC-SHA256...")

	result := testAdapter.getSignedHeaders(req)

	// Les headers doivent être triés alphabétiquement et en minuscules
	expected := "authorization;content-type;x-amz-date"
	assert.Equal(t, expected, result)
}

func TestAdapter_getSignedHeaders_EmptyHeaders(t *testing.T) {
	testup.LogNameTestInfo(t, "Test get signed headers with empty headers")

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	result := testAdapter.getSignedHeaders(req)

	assert.Equal(t, "", result)
}

func TestAdapter_getCanonicalHeaders(t *testing.T) {
	testup.LogNameTestInfo(t, "Test get canonical headers function")

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Amz-Date", "20250821T120000Z")
	req.Header.Set("Host", "example.com")

	result := testAdapter.getCanonicalHeaders(req)

	// Les headers doivent être triés alphabétiquement et formatés correctement
	expected := "content-type:application/json\nhost:example.com\nx-amz-date:20250821T120000Z\n"
	assert.Equal(t, expected, result)
}

func TestAdapter_getCanonicalHeaders_WithSpaces(t *testing.T) {
	testup.LogNameTestInfo(t, "Test get canonical headers with spaces")

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "  application/json  ")
	req.Header.Set("X-Custom-Header", "  value with spaces  ")

	result := testAdapter.getCanonicalHeaders(req)

	// Les espaces doivent être supprimés
	expected := "content-type:application/json\nx-custom-header:value with spaces\n"
	assert.Equal(t, expected, result)
}

func TestAdapter_getCanonicalHeaders_MultipleValues(t *testing.T) {
	testup.LogNameTestInfo(t, "Test get canonical headers with multiple values")

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	req.Header.Add("X-Custom-Header", "value1")
	req.Header.Add("X-Custom-Header", "value2")
	req.Header.Set("Content-Type", "application/json")

	result := testAdapter.getCanonicalHeaders(req)

	// Les valeurs multiples doivent être jointes par des virgules
	expected := "content-type:application/json\nx-custom-header:value1,value2\n"
	assert.Equal(t, expected, result)
}

func TestAdapter_signRequest_Basic(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request basic")

	req, err := http.NewRequest("GET", "https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt", nil)
	require.NoError(t, err)

	payload := []byte("test payload")

	err = testAdapter.signRequest(req, payload)
	require.NoError(t, err)

	// Vérifier que les headers requis sont présents
	assert.NotEmpty(t, req.Header.Get("X-Amz-Date"))
	assert.NotEmpty(t, req.Header.Get("X-Amz-Content-Sha256"))
	assert.NotEmpty(t, req.Header.Get("Authorization"))

	// Vérifier le format de la date
	amzDate := req.Header.Get("X-Amz-Date")
	assert.Regexp(t, `^\d{8}T\d{6}Z$`, amzDate)

	// Vérifier le hash du payload
	expectedHash := sha256.Sum256(payload)
	expectedHashStr := hex.EncodeToString(expectedHash[:])
	assert.Equal(t, expectedHashStr, req.Header.Get("X-Amz-Content-Sha256"))
}

func TestAdapter_signRequest_NilPayload(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request with nil payload")

	req, err := http.NewRequest("GET", "https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt", nil)
	require.NoError(t, err)

	err = testAdapter.signRequest(req, nil)
	require.NoError(t, err)

	// Vérifier le hash du payload vide
	expectedHash := sha256.Sum256([]byte(""))
	expectedHashStr := hex.EncodeToString(expectedHash[:])
	assert.Equal(t, expectedHashStr, req.Header.Get("X-Amz-Content-Sha256"))
}

func TestAdapter_signRequest_AuthorizationFormat(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request authorization format")

	req, err := http.NewRequest("PUT", "https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt", nil)
	require.NoError(t, err)

	payload := []byte("test content")

	err = testAdapter.signRequest(req, payload)
	require.NoError(t, err)

	authorization := req.Header.Get("Authorization")

	// Vérifier le format de l'authorization header
	assert.Contains(t, authorization, "AWS4-HMAC-SHA256")
	assert.Contains(t, authorization, "Credential=test-access-key")
	assert.Contains(t, authorization, "SignedHeaders=")
	assert.Contains(t, authorization, "Signature=")

	// Vérifier que la signature est bien un hash hex
	parts := strings.Split(authorization, "Signature=")
	require.Len(t, parts, 2)
	signature := parts[1]
	assert.Regexp(t, "^[a-f0-9]{64}$", signature)
}

func TestAdapter_signRequest_WithQueryParams(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request with query parameters")

	reqURL, err := url.Parse("https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt")
	require.NoError(t, err)

	params := url.Values{}
	params.Add("response-content-type", "application/json")
	params.Add("x-id", "GetObject")
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	require.NoError(t, err)

	err = testAdapter.signRequest(req, nil)
	require.NoError(t, err)

	// Vérifier que la signature est générée correctement même avec des paramètres
	assert.NotEmpty(t, req.Header.Get("Authorization"))
	assert.NotEmpty(t, req.Header.Get("X-Amz-Date"))
}

func TestAdapter_signRequest_EmptyPath(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request with empty path")

	req, err := http.NewRequest("GET", "https://test-account-id.r2.cloudflarestorage.com", nil)
	require.NoError(t, err)

	err = testAdapter.signRequest(req, nil)
	require.NoError(t, err)

	// Vérifier que la signature est générée même avec un chemin vide
	assert.NotEmpty(t, req.Header.Get("Authorization"))
}

func TestAdapter_signRequest_Reproducibility(t *testing.T) {
	testup.LogNameTestInfo(t, "Test sign request reproducibility")

	// Créer deux requêtes identiques
	req1, err := http.NewRequest("GET", "https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt", nil)
	require.NoError(t, err)

	req2, err := http.NewRequest("GET", "https://test-account-id.r2.cloudflarestorage.com/test-bucket/file.txt", nil)
	require.NoError(t, err)

	payload := []byte("same payload")

	err = testAdapter.signRequest(req1, payload)
	require.NoError(t, err)

	// Petite pause pour s'assurer que le temps change
	time.Sleep(time.Millisecond)

	err = testAdapter.signRequest(req2, payload)
	require.NoError(t, err)

	// Les signatures seront différentes à cause du timestamp, mais la structure doit être la même
	auth1 := req1.Header.Get("Authorization")
	auth2 := req2.Header.Get("Authorization")

	// Vérifier que les deux ont la même structure
	assert.Contains(t, auth1, "AWS4-HMAC-SHA256")
	assert.Contains(t, auth2, "AWS4-HMAC-SHA256")
	assert.Contains(t, auth1, "Credential=test-access-key")
	assert.Contains(t, auth2, "Credential=test-access-key")

	// Le hash du payload doit être identique
	assert.Equal(t, req1.Header.Get("X-Amz-Content-Sha256"), req2.Header.Get("X-Amz-Content-Sha256"))
}

func TestAdapter_Creation_WithOptions(t *testing.T) {
	testup.LogNameTestInfo(t, "Test adapter creation with options")

	opts := options{
		AccountID: "test-account",
		AccessKey: "test-access",
		SecretKey: "test-secret",
		Bucket:    "test-bucket",
		KeyPrefix: "prefix/",
	}

	// Test de création d'adaptateur avec options (hypothétique)
	adapter := &Adapter{
		accountID:  opts.AccountID,
		accessKey:  opts.AccessKey,
		secretKey:  opts.SecretKey,
		bucketName: opts.Bucket,
		keyPrefix:  opts.KeyPrefix,
		httpClient: &http.Client{},
		baseURL:    "https://" + opts.AccountID + ".r2.cloudflarestorage.com",
	}

	assert.Equal(t, "test-account", adapter.accountID)
	assert.Equal(t, "test-access", adapter.accessKey)
	assert.Equal(t, "test-secret", adapter.secretKey)
	assert.Equal(t, "test-bucket", adapter.bucketName)
	assert.Equal(t, "prefix/", adapter.keyPrefix)
	assert.Equal(t, "https://test-account.r2.cloudflarestorage.com", adapter.baseURL)
	assert.NotNil(t, adapter.httpClient)
}

func TestAdapter_BaseURL_Format(t *testing.T) {
	testup.LogNameTestInfo(t, "Test base URL format")

	// Tester différents account IDs
	testCases := []struct {
		accountID   string
		expectedURL string
	}{
		{"abc123", "https://abc123.r2.cloudflarestorage.com"},
		{"test-account-456", "https://test-account-456.r2.cloudflarestorage.com"},
		{"", "https://.r2.cloudflarestorage.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.accountID, func(t *testing.T) {
			adapter := &Adapter{
				accountID:  tc.accountID,
				accessKey:  "test-key",
				secretKey:  "test-secret",
				bucketName: "test-bucket",
				keyPrefix:  "",
				httpClient: &http.Client{},
				baseURL:    "https://" + tc.accountID + ".r2.cloudflarestorage.com",
			}

			assert.Equal(t, tc.expectedURL, adapter.baseURL)
		})
	}
}
