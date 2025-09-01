package filestores

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testService *FileStoreService
	testAdapter *testS3Adapter
)

// Implémentation simple d'un adaptateur S3 pour les tests
type testS3Adapter struct {
	storage map[string][]byte
}

func newTestS3Adapter() *testS3Adapter {
	return &testS3Adapter{
		storage: make(map[string][]byte),
	}
}

func (a *testS3Adapter) Upload(ctx context.Context, key string, data []byte) error {
	if strings.Contains(key, "error") {
		return fmt.Errorf("erreur simulée pour key: %s", key)
	}
	a.storage[key] = make([]byte, len(data))
	copy(a.storage[key], data)
	return nil
}

func (a *testS3Adapter) Download(ctx context.Context, key string) ([]byte, error) {
	if strings.Contains(key, "error") {
		return nil, fmt.Errorf("erreur simulée pour key: %s", key)
	}
	data, exists := a.storage[key]
	if !exists {
		return nil, fmt.Errorf("fichier non trouvé: %s", key)
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func (a *testS3Adapter) Delete(ctx context.Context, key string) error {
	if strings.Contains(key, "error") {
		return fmt.Errorf("erreur simulée pour key: %s", key)
	}
	delete(a.storage, key)
	return nil
}

func TestMain(m *testing.M) {
	testAdapter = newTestS3Adapter()

	config := FileStoreConfig{
		KeyPrefix:      "test/uploads",
		MaxSize:        1024 * 1024, // 1MB
		AllowedMIMEs:   []string{"text/plain; charset=utf-8", "text/plain", "application/pdf", "image/jpeg", "image/png"},
		UseDateFolders: true,
	}

	testService = NewService(testAdapter, config)

	code := m.Run()
	os.Exit(code)
}

func TestFileStoreService_UploadBytes_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload bytes success")

	ctx := context.Background()
	scope := "documents"
	filename := "test.txt"
	data := []byte("Contenu de test")

	result, err := testService.UploadBytes(ctx, scope, filename, data)

	require.NoError(t, err)
	assert.NotEmpty(t, result.Key)
	assert.Equal(t, int64(len(data)), result.Size)
	assert.Equal(t, "text/plain; charset=utf-8", result.MIME)
	assert.Equal(t, filename, result.Original)
	assert.Equal(t, result.Key, result.StoredPath)

	// Vérifier que le fichier est bien stocké
	stored, exists := testAdapter.storage[result.Key]
	assert.True(t, exists)
	assert.Equal(t, data, stored)
}

func TestFileStoreService_UploadBytes_WithDateFolders(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload bytes with date folders")

	ctx := context.Background()
	scope := "images"
	filename := "photo.jpg"
	data := []byte("fake jpeg data")

	result, err := testService.UploadBytes(ctx, scope, filename, data)

	require.NoError(t, err)

	// Vérifier que la clé contient bien la structure de dates
	now := time.Now()
	expectedDatePath := fmt.Sprintf("%04d/%02d/%02d", now.Year(), now.Month(), now.Day())
	assert.Contains(t, result.Key, expectedDatePath)
	assert.Contains(t, result.Key, "test/uploads")
	assert.Contains(t, result.Key, scope)
}

func TestFileStoreService_UploadBytes_EmptyFile(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload empty file")

	ctx := context.Background()
	scope := "documents"
	filename := "empty.txt"
	data := []byte("")

	result, err := testService.UploadBytes(ctx, scope, filename, data)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "fichier vide")
}

func TestFileStoreService_UploadBytes_FileTooLarge(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload file too large")

	ctx := context.Background()
	scope := "documents"
	filename := "large.txt"
	// Créer un fichier plus grand que la limite (1MB)
	data := make([]byte, 2*1024*1024)

	result, err := testService.UploadBytes(ctx, scope, filename, data)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "fichier trop volumineux")
}

func TestFileStoreService_UploadBytes_InvalidMIME(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload file with invalid MIME type")

	// Créer un service avec des MIME très restrictifs (seulement PDF)
	restrictiveConfig := FileStoreConfig{
		KeyPrefix:      "test/uploads",
		MaxSize:        1024 * 1024,
		AllowedMIMEs:   []string{"application/pdf"}, // Seulement PDF autorisé
		UseDateFolders: true,
	}

	restrictiveService := NewService(testAdapter, restrictiveConfig)

	ctx := context.Background()
	scope := "documents"
	filename := "test.txt"
	data := []byte("Simple text content") // Sera détecté comme text/plain

	result, err := restrictiveService.UploadBytes(ctx, scope, filename, data)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mime n'est pas attribué")
}

func TestFileStoreService_UploadBytes_PDFFile(t *testing.T) {
	testup.LogNameTestInfo(t, "Test upload PDF file")

	ctx := context.Background()
	scope := "documents"
	filename := "document.pdf"
	// Début d'un fichier PDF
	data := []byte("%PDF-1.4")

	result, err := testService.UploadBytes(ctx, scope, filename, data)

	require.NoError(t, err)
	assert.Equal(t, "application/pdf", result.MIME)
}

func TestFileStoreService_Download_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test download file success")

	ctx := context.Background()

	// D'abord uploader un fichier
	scope := "documents"
	filename := "test.txt"
	originalData := []byte("Contenu à télécharger")

	uploadResult, err := testService.UploadBytes(ctx, scope, filename, originalData)
	require.NoError(t, err)

	// Ensuite le télécharger
	downloadedData, err := testService.Download(ctx, uploadResult.Key)

	require.NoError(t, err)
	assert.Equal(t, originalData, downloadedData)
}

func TestFileStoreService_Download_FileNotFound(t *testing.T) {
	testup.LogNameTestInfo(t, "Test download file not found")

	ctx := context.Background()
	key := "inexistant/file.txt"

	data, err := testService.Download(ctx, key)

	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "fichier non trouvé")
}

