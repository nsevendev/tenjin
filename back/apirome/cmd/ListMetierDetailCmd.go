package cmd

import (
	"encoding/json"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version40/metiers40"
	"time"
)

var SyncAndWriteInFileListMetierDetail = &cobra.Command{
	Use:   "list-metier-detail",
	Short: "Télécharge metier par metier avec leur detail, crée un tableau de metier et enregistre ce tableau en JSON dans un fichier",
	Run: func(cmd *cobra.Command, args []string) {
		type Metier struct {
			Code    string `json:"code"`
			Libelle string `json:"libelle"`
		}

		file := libs.GetLastFileByDate("./apirome/version40/data/listmetier")
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

			req := metiers40.RequestGetOneMetier(token, metier.Code)
			body := libs.ExecuteRequest(req)

			// on ajoute le body en cours dans le tableau allBodies
			allBodies = append(allBodies, json.RawMessage(body))
			logger.If("Metier bien ajouter : %s", metier.Code)
		}

		// Sauvegarder tous le tableau de body dans un fichier JSON
		libs.PrintSliceBrutInFile("./apirome/version40/data/metierdetail", allBodies)
		logger.If("Détails des métiers sauvegardés dans ./apirome/version40/data/metierdetail")
	},
}
