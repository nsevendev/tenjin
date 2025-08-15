package competence

import (
	"encoding/json"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/apirome/libs"
	"tenjin/back/internal/utils/db"
)

// ImportListCompetenceCompletInDatabaseCmd est la commande pour importer les compétences détaillées en JSON dans la base de données MongoDB
// Cette commande lit un fichier JSON contenant les détails des compétences et les insère dans la collection
var ImportListCompetenceCompletInDatabaseCmd = &cobra.Command{
	Use:   "import-competence-complet-database",
	Short: "Importe toutes les compétences détaillées en JSON dans la base de données",
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnexionDatabase("dev")
		collection := db.Client.Collection("competences")

		// 1. Ouvre le fichier JSON généré
		file := libs.GetLastFileByDate("./apirome/version_4_0/data/listcompetencecomplet")
		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
		}
		defer fileData.Close()

		// 2. Parse le contenu en slice de structs
		var competences []map[string]interface{} // adapte au nom réel
		if err := json.NewDecoder(fileData).Decode(&competences); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
		}
		logger.If("Nombre de competence à importer : %d", len(competences))

		// 3. Import en bulk (ou un par un)
		var docs []interface{}
		for _, competence := range competences {
			docs = append(docs, competence)
		}
		result, err := collection.InsertMany(nil, docs)
		if err != nil {
			logger.Ff("Erreur insertion Mongo : %v", err)
		}
		logger.If("Import réussi, %d competences insérés", len(result.InsertedIDs))
	},
}
