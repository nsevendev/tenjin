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

// ApiromeImportListCompetenceCompletInDatabaseCmd est la commande pour importer les compétences détaillées en JSON dans la base de données MongoDB
// Cette commande lit un fichier JSON contenant les détails des compétences et les insère dans la collection
var ApiromeImportListCompetenceCompletInDatabaseCmd = &cobra.Command{
	Use:   "import-competence-complet-database",
	Short: "Importe toutes les compétences détaillées en JSON dans la base de données",
	Run: func(cmd *cobra.Command, args []string) {
		clearExisting, _ := cmd.Flags().GetBool("clear")
		fileExplorer := cliutils.NewFileExplorer("./cli/data/apirome/list-competence-complet")
		fileManager := cliutils.NewFileManager("./cli/data/apirome/log", cliutils.FileTypeLOG)
		database.ConnexionDatabase("dev")
		collection := database.Client.Collection("competences")

		file, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtJSON)
		if err != nil {
			logger.Ff("Erreur lors de la récupération du dernier fichier JSON : %v", err)
			return
		}
		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
		}
		defer fileData.Close()

		var competences []map[string]interface{} // adapte au nom réel
		if err := json.NewDecoder(fileData).Decode(&competences); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
			return
		}
		logger.If("Nombre de competence à importer : %d", len(competences))

		if clearExisting {
			logger.If("Suppression des données existantes...")
			if deleteResult, err := collection.DeleteMany(nil, map[string]interface{}{}); err != nil {
				logger.Wf("Erreur lors du nettoyage : %v", err)
			} else {
				logger.If("Collection vidée : %d documents supprimés", deleteResult.DeletedCount)
			}
		}

		// bulk
		var docs []interface{}
		for _, competence := range competences {
			docs = append(docs, competence)
		}
		result, err := collection.InsertMany(nil, docs)
		if err != nil {
			logger.Ff("Erreur insertion Mongo : %v", err)
			return
		}
		importSuccessString := fmt.Sprintf("Import réussi, %d competences insérés", len(result.InsertedIDs))
		if err := fileManager.AppendLog(importSuccessString, nil); err != nil {
			logger.Ff("Erreur lors de l'écriture du log : %v", err)
			return
		}
		logger.If(importSuccessString)
	},
}

func init() {
	ApiromeImportListCompetenceCompletInDatabaseCmd.Flags().BoolP("clear", "c", false, "Vider la collection avant import")
}
