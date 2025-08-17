package apiromecmd

import (
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/cli/cliutils"
	"tenjin/back/internal/utils/database"
)

// ApiromeImportListMetiersDetailInDatabaseCmd est la commande pour importer les métiers détaillés en JSON dans la base de données MongoDB
// Cette commande lit un fichier JSON contenant les détails des métiers et les insère dans la collection "metiers"
var ApiromeImportListMetiersDetailInDatabaseCmd = &cobra.Command{
	Use:   "import-metier-detail-database",
	Short: "Importe tous les métiers détaillés en json dans la database",
	Run: func(cmd *cobra.Command, args []string) {
		clearExisting, _ := cmd.Flags().GetBool("clear")
		database.ConnexionDatabase("dev")
		collection := database.Client.Collection("metiers")
		fileExplorer := cliutils.NewFileExplorer("./cli/data/apirome/list-metier-detail")
		fileManagerLog := cliutils.NewFileManager("./cli/data/apirome/log", cliutils.FileTypeLOG)

		file, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtJSON)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
			return
		}
		defer fileData.Close()

		var metiers []map[string]interface{}
		if err := json.NewDecoder(fileData).Decode(&metiers); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
			return
		}
		logger.If("Nombre de métiers à importer : %d", len(metiers))

		if clearExisting {
			logger.If("Suppression des données existantes...")
			if deleteResult, err := collection.DeleteMany(nil, map[string]interface{}{}); err != nil {
				logger.Wf("Erreur lors du nettoyage : %v", err)
			} else {
				logger.If("Collection vidée : %d documents supprimés", deleteResult.DeletedCount)
			}
		}

		var docs []interface{}
		for _, metier := range metiers {
			docs = append(docs, metier)
		}

		logger.If("Import en cours...")
		result, err := collection.InsertMany(nil, docs)
		if err != nil {
			logger.Ff("Erreur insertion Mongo : %v", err)
		}

		totalMetier := fmt.Sprintf("Total metiers insérées : %v", len(result.InsertedIDs))
		if err := fileManagerLog.AppendLog(totalMetier, nil); err != nil {
			logger.Wf("%v", err)
			return
		}
		logger.If(totalMetier)
	},
}

func init() {
	ApiromeImportListMetiersDetailInDatabaseCmd.Flags().BoolP("clear", "c", false, "Vider la collection avant import")
}
