package main

import (
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/apirome/cmd"
)

func init() {
	appEnv := env.Get("APP_ENV")
	logger.Init(appEnv)
}

func main() {
	rootCmd := &cobra.Command{Use: "apirome"}
	rootCmd.AddCommand(cmd.GenerateModelCmd)
	rootCmd.AddCommand(cmd.SyncAndWriteInFileListMetierSummary)
	rootCmd.AddCommand(cmd.SyncAndWriteInFileListMetierDetail)

	if err := rootCmd.Execute(); err != nil {
		logger.Ff("erreur lors de l'exécution de la commande : %v", err)
	}
}
