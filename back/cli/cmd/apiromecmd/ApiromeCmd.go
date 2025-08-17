package apiromecmd

import "github.com/spf13/cobra"

// ApiromeCmd groupe commande rncp
var ApiromeCmd = &cobra.Command{
	Use:   "apirome",
	Short: "Gestion des api romes",
	Long:  "Commandes pour la gestion des api romes de france travail",
}

func init() {
	ApiromeCmd.AddCommand(ApiromeListMetierSummaryCmd)
	ApiromeCmd.AddCommand(ApiromeListMetierDetailCmd)
	ApiromeCmd.AddCommand(ApiromeImportListMetiersDetailInDatabaseCmd)
	ApiromeCmd.AddCommand(ApiromeListCompetenceSummaryCmd)
	ApiromeCmd.AddCommand(ApiromeListCompetenceCompletCmd)
	ApiromeCmd.AddCommand(ApiromeImportListCompetenceCompletInDatabaseCmd)
}
