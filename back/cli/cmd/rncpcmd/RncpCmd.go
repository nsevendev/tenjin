package rncpcmd

import "github.com/spf13/cobra"

// RncpCmd groupe commande rncp
var RncpCmd = &cobra.Command{
	Use:   "rncp",
	Short: "Gestion des rncps",
	Long:  "Commandes pour la gestion des rncps",
}

func init() {
	RncpCmd.AddCommand(RncpXmlToJsonCmd)
	RncpCmd.AddCommand(RncpImportDataInDatabaseCmd)
}
