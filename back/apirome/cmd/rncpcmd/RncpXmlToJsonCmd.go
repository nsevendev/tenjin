package rncpcmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/rncp"
	"time"

	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
)

// ConvertXMLToJSONCmd est la commande pour convertir le fichier XML RNCP en JSON
// Elle lit le dernier fichier XML trouvé dans le répertoire spécifié, le parse, le convertit en structures MongoDB,
// et sauvegarde le résultat en JSON dans le même répertoire avec un nom de fichier basé sur la date et l'heure actuelle.
var ConvertXMLToJSONCmd = &cobra.Command{
	Use:   "rncp-xml-to-json",
	Short: "Convertit le fichier XML RNCP en format JSON",
	Run: func(cmd *cobra.Command, args []string) {
		fileDir := "./apirome/rncp/data/"

		// Utiliser la même logique que vos autres commandes
		xmlFile := libs.GetLastXMLFileByDate(fileDir)
		if xmlFile == "" {
			logger.Ff("Aucun fichier trouvé dans le répertoire %s", fileDir)
			return
		}

		logger.If("Début de la conversion XML → JSON pour : %s", xmlFile)

		// 1. Lire le fichier XML
		xmlData, err := os.ReadFile(xmlFile)
		if err != nil {
			logger.Ff("Erreur lors de la lecture du fichier XML : %v", err)
			return
		}

		// 2. Parser le XML
		var xmlFiches rncp.XMLFiches
		if err := xml.Unmarshal(xmlData, &xmlFiches); err != nil {
			logger.Ff("Erreur lors du parsing XML : %v", err)
			return
		}

		logger.If("XML parsé avec succès - Version flux: %s", xmlFiches.VersionFlux)
		logger.If("Nombre de fiches trouvées : %d", len(xmlFiches.Fiches))

		// 3. Convertir en structures MongoDB
		var certifications []rncp.Certification
		for _, xmlFiche := range xmlFiches.Fiches {
			cert := rncp.ConvertXMLToMongo(xmlFiche)
			certifications = append(certifications, cert)
		}

		logger.If("Conversion terminée : %d certifications", len(certifications))

		// 4. Créer le répertoire de sortie si nécessaire
		jsonOutputDir := fileDir
		if err := os.MkdirAll(jsonOutputDir, 0755); err != nil {
			logger.Ff("Erreur création répertoire : %v", err)
			return
		}

		// 5. Générer le nom de fichier avec timestamp
		timestamp := time.Now().Format("20060102_150405")
		outputFile := filepath.Join(jsonOutputDir, fmt.Sprintf("%s.json", timestamp))

		// 6. Sauvegarder en JSON
		jsonData, err := json.MarshalIndent(certifications, "", "  ")
		if err != nil {
			logger.Ff("Erreur lors de la sérialisation JSON : %v", err)
			return
		}

		if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
			logger.Ff("Erreur lors de l'écriture du fichier JSON : %v", err)
			return
		}

		logger.If("Fichier JSON généré avec succès : %s", outputFile)
		logger.If("Taille du fichier : %.2f MB", float64(len(jsonData))/(1024*1024))

		// 7. Affichage de statistiques
		showStats(certifications)
	},
}

func showStats(certifications []rncp.Certification) {
	logger.If("=== STATISTIQUES ===")

	// Compter par niveau
	niveaux := make(map[string]int)
	etats := make(map[string]int)
	actifs := 0

	for _, cert := range certifications {
		niveaux[cert.Niveau]++
		etats[cert.EtatFiche]++
		if cert.Actif {
			actifs++
		}
	}

	logger.If("Répartition par niveau :")
	for niveau, count := range niveaux {
		logger.If("  %s: %d", niveau, count)
	}

	logger.If("Répartition par état :")
	for etat, count := range etats {
		logger.If("  %s: %d", etat, count)
	}

	logger.If("Certifications actives : %d/%d", actifs, len(certifications))
}

// findXMLFile cherche le fichier XML dans le répertoire spécifié
func findXMLFile(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le répertoire %s : %w", dir, err)
	}

	// Chercher un fichier .xml
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".xml" {
			fullPath := filepath.Join(dir, file.Name())
			logger.If("Fichier XML trouvé : %s", fullPath)
			return fullPath, nil
		}
	}

	return "", fmt.Errorf("aucun fichier XML trouvé dans le répertoire %s", dir)
}
