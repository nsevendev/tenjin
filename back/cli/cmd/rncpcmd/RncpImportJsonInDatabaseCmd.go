package rncpcmd

import (
	"encoding/json"
	"fmt"
	"os"
	"tenjin/back/cli/cliutils"
	"tenjin/back/cli/internal/rncp"

	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/internal/utils/database"
)

// RncpImportDataInDatabaseCmd est la commande pour importer les certifications RNCP en base MongoDB
// Elle lit le dernier fichier JSON généré, le parse, et insère les données dans la collection "certifications"
// Optionnellement, elle peut vider la collection existante avant l'import
var RncpImportDataInDatabaseCmd = &cobra.Command{
	Use:   "import-rncp-database",
	Short: "Importe toutes les certifications RNCP en base MongoDB",
	Run: func(cmd *cobra.Command, args []string) {
		clearExisting, _ := cmd.Flags().GetBool("clear")
		database.ConnexionDatabase("dev")
		collection := database.Client.Collection("certifications")
		fileExplorer := cliutils.NewFileExplorer("./cli/data/rncp/data-json")
		fileManagerLog := cliutils.NewFileManager("./cli/data/rncp/log", cliutils.FileTypeLOG)

		file, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtJSON)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.If("Fichier à importer : %s", file)

		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("Erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
			return
		}
		defer fileData.Close()

		var certifications []rncp.Certification
		if err := json.NewDecoder(fileData).Decode(&certifications); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
			return
		}

		logger.If("Nombre de certifications à importer : %d", len(certifications))

		if clearExisting {
			logger.If("Suppression des données existantes...")
			if _, err := collection.DeleteMany(nil, map[string]interface{}{}); err != nil {
				logger.Wf("Erreur lors du nettoyage : %v", err)
			} else {
				logger.If("Collection vidée avec succès")
			}
		}

		// bulk
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

		totalCertificationString := fmt.Sprintf("Total certifications insérées : %v", len(result.InsertedIDs))
		if err := fileManagerLog.AppendLog(totalCertificationString, nil); err != nil {
			logger.Wf("%v", err)
			return
		}
		logger.If("%v", totalCertificationString)
	},
}

func init() {
	RncpImportDataInDatabaseCmd.Flags().BoolP("clear", "c", false, "Vider la collection avant import")
}
