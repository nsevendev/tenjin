package filestores

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"tenjin/back/internal/utils/s3adapter"
	"time"
)

type FileStoreConfig struct {
	// Préfixe commun dans le bucket (ex: "tenjin/uploads/").
	KeyPrefix string
	// Taille max autorisée (en octets). 0 = pas de limite.
	MaxSize int64
	// Liste blanche des MIME (vide = tout accepté).
	AllowedMIMEs []string
	// Génération de chemin daté (ex: 2025/08/18)
	UseDateFolders bool
}

type UploadStoreResult struct {
	Key        string // clé stockée dans R2
	Size       int64  // taille du fichier en octets
	MIME       string // type MIME du fichier
	Original   string // nom de fichier original
	StoredPath string // = Key (alias, utile s’il y a un CDN plus tard)
}

type FileStoreService struct {
	s3Adapter s3adapter.AdapterInterface
	config    FileStoreConfig
}

func NewService(s3Adapter s3adapter.AdapterInterface, cfg FileStoreConfig) *FileStoreService {
	return &FileStoreService{s3Adapter: s3Adapter, config: cfg}
}

// UploadBytes est la primitive d’upload.
// scope devient un sous-dossier logique (ex: "avatars", "docs", ...).
func (s *FileStoreService) UploadBytes(ctx context.Context, scope, filename string, data []byte) (*UploadStoreResult, error) {
	if len(data) == 0 {
		return nil, errors.New("fichier vide")
	}
	if s.config.MaxSize > 0 && int64(len(data)) > s.config.MaxSize {
		return nil, fmt.Errorf("fichier trop volumineux: %d > %d", len(data), s.config.MaxSize)
	}

	// MIME: essaie par contenu, fallback extension
	mimeType := http.DetectContentType(peek(data, 512))
	if mimeType == "application/octet-stream" {
		if ext := strings.ToLower(filepath.Ext(filename)); ext != "" {
			if byExt := mime.TypeByExtension(ext); byExt != "" {
				mimeType = byExt
			}
		}
	}

	// Filtrage MIME si configuré
	if len(s.config.AllowedMIMEs) > 0 && !contains(s.config.AllowedMIMEs, mimeType) {
		return nil, fmt.Errorf("mime n'est pas attribué: %s", mimeType)
	}

	key := s.buildKey(scope, filename)

	logger.If("DEBUG AdapterInterface FileStoreService:")
	logger.If("  - scope: %q", scope)
	logger.If("  - filename: %q", filename)
	logger.If("  - config.KeyPrefix: %q", s.config.KeyPrefix)
	logger.If("  - config.UseDateFolders: %v", s.config.UseDateFolders)
	logger.If("  - key générée: %q", key)

	if err := s.s3Adapter.Upload(ctx, key, data); err != nil {
		return nil, fmt.Errorf("erreur lors de l'upload vers le s3: %w", err)
	}

	return &UploadStoreResult{
		Key:        key,
		Size:       int64(len(data)),
		MIME:       mimeType,
		Original:   filename,
		StoredPath: key,
	}, nil
}

func (s *FileStoreService) Download(ctx context.Context, key string) ([]byte, error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("key manquante")
	}
	return s.s3Adapter.Download(ctx, key)
}

func (s *FileStoreService) Delete(ctx context.Context, key string) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("key manquante")
	}
	return s.s3Adapter.Delete(ctx, key)
}

// ---------- helpers ----------

func (s *FileStoreService) buildKey(scope, filename string) string {
	scope = cleanPart(scope)
	prefix := strings.Trim(s.config.KeyPrefix, "/")
	now := time.Now()

	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		ext = guessExtByName(filename)
	}
	name := randHex(16) + ext

	var parts []string
	if prefix != "" {
		parts = append(parts, prefix)
	}
	if scope != "" {
		parts = append(parts, scope)
	}
	if s.config.UseDateFolders {
		parts = append(parts,
			fmt.Sprintf("%04d", now.Year()),
			fmt.Sprintf("%02d", now.Month()),
			fmt.Sprintf("%02d", now.Day()),
		)
	}
	parts = append(parts, name)
	return path.Clean(strings.Join(parts, "/"))
}

func cleanPart(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "/")
	s = strings.ReplaceAll(s, "..", "")
	return s
}

func randHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func peek(b []byte, n int) []byte {
	if len(b) < n {
		return b
	}
	return b[:n]
}

func contains(list []string, v string) bool {
	for _, e := range list {
		if e == v {
			return true
		}
	}
	return false
}

func guessExtByName(name string) string {
	ext := strings.ToLower(filepath.Ext(name))
	if ext != "" {
		return ext
	}
	return ""
}
