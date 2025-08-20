package filestores

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/logger/v2/logger"
	"io"
	"mime/multipart"
)

// FileData représente les données d'un fichier extrait du form-data
type FileData struct {
	Data     []byte
	Filename string
	Size     int64
	Header   *multipart.FileHeader
}

// ExtractFileFromForm extrait un fichier du form-file et retourne ses données
// fieldName: nom du champ dans le form (ex: "file")
// Retourne: *FileData si succès, nil si erreur (l'erreur est automatiquement envoyée au client)
func (s *FileStoreService) ExtractFileFromForm(c *gin.Context, fieldName string) (*FileData, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return nil, fmt.Errorf("fichier manquant : %v", err)
	}

	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier uploadé : %v", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	logger.Sf("Fichier '%s' extrait avec succès (%d bytes)", file.Filename, len(data))

	return &FileData{
		Data:     data,
		Filename: file.Filename,
		Size:     file.Size,
		Header:   file,
	}, nil
}
