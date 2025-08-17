package apiromecmd

import (
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"tenjin/back/cli/cliutils"
	"tenjin/back/cli/internal/apirome"
	"time"
)

type errorCompetence struct {
	Code      string    `json:"code"`
	Libelle   string    `json:"libelle"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

// ApiromeListCompetenceCompletCmd est la commande pour télécharger les détails complets de chaque compétence et les enregistrer en JSON dans un fichier
// Cette commande est plus longue car elle télécharge les détails de chaque compétence un par un avec un délai de 2 secondes entre chaque requête
// pour respecter la contrainte du serveur API france travail
var ApiromeListCompetenceCompletCmd = &cobra.Command{
	Use:   "list-competence-complet",
	Short: "Télécharge les détails complets de chaque compétence et les enregistre dans un fichier JSON",
	Run: func(cmd *cobra.Command, args []string) {
		apiromeAuth := apirome.NewAuth()
		httpclient := cliutils.NewHTTPClient(60 * time.Second)
		fileManager := cliutils.NewFileManager("./cli/data/apirome/list-competence-complet", cliutils.FileTypeJSON)
		fileManagerError := cliutils.NewFileManager("./cli/data/apirome/log/error", cliutils.FileTypeLOG)
		successCount := 0
		errorCount := 0
		var allBodies []json.RawMessage
		var errors []errorCompetence
		var listCompetence []apirome.CompetenceSummary

		competenceJsonToStruct(&listCompetence)
		lenListCompetence := len(listCompetence)

		logger.If("Début du téléchargement de %d compétences...", lenListCompetence)

		for i, competence := range listCompetence {
			token, err := apiromeAuth.GetToken()
			if err != nil {
				logger.Ff("%v", err)
			}
			// contrainte du serveur api, limite de temops par requête
			time.Sleep(2 * time.Second)

			logger.If("Progression: %d/%d - Traitement de la compétence: %v", i+1, lenListCompetence, competence.Code)

			req := apirome.RequestGetOneCompetenceComplet(token, competence.Code)
			if req == nil {
				errMessage := fmt.Sprintf("%v", err)
				addErrorEntry(&errors, &errorCount, competence.Code, competence.Libelle, errMessage)
				logger.Ef("Erreur requête pour %v: %v", competence.Code, errMessage)
				continue
			}

			body, err := httpclient.ExecuteRequest(req)
			if err != nil {
				errMessage := fmt.Sprintf("%v", err)
				addErrorEntry(&errors, &errorCount, competence.Code, competence.Libelle, errMessage)
				logger.Ef("Erreur execution pour %v - erreur: %v", competence.Code, errMessage)
				continue
			}

			// vérifie le contenu avant JSON
			bodyStr := string(body)
			if strings.HasPrefix(bodyStr, "<") { // C'est du HTML, pas du JSON !
				errMessage := fmt.Sprintf("API retourne du HTML: %v", bodyStr[:min(200, len(bodyStr))])
				addErrorEntry(&errors, &errorCount, competence.Code, competence.Libelle, errMessage)
				logger.Ef("API retourne du HTML pour %v: %v", competence.Code, errMessage)
				continue
			}

			// test si body est du json
			var testJson json.RawMessage
			if err := json.Unmarshal(body, &testJson); err != nil {
				errMessage := fmt.Sprintf("réponse invalide (pas de JSON valide): %v", err)
				addErrorEntry(&errors, &errorCount, competence.Code, competence.Libelle, errMessage)
				logger.Ef("JSON invalide pour %v: %v", competence.Code, errMessage)
				continue
			}

			allBodies = append(allBodies, body)
			successCount++
			logger.If("Competence complete bien ajoutée : %s", competence.Code)
		}

		// Sauvegarder les erreurs s'il y en a
		if len(errors) > 0 {
			if _, err := fileManagerError.WriteData(errors, nil); err != nil {
				logger.Ef("Erreur sauvegarde fichier erreurs: %v", err)
			} else {
				logger.Wf("Fichier d'erreurs sauvegardé avec succès")
			}
		}

		// Sauvegarder tous les body dans un fichier JSON
		if len(allBodies) > 0 {
			logger.If("Sauvegarde de %d compétences réussies...", len(allBodies))
			if _, err := fileManager.WriteData(allBodies, nil); err != nil {
				logger.Ef("Erreur sauvegarde fichier compétences: %v", err)
			} else {
				logger.If("Fichier des compétences sauvegardé avec succès")
			}
		} else {
			logger.Wf("Aucune compétence valide à sauvegarder")
		}

		// Statistiques finales
		logger.If("=== RÉSUMÉ ===")
		logger.If("Total traité: %d", len(listCompetence))
		logger.If("Succès: %d", successCount)
		logger.If("Erreurs: %d", errorCount)
		if len(listCompetence) > 0 {
			successRate := float64(successCount) / float64(len(listCompetence)) * 100
			logger.If("Taux de succès: %.2f%%", successRate)
		}

		if errorCount > 0 {
			logger.Wf("%d compétences ont échoué - voir le fichier d'erreurs", errorCount)
		} else {
			logger.If("Toutes les compétences ont été traitées avec succès !")
		}
	},
}

// competenceJsonToStruct lit le fichier JSON des compétences et le convertit en slice de struct competence
// Le fichier doit être au format JSON valide et contenir un tableau d'objets compétence
func competenceJsonToStruct(listCompetence *[]apirome.CompetenceSummary) {
	fileExplorer := cliutils.NewFileExplorer("./cli/data/apirome/list-competence-summary")
	file, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtJSON)
	if err != nil {
		logger.Ef("erreur lors de la récupération du dernier fichier de compétences : %v", err)
		return
	}
	fileData, err := os.Open(file)
	if err != nil {
		logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, err)
	}
	defer fileData.Close()

	err = json.NewDecoder(fileData).Decode(&listCompetence)
	if err != nil {
		logger.Ff("erreur lors du decodage json du fichier %s : %v", file, err)
	}
}

// addErrorEntry ajoute une entrée d'erreur à la slice d'erreurs et incrémente le compteur d'erreurs
func addErrorEntry(e *[]errorCompetence, errorCount *int, code string, libelle string, errorMsg string) {
	eEntry := errorCompetence{
		Code:      code,
		Libelle:   libelle,
		Error:     errorMsg,
		Timestamp: time.Now(),
	}
	*e = append(*e, eEntry)
	*errorCount++
}
