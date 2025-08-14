package metier

import (
	"github.com/spf13/cobra"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version_4_0/metiers_4_0"
)

// SyncAndWriteInFileListMetierSummary est la commande pour télécharger le résumé des métiers et l'enregistrer en JSON dans un fichier
// Ce fichier créé sera utiliser pour telechager les details de chaque métier
// chaque rentrées contient un code et un libellé c'est ce code qui sera utilisé pour récuperer les détails de chaque métier
// c'est une commande assez rapide
var SyncAndWriteInFileListMetierSummary = &cobra.Command{
	Use:   "list-metier-summary",
	Short: "Télécharge le résumé des métiers ('code', 'libelle') et l'enregistre en JSON dans un fichier",
	Run: func(cmd *cobra.Command, args []string) {
		token := libs.GetToken()

		req := metiers_4_0.RequestGetListMetier(token)

		bodyListMetierSummary := libs.ExecuteRequest(req)

		libs.PrintBrutInFile("./apirome/version_4_0/data/listmetiersummary", bodyListMetierSummary)
	},
}
