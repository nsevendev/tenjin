package competence

import (
	"encoding/json"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version_4_0/metiers_4_0"
	"time"
)

// SyncAndWriteInFileListCompetenceComplet est la commande pour télécharger les détails complets de chaque compétence et les enregistrer en JSON dans un fichier
// Cette commande est plus longue car elle télécharge les détails de chaque compétence un par un avec un délai de 2 secondes entre chaque requête
// pour respecter la contrainte du serveur API france travail
var SyncAndWriteInFileListCompetenceComplet = &cobra.Command{
	Use:   "list-competence-complet",
	Short: "Télécharge les détails complets de chaque compétence et les enregistre dans un fichier JSON",
	Run: func(cmd *cobra.Command, args []string) {
		type Competence struct {
			Type    string `json:"type"`
			Code    string `json:"code"`
			Libelle string `json:"libelle"`
		}

		file := libs.GetLastFileByDate("./apirome/version_4_0/data/listcompetence")
		fileData, err := os.Open(file)
		if err != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, err)
		}
		defer fileData.Close()

		var listCompetence []Competence
		err = json.NewDecoder(fileData).Decode(&listCompetence)
		if err != nil {
			logger.Ff("erreur lors du decodage json du fichier %s : %v", file, err)
		}

		var allBodies []json.RawMessage
		for _, competence := range listCompetence {
			token := libs.GetToken()
			time.Sleep(2 * time.Second) // contrainte du serveur api, limite de temops par requête

			req := metiers_4_0.RequestGetOneCompetenceComplet(token, competence.Code)
			body := libs.ExecuteRequest(req)

			// on ajoute le body en cours dans le tableau allBodies
			allBodies = append(allBodies, body)
			logger.If("Competence complete bien ajouter : %s", competence.Code)
		}

		// Sauvegarder tous le tableau de body dans un fichier JSON
		libs.PrintSliceBrutInFile("./apirome/version_4_0/data/listcompetencecomplet", allBodies)
	},
}
