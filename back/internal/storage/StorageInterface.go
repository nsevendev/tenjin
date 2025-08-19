package storage

import "context"

// StorageInterface est le contrat pour les adaptateurs de stockage.
// Il définit les méthodes nécessaires pour uploader, télécharger et supprimer des fichiers.
// Chaque adaptateur (R2, S3, etc.) doit implémenter cette interface
type StorageInterface interface {
	Upload(ctx context.Context, key string, data []byte) error
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}
