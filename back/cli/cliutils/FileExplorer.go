package cliutils

import (
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// FileExtension définit les types de fichiers supportés
type FileExtension string

const (
	ExtJSON FileExtension = "json"
	ExtXML  FileExtension = "xml"
)

// FileConfig contient la configuration pour chaque type de fichier
type FileConfig struct {
	Regex       *regexp.Regexp
	DateFormats []string
}

type FileExplorer struct {
	BasePath    string
	fileConfigs map[FileExtension]FileConfig
}

type datedFile struct {
	name string
	date time.Time
}

func NewFileExplorer(basePath string) *FileExplorer {
	// Configuration des types de fichiers
	fileConfigs := map[FileExtension]FileConfig{
		ExtJSON: {
			Regex:       regexp.MustCompile(`(\d{8}_\d{6})\.json$`),
			DateFormats: []string{"20060102_150405"},
		},
		ExtXML: {
			Regex:       regexp.MustCompile(`(\d{4}-\d{2}-\d{2}|\d{8})\.xml$`),
			DateFormats: []string{"2006-01-02", "20060102"},
		},
	}

	return &FileExplorer{
		BasePath:    basePath,
		fileConfigs: fileConfigs,
	}
}

// GetLastFileByDateAndType récupère le dernier fichier du type spécifié dans le dossier donné, trié par date
func (fe *FileExplorer) GetLastFileByDateAndType(dir string, fileType FileExtension) (string, error) {
	fullPath := filepath.Join(fe.BasePath, dir)
	files, err := os.ReadDir(fullPath)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du dossier %s : %v", fullPath, err)
	}

	// Récupérer la configuration pour ce type de fichier
	config, exists := fe.fileConfigs[fileType]
	if !exists {
		return "", fmt.Errorf("type de fichier non supporté: %s", fileType)
	}

	var datedFiles []datedFile

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		matches := config.Regex.FindStringSubmatch(file.Name())
		if len(matches) != 2 {
			continue
		}

		var parsed time.Time
		var parseErr error

		// Essayer tous les formats de date possibles pour ce type
		for _, format := range config.DateFormats {
			parsed, parseErr = time.Parse(format, matches[1])
			if parseErr == nil {
				break // Format trouvé
			}
		}

		if parseErr != nil {
			logger.Ef("Erreur lors de la parsing de la date pour le fichier %s : %v", file.Name(), parseErr)
			continue // Aucun format ne correspond
		}

		datedFiles = append(datedFiles, datedFile{
			name: filepath.Join(fullPath, file.Name()),
			date: parsed,
		})
	}

	if len(datedFiles) == 0 {
		return "", fmt.Errorf("aucun fichier %s trouvé dans le dossier %s", strings.ToUpper(string(fileType)), fullPath)
	}

	// Trie par date décroissante (plus récent en premier), return le resultat dans le meme slice
	sort.Slice(datedFiles, func(i, j int) bool {
		return datedFiles[i].date.After(datedFiles[j].date)
	})

	return datedFiles[0].name, nil
}
