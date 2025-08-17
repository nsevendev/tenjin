package apiromecmd

import (
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/cli/cliutils"
	"tenjin/back/cli/internal/apirome"
	"time"
)

// ApiromeListMetierSummaryCmd est la commande pour télécharger le résumé des métiers et l'enregistrer en JSON dans un fichier
// Ce fichier créé sera utiliser pour telechager les details de chaque métier
// chaque rentrées contient un code et un libellé c'est ce code qui sera utilisé pour récuperer les détails de chaque métier
// c'est une commande assez rapide
var ApiromeListMetierSummaryCmd = &cobra.Command{
	Use:   "list-metier-summary",
	Short: "Télécharge la liste des métiers résumé et enregistre en JSON dans un fichier",
	Run: func(cmd *cobra.Command, args []string) {
		apiromeAuth := apirome.NewAuth()
		fileManager := cliutils.NewFileManager("./cli/data/apirome/list-metier-summary", cliutils.FileTypeJSON)
		fileManagerLog := cliutils.NewFileManager("./cli/data/apirome/log", cliutils.FileTypeLOG)
		httpclient := cliutils.NewHTTPClient(60 * time.Second)

		token, err := apiromeAuth.GetToken()
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.If("Token obtenu avec succès")

		req := apirome.RequestGetListMetier(token)
		bodyListMetierSummary, err := httpclient.ExecuteRequest(req)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		var metiers []apirome.CompetenceSummary
		if err := json.Unmarshal(bodyListMetierSummary, &metiers); err != nil {
			logger.Wf("Impossible de parser les statistiques: %v", err)
		} else {
			totalMetierString := fmt.Sprintf("Total métiers résumé enregistrés en json: %v", len(metiers))
			if err := fileManagerLog.AppendLog(totalMetierString, nil); err != nil {
				logger.Wf("%v", err)
				return
			}
			logger.If("%v", totalMetierString)
		}

		if _, err := fileManager.WriteData(json.RawMessage(bodyListMetierSummary), nil); err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.I("Liste des métiers résumé téléchargée et enregistrée avec succès")
	},
}
