package main

import (
	"fmt"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/cli/cmd/apiromecmd"
	"tenjin/back/cli/cmd/rncpcmd"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "CLI",
	Long:  "Cli Tenjin gestion de tâches",
}

func init() {
	appEnv := env.Get("APP_ENV")
	logger.Init(appEnv)
	rootCmd.AddCommand(rncpcmd.RncpCmd)
	rootCmd.AddCommand(apiromecmd.ApiromeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
