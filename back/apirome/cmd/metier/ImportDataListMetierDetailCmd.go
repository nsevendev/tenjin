package metier

import (
	"encoding/json"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/apirome/libs"
	"tenjin/back/internal/utils/db"
)

var ImportDataListMetiersDetailCmd = &cobra.Command{
	Use:   "import-metier-detail",
	Short: "Importe tous les métiers détaillés en base MongoDB",
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnexionDatabase("dev")
		collection := db.Client.Collection("metiers")

		// 1. Ouvre le fichier JSON généré
		file := libs.GetLastFileByDate("./apirome/version_4_0/data/listmetierdetail")
		fileData, errFileData := os.Open(file)
		if errFileData != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, errFileData)
		}
		defer fileData.Close()

		// 2. Parse le contenu en slice de structs
		var metiers []map[string]interface{} // adapte au nom réel
		if err := json.NewDecoder(fileData).Decode(&metiers); err != nil {
			logger.Ff("Erreur decoding JSON : %v", err)
		}
		logger.If("Nombre de métiers à importer : %d", len(metiers))

		// 3. Import en bulk (ou un par un)
		var docs []interface{}
		for _, metier := range metiers {
			docs = append(docs, metier)
		}
		result, err := collection.InsertMany(nil, docs)
		if err != nil {
			logger.Ff("Erreur insertion Mongo : %v", err)
		}
		logger.If("Import réussi, %d métiers insérés", len(result.InsertedIDs))
	},
}
