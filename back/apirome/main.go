package main

import (
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/apirome/cmd/competence"
	"tenjin/back/apirome/cmd/metier"
	"tenjin/back/apirome/cmd/rncpcmd"
)

func init() {
	appEnv := env.Get("APP_ENV")
	logger.Init(appEnv)
}

func main() {
	rootCmd := &cobra.Command{Use: "apirome"}
	rootCmd.AddCommand(metier.SyncAndWriteInFileListMetierSummary)
	rootCmd.AddCommand(metier.SyncAndWriteInFileListMetierDetail)
	rootCmd.AddCommand(metier.ImportListMetiersDetailInDatabaseCmd)
	rootCmd.AddCommand(rncpcmd.ConvertXMLToJSONCmd)
	rootCmd.AddCommand(rncpcmd.ImportDataRNCPCmd)
	rootCmd.AddCommand(competence.SyncAndWriteInFileListCompetence)
	rootCmd.AddCommand(competence.SyncAndWriteInFileListCompetenceComplet)
	rootCmd.AddCommand(competence.ImportListCompetenceCompletInDatabaseCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Ff("erreur lors de l'exécution de la commande : %v", err)
	}
}
