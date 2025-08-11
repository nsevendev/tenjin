package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var inputPath string
var outputPath string
var goPackage string

var GenerateModelCmd = &cobra.Command{
	Use:   "generate-model",
	Short: "Génère les structs Go à partir d'un fichier OpenAPI (oapi-codegen)",
	Run: func(cmd *cobra.Command, args []string) {
		if inputPath == "" || outputPath == "" || goPackage == "" {
			log.Fatalf("Tous les paramètres --input, --output, --package sont obligatoires.")
		}
		cmdShell := exec.Command(
			"oapi-codegen",
			"-generate", "types",
			"-o", outputPath,
			"-package", goPackage,
			inputPath,
		)
		output, err := cmdShell.CombinedOutput()
		if err != nil {
			log.Fatalf("Erreur exec : %v\nSortie : %s", err, string(output))
		}
		log.Printf("Sortie :\n%s", string(output))
	},
}

func init() {
	GenerateModelCmd.Flags().StringVar(&inputPath, "input", "", "Chemin du fichier OpenAPI (json/yaml)")
	GenerateModelCmd.Flags().StringVar(&outputPath, "output", "", "Chemin de sortie du fichier Go généré")
	GenerateModelCmd.Flags().StringVar(&goPackage, "package", "", "Nom du package Go")
}
