package competence

import (
	"github.com/spf13/cobra"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version_4_0/metiers_4_0"
)

// SyncAndWriteInFileListCompetence est la commande pour télécharger la liste des compétences réduit et l'enregistrer en JSON dans un fichier
// Ce fichier créé sera utiliser pour telechager les détails de chaque compétence
// chaque rentrées contient un code, c'est ce code qui sera utilisé pour récuperer les détails de chaque compétence
// c'est une commande assez rapide
var SyncAndWriteInFileListCompetence = &cobra.Command{
	Use:   "list-competence",
	Short: "Recuperer toutes les compétences",
	Run: func(cmd *cobra.Command, args []string) {
		token := libs.GetToken()

		req := metiers_4_0.RequestGetListCompetence(token)

		bodyListCompetence := libs.ExecuteRequest(req)

		libs.PrintBrutInFile("./apirome/version_4_0/data/listcompetence", bodyListCompetence)
	},
}
