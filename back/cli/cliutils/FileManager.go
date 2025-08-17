package cliutils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"os"
	"path/filepath"
	"time"
)

// FileType définit le type de fichier à créer
type FileType string

const (
	FileTypeJSON FileType = "json"
	FileTypeXML  FileType = "xml"
	FileTypeLOG  FileType = "log"
	FileTypeTXT  FileType = "txt"
)

// FileManager gère la création et l'écriture de fichiers
type FileManager struct {
	BasePath   string   // Répertoire de base
	FileType   FileType // Type de fichier par défaut
	CreateDirs bool     // Créer les répertoires automatiquement
	Timestamp  bool     // Ajouter timestamp au nom de fichier
}

// FileOptions permet de personnaliser l'écriture
type FileOptions struct {
	FileType     *FileType   // Override du type de fichier
	CustomName   string      // Nom personnalisé (sans extension)
	SubDir       string      // Sous-répertoire
	AddTimestamp *bool       // Override du timestamp
	Indent       bool        // Indentation pour FileTypeJSON/FileTypeXML
	FileMode     os.FileMode // Permissions du fichier (défaut: 0644)
}

// NewFileManager crée un nouveau gestionnaire de fichiers
func NewFileManager(basePath string, defaultType FileType) *FileManager {
	return &FileManager{
		BasePath:   basePath,
		FileType:   defaultType,
		CreateDirs: true,
		Timestamp:  true,
	}
}

// SetDefaults configure les options par défaut
func (fm *FileManager) SetDefaults(createDirs, timestamp bool) *FileManager {
	fm.CreateDirs = createDirs
	fm.Timestamp = timestamp
	return fm
}

// WriteData écrit des données dans un fichier selon le type spécifié
func (fm *FileManager) WriteData(data interface{}, options *FileOptions) (string, error) {
	// Configuration par défaut
	if options == nil {
		options = &FileOptions{}
	}

	// Déterminer le type de fichier
	fileType := fm.FileType
	if options.FileType != nil {
		fileType = *options.FileType
	}

	// Générer le nom de fichier
	fileName, err := fm.generateFileName(options, fileType)
	if err != nil {
		return "", fmt.Errorf("erreur génération nom fichier: %w", err)
	}

	// Construire le chemin complet
	fullPath := filepath.Join(fm.BasePath, options.SubDir, fileName)

	// Créer les répertoires si nécessaire
	if fm.CreateDirs {
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return "", fmt.Errorf("erreur création répertoire %s: %w", filepath.Dir(fullPath), err)
		}
	}

	// Sérialiser les données selon le type
	content, err := fm.serializeData(data, fileType, options.Indent)
	if err != nil {
		return "", fmt.Errorf("erreur sérialisation: %w", err)
	}

	// Déterminer les permissions
	fileMode := os.FileMode(0644)
	if options.FileMode != 0 {
		fileMode = options.FileMode
	}

	// Écrire le fichier
	if err := os.WriteFile(fullPath, content, fileMode); err != nil {
		return "", fmt.Errorf("erreur écriture fichier %s: %w", fullPath, err)
	}

	logger.If("Fichier généré avec succès : %v", fullPath)

	return fullPath, nil
}

// WriteErrors est un helper spécialisé pour les fichiers d'erreur
func (fm *FileManager) WriteErrors(errors interface{}) (string, error) {
	options := &FileOptions{
		FileType:     &[]FileType{FileTypeJSON}[0],
		CustomName:   "erreurs",
		AddTimestamp: &[]bool{true}[0],
		Indent:       true,
	}
	return fm.WriteData(errors, options)
}

// WriteLog est un helper spécialisé pour les logs
func (fm *FileManager) WriteLog(logContent string) (string, error) {
	options := &FileOptions{
		FileType:     &[]FileType{FileTypeLOG}[0],
		CustomName:   "application",
		AddTimestamp: &[]bool{true}[0],
	}
	return fm.WriteData(logContent, options)
}

// AppendLog ajoute du contenu à un fichier de log existant
func (fm *FileManager) AppendLog(logContent string, options *FileOptions) error {
	if options == nil {
		options = &FileOptions{}
	}

	// Forcer le type FileTypeLOG
	logType := FileTypeLOG
	options.FileType = &logType

	fileName, err := fm.generateFileName(options, FileTypeLOG)
	if err != nil {
		return fmt.Errorf("erreur génération nom du fichier: %w", err)
	}

	fullPath := filepath.Join(fm.BasePath, options.SubDir, fileName)

	// Créer les répertoires si nécessaire
	if fm.CreateDirs {
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("erreur création répertoire: %w", err)
		}
	}

	// Ouvrir en mode append
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("erreur ouverture fichier: %w", err)
	}
	defer file.Close()

	// Ajouter timestamp au contenu de log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("[%s] %s\n", timestamp, logContent)

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("erreur écriture dans fichier: %w", err)
	}

	return nil
}