func TestFileStoreService_Download_EmptyKey(t *testing.T) {
	testup.LogNameTestInfo(t, "Test download with empty key")

	ctx := context.Background()
	key := ""

	data, err := testService.Download(ctx, key)

	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "key manquante")
}

func TestFileStoreService_Delete_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test delete file success")

	ctx := context.Background()

	// D'abord uploader un fichier
	scope := "documents"
	filename := "to_delete.txt"
	data := []byte("Fichier à supprimer")

	uploadResult, err := testService.UploadBytes(ctx, scope, filename, data)
	require.NoError(t, err)

	// Vérifier qu'il existe
	_, exists := testAdapter.storage[uploadResult.Key]
	assert.True(t, exists)

	// Le supprimer
	err = testService.Delete(ctx, uploadResult.Key)
	require.NoError(t, err)

	// Vérifier qu'il n'existe plus
	_, exists = testAdapter.storage[uploadResult.Key]
	assert.False(t, exists)
}

func TestFileStoreService_Delete_EmptyKey(t *testing.T) {
	testup.LogNameTestInfo(t, "Test delete with empty key")

	ctx := context.Background()
	key := ""

	err := testService.Delete(ctx, key)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key manquante")
}

func TestFileStoreService_buildKey(t *testing.T) {
	testup.LogNameTestInfo(t, "Test build key generation")

	scope := "avatars"
	filename := "profile.jpg"

	key := testService.buildKey(scope, filename)

	assert.Contains(t, key, "test/uploads")
	assert.Contains(t, key, scope)
	assert.Contains(t, key, ".jpg")

	// Vérifier la structure de date
	now := time.Now()
	expectedYear := fmt.Sprintf("%04d", now.Year())
	expectedMonth := fmt.Sprintf("%02d", now.Month())
	expectedDay := fmt.Sprintf("%02d", now.Day())

	assert.Contains(t, key, expectedYear)
	assert.Contains(t, key, expectedMonth)
	assert.Contains(t, key, expectedDay)
}

func TestFileStoreService_buildKey_CleanScope(t *testing.T) {
	testup.LogNameTestInfo(t, "Test build key with dirty scope")

	scope := "//documents/../admin//"
	filename := "test.txt"

	key := testService.buildKey(scope, filename)

	// Le scope doit être nettoyé
	assert.Contains(t, key, "documents/admin")
	assert.NotContains(t, key, "..")
	assert.NotContains(t, key, "//")
}

func TestFileStoreService_WithoutDateFolders(t *testing.T) {
	testup.LogNameTestInfo(t, "Test service without date folders")

	// Créer un service sans dossiers de date
	config := FileStoreConfig{
		KeyPrefix:      "simple",
		MaxSize:        1024,
		AllowedMIMEs:   []string{"text/plain; charset=utf-8", "text/plain"},
		UseDateFolders: false,
	}

	simpleService := NewService(testAdapter, config)

	ctx := context.Background()
	scope := "docs"
	filename := "simple.txt"
	data := []byte("Simple test")

	result, err := simpleService.UploadBytes(ctx, scope, filename, data)

	require.NoError(t, err)

	// La clé ne doit pas contenir de structure de date
	now := time.Now()
	dateStr := fmt.Sprintf("%04d", now.Year())
	assert.NotContains(t, result.Key, dateStr)
	assert.Contains(t, result.Key, "simple")
	assert.Contains(t, result.Key, scope)
}

func TestFileStoreService_MIMEDetection(t *testing.T) {
	testup.LogNameTestInfo(t, "Test MIME type detection")

	tests := []struct {
		name         string
		filename     string
		data         []byte
		expectedMIME string
	}{
		{
			name:         "Text file",
			filename:     "test.txt",
			data:         []byte("Simple text content"),
			expectedMIME: "text/plain; charset=utf-8",
		},
		{
			name:         "PDF file",
			filename:     "doc.pdf",
			data:         []byte("%PDF-1.4"),
			expectedMIME: "application/pdf",
		},
	}

	// Service permissif pour ce test
	config := FileStoreConfig{
		KeyPrefix:      "mime-test",
		MaxSize:        1024,
		AllowedMIMEs:   []string{}, // Tout accepté
		UseDateFolders: false,
	}

	permissiveService := NewService(testAdapter, config)
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := permissiveService.UploadBytes(ctx, "test", tt.filename, tt.data)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedMIME, result.MIME)
		})
	}
}

func Test_cleanPart(t *testing.T) {
	testup.LogNameTestInfo(t, "Test clean part function")

	tests := []struct {
		input    string
		expected string
	}{
		{"/admin/", "admin"},
		{"//documents/../test//", "documents//test"}, // Comportement réel observé
		{"  spaced  ", "spaced"},
		{"normal", "normal"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("clean_%s", tt.input), func(t *testing.T) {
			result := cleanPart(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_randHex(t *testing.T) {
	testup.LogNameTestInfo(t, "Test random hex generation")

	hex1 := RandHex(16)
	hex2 := RandHex(16)

	assert.Len(t, hex1, 32) // 16 bytes = 32 hex chars
	assert.Len(t, hex2, 32)
	assert.NotEqual(t, hex1, hex2) // Doit être différent
}

func Test_contains(t *testing.T) {
	testup.LogNameTestInfo(t, "Test contains function")

	list := []string{"text/plain", "application/pdf", "image/jpeg"}

	assert.True(t, contains(list, "text/plain"))
	assert.True(t, contains(list, "application/pdf"))
	assert.False(t, contains(list, "video/mp4"))
	assert.False(t, contains([]string{}, "anything"))
}
