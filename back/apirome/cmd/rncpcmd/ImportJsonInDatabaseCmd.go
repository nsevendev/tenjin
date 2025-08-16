package rncpcmd

import (
	"encoding/json"
	"os"

	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/rncp"
	"tenjin/back/internal/utils/database"
)

// ImportDataRNCPCmd est la commande pour importer les certifications RNCP en base MongoDB
// Elle lit le dernier fichier JSON généré, le parse, et insère les données dans la collection "certifications"
// Optionnellement, elle peut vider la collection existante avant l'import
var ImportDataRNCPCmd = &cobra.Command{
	Use:   "import-rncp-database",
	Short: "Importe toutes les certifications RNCP en base MongoDB",
	Run: func(cmd *cobra.Command, args []string) {
		// Connexion à la base
		database.ConnexionDatabase("dev")
		collection := database.Client.Collection("certifications")

		// 1. Trouve le dernier fichier JSON généré
		file := libs.GetLastFileByDate("./apirome/rncp/data")
		logger.If("Fichier à importer : %s", file)

		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("Erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
			return
		}
		defer fileData.Close()

		// 2. Parse le contenu JSON
		var certifications []rncp.Certification
		if err := json.NewDecoder(fileData).Decode(&certifications); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
			return
		}

		logger.If("Nombre de certifications à importer : %d", len(certifications))

		// 3. Optionnel : vider la collection existante
		if clearExisting, _ := cmd.Flags().GetBool("clear"); clearExisting {
			logger.If("Suppression des données existantes...")
			if _, err := collection.DeleteMany(nil, map[string]interface{}{}); err != nil {
				logger.Wf("Erreur lors du nettoyage : %v", err)
			} else {
				logger.If("Collection vidée avec succès")
			}
		}

		// 4. Import en bulk
		var docs []interface{}
		for _, certification := range certifications {
			docs = append(docs, certification)
		}

		logger.If("Import en cours...")
		result, err := collection.InsertMany(nil, docs)
		if err != nil {
			logger.Ff("Erreur insertion MongoDB : %v", err)
			return
		}

		logger.If("Import réussi !")
		logger.If("📊 %d certifications insérées", len(result.InsertedIDs))

		// 5. Statistiques finales
		showImportStats(certifications)
	},
}

func showImportStats(certifications []rncp.Certification) {
	logger.If("=== STATISTIQUES D'IMPORT ===")

	// Compter les certifications par état et niveau
	niveauxTotal := make(map[string]int)
	niveauxActifs := make(map[string]int)
	etats := make(map[string]int)
	totalActifs := 0
	codesROMEUniques := make(map[string]bool)
	certificateursUniques := make(map[string]bool)

	for _, cert := range certifications {
		// Compter tous les niveaux
		niveauxTotal[cert.Niveau]++

		// Compter les états
		etats[cert.EtatFiche]++

		// Compter les actifs par niveau
		if cert.Actif {
			niveauxActifs[cert.Niveau]++
			totalActifs++
		}

		// Compter les codes ROME uniques
		for _, rome := range cert.CodesROME {
			if rome.Code != "" {
				codesROMEUniques[rome.Code] = true
			}
		}

		// Compter les certificateurs uniques
		for _, certificateur := range cert.Certificateurs {
			if certificateur.Nom != "" {
				certificateursUniques[certificateur.Nom] = true
			}
		}
	}

	logger.If("Répartition par niveau (total) :")
	for niveau, count := range niveauxTotal {
		if niveau != "" {
			actifs := niveauxActifs[niveau]
			logger.If("  %s: %d (dont %d actifs)", niveau, count, actifs)
		} else {
			logger.If("  [Niveau vide]: %d", count)
		}
	}

	logger.If("Répartition par état :")
	for etat, count := range etats {
		logger.If("  %s: %d", etat, count)
	}

	logger.If("Total certifications actives : %d/%d (%.1f%%)",
		totalActifs, len(certifications),
		float64(totalActifs)/float64(len(certifications))*100)
	logger.If("Nombre de codes ROME différents : %d", len(codesROMEUniques))
	logger.If("Nombre de certificateurs différents : %d", len(certificateursUniques))

	logger.If("✅ Import terminé avec succès")
}

func init() {
	ImportDataRNCPCmd.Flags().Bool("clear", false, "Vider la collection avant import")
}