// generateFileName génère un nom de fichier selon les options
func (fm *FileManager) generateFileName(options *FileOptions, fileType FileType) (string, error) {
	var baseName string

	// Nom de base
	if options.CustomName != "" {
		baseName = options.CustomName
	} else {
		baseName = ""
	}

	// Ajouter timestamp si demandé
	addTimestamp := fm.Timestamp
	if options.AddTimestamp != nil {
		addTimestamp = *options.AddTimestamp
	}

	if addTimestamp {
		timestamp := time.Now().Format("20060102_150405")
		if baseName != "" {
			baseName = fmt.Sprintf("%v_%v", timestamp, baseName)
		} else {
			baseName = fmt.Sprintf("%v", timestamp)
		}
	}

	// Ajouter extension
	extension := string(fileType)
	fileName := fmt.Sprintf("%v.%v", baseName, extension)

	return fileName, nil
}

// serializeData sérialise les données selon le type de fichier
func (fm *FileManager) serializeData(data interface{}, fileType FileType, indent bool) ([]byte, error) {
	switch fileType {
	case FileTypeJSON:
		if indent {
			return json.MarshalIndent(data, "", "  ")
		}
		return json.Marshal(data)

	case FileTypeXML:
		if indent {
			return xml.MarshalIndent(data, "", "  ")
		}
		return xml.Marshal(data)

	case FileTypeLOG, FileTypeTXT:
		// Pour les logs et texte, on s'attend à recevoir une string
		if str, ok := data.(string); ok {
			return []byte(str), nil
		}
		return nil, fmt.Errorf("pour les fichiers FileTypeLOG/FileTypeTXT, les données doivent être une string")

	default:
		return nil, fmt.Errorf("type de fichier non supporté: %s", fileType)
	}
}

// FileExists vérifie si un fichier existe
func (fm *FileManager) FileExists(fileName string, subDir string) bool {
	fullPath := filepath.Join(fm.BasePath, subDir, fileName)
	_, err := os.Stat(fullPath)
	return err == nil
}

// ListFiles liste les fichiers d'un répertoire avec un pattern optionnel
func (fm *FileManager) ListFiles(subDir string, pattern string) ([]string, error) {
	dirPath := filepath.Join(fm.BasePath, subDir)

	if pattern == "" {
		pattern = "*"
	}

	files, err := filepath.Glob(filepath.Join(dirPath, pattern))
	if err != nil {
		return nil, fmt.Errorf("erreur listage fichiers: %w", err)
	}

	// Retourner seulement les noms de fichier, pas les chemins complets
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, filepath.Base(file))
	}

	return fileNames, nil
}

// Exemples d'utilisation
/*func ExampleUsage() {
	// Création du gestionnaire
	fm := NewFileManager("./data", FileTypeJSON).SetDefaults(true, true)

	// === EXEMPLE 1: Fichier d'erreurs ===
	type ErrorEntry struct {
		Code      string    `json:"code"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
	}

	errors := []ErrorEntry{
		{Code: "ERR001", Message: "Erreur test", Timestamp: time.Now()},
	}

	// Sauvegarde automatique avec timestamp
	errorFile, err := fm.WriteErrors(errors)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		fmt.Printf("Fichier d'erreurs créé: %s\n", errorFile)
	}

	// === EXEMPLE 2: Données personnalisées ===
	data := map[string]interface{}{
		"competences": []string{"Go", "API", "FileTypeJSON"},
		"count":       3,
	}

	options := &FileOptions{
		FileType:   &[]FileType{FileTypeXML}[0], // Override vers FileTypeXML
		CustomName: "competences",
		SubDir:     "exports",
		Indent:     true,
	}

	dataFile, err := fm.WriteData(data, options)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
	} else {
		fmt.Printf("Fichier de données créé: %s\n", dataFile)
	}

	// === EXEMPLE 3: Log avec append ===
	logOptions := &FileOptions{
		CustomName: "application",
		SubDir:     "logs",
	}

	err = fm.AppendLog("Application démarrée", logOptions)
	if err != nil {
		fmt.Printf("Erreur log: %v\n", err)
	}
}*/
