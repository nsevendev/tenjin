package metier

import (
	"encoding/json"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version_4_0/metiers_4_0"
	"time"
)

// SyncAndWriteInFileListMetierDetail est la commande pour télécharger les détails de chaque métier et les enregistrer en JSON dans un fichier
// Cette commande est plus longue car elle télécharge les détails de chaque métier un par un avec un delais de 2 secondes entre chaque requête
// pour respecter la contrainte du serveur API france travail
var SyncAndWriteInFileListMetierDetail = &cobra.Command{
	Use:   "list-metier-detail",
	Short: "Télécharge metier par metier avec leur detail et enregistre en JSON dans un fichier (long)",
	Run: func(cmd *cobra.Command, args []string) {
		type Metier struct {
			Code    string `json:"code"`
			Libelle string `json:"libelle"`
		}

		file := libs.GetLastFileByDate("./apirome/version_4_0/data/listmetiersummary")
		fileData, err := os.Open(file)
		if err != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, err)
		}
		defer fileData.Close()

		var listMetier []Metier
		err = json.NewDecoder(fileData).Decode(&listMetier)
		if err != nil {
			logger.Ff("erreur lors du decodage json du fichier %s : %v", file, err)
		}

		var allBodies []json.RawMessage
		for _, metier := range listMetier {
			token := libs.GetToken()
			time.Sleep(2 * time.Second) // contrainte du serveur api, limite de temops par requête

			req := metiers_4_0.RequestGetOneMetier(token, metier.Code)
			body := libs.ExecuteRequest(req)

			// on ajoute le body en cours dans le tableau allBodies
			allBodies = append(allBodies, json.RawMessage(body))
			logger.If("Metier bien ajouter : %s", metier.Code)
		}

		// Sauvegarder tous le tableau de body dans un fichier JSON
		libs.PrintSliceBrutInFile("./apirome/version_4_0/data/listmetierdetail", allBodies)
	},
}
