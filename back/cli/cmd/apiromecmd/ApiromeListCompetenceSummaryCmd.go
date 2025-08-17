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

// ApiromeListCompetenceSummaryCmd est la commande pour télécharger la liste des compétences réduit et l'enregistrer en JSON dans un fichier
// Ce fichier créé sera utiliser pour telechager les détails de chaque compétence
// chaque rentrées contient un code, c'est ce code qui sera utilisé pour récuperer les détails de chaque compétence
// c'est une commande assez rapide
var ApiromeListCompetenceSummaryCmd = &cobra.Command{
	Use:   "list-competence-summary",
	Short: "Recuperer toutes les compétences résumé",
	Run: func(cmd *cobra.Command, args []string) {
		apiromeAuth := apirome.NewAuth()
		httpclient := cliutils.NewHTTPClient(60 * time.Second)
		fileManagerLog := cliutils.NewFileManager("./cli/data/apirome/log", cliutils.FileTypeLOG)
		fileManager := cliutils.NewFileManager("./cli/data/apirome/list-competence-summary", cliutils.FileTypeJSON)

		token, err := apiromeAuth.GetToken()
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		req := apirome.RequestGetListCompetence(token)
		bodyListCompetence, err := httpclient.ExecuteRequest(req)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		var competences []apirome.CompetenceSummary
		if err := json.Unmarshal(bodyListCompetence, &competences); err != nil {
			logger.Wf("Impossible de parser les statistiques: %v", err)
		} else {
			totalCompetenceString := fmt.Sprintf("Total competences résumé enregistrés en json: %v", len(competences))
			if err := fileManagerLog.AppendLog(totalCompetenceString, nil); err != nil {
				logger.Wf("%v", err)
				return
			}
			logger.If(totalCompetenceString)
		}

		if _, err := fileManager.WriteData(json.RawMessage(bodyListCompetence), nil); err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.I("Liste des competences résumé téléchargée et enregistrée avec succès")
	},
}
