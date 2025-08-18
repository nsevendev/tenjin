package storage

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	// Préfixe commun dans le bucket (ex: "tenjin/uploads/").
	KeyPrefix string
	// Taille max autorisée (en octets). 0 = pas de limite.
	MaxSize int64
	// Liste blanche des MIME (vide = tout accepté).
	AllowedMIMEs []string
	// Génération de chemin daté (ex: 2025/08/18)
	UseDateFolders bool
}

type UploadResult struct {
	Key        string // clé stockée dans R2
	Size       int64
	MIME       string
	Original   string // nom de fichier original
	StoredPath string // = Key (alias, utile s’il y a un CDN plus tard)
}

type Service struct {
	store  Storage
	config Config
}

func NewService(store Storage, cfg Config) *Service {
	return &Service{store: store, config: cfg}
}

// UploadBytes est la primitive d’upload.
// scope devient un sous-dossier logique (ex: "avatars", "docs", ...).
func (s *Service) UploadBytes(ctx context.Context, scope, filename string, data []byte) (*UploadResult, error) {
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

	// DEBUG: Ajouter des logs pour voir ce qui se passe
	fmt.Printf("DEBUG Storage Service:\n")
	fmt.Printf("  - scope: %q\n", scope)
	fmt.Printf("  - filename: %q\n", filename)
	fmt.Printf("  - config.KeyPrefix: %q\n", s.config.KeyPrefix)
	fmt.Printf("  - config.UseDateFolders: %v\n", s.config.UseDateFolders)
	fmt.Printf("  - key générée: %q\n", key)

	if err := s.store.Upload(ctx, key, data); err != nil {
		return nil, err
	}

	return &UploadResult{
		Key:        key,
		Size:       int64(len(data)),
		MIME:       mimeType,
		Original:   filename,
		StoredPath: key,
	}, nil
}

func (s *Service) Download(ctx context.Context, key string) ([]byte, error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("missing key")
	}
	return s.store.Download(ctx, key)
}

func (s *Service) Delete(ctx context.Context, key string) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("missing key")
	}
	return s.store.Delete(ctx, key)
}

// ---------- helpers ----------

func (s *Service) buildKey(scope, filename string) string {
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
