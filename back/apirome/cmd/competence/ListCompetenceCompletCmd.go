package competence

import (
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"tenjin/back/apirome/libs"
	"tenjin/back/apirome/version_4_0/metiers_4_0"
	"time"
)

// SyncAndWriteInFileListCompetenceComplet est la commande pour télécharger les détails complets de chaque compétence et les enregistrer en JSON dans un fichier
// Cette commande est plus longue car elle télécharge les détails de chaque compétence un par un avec un délai de 2 secondes entre chaque requête
// pour respecter la contrainte du serveur API france travail
var SyncAndWriteInFileListCompetenceComplet = &cobra.Command{
	Use:   "list-competence-complet",
	Short: "Télécharge les détails complets de chaque compétence et les enregistre dans un fichier JSON",
	Run: func(cmd *cobra.Command, args []string) {
		type Competence struct {
			Type    string `json:"type"`
			Code    string `json:"code"`
			Libelle string `json:"libelle"`
		}

		type ErrorEntry struct {
			Code      string    `json:"code"`
			Libelle   string    `json:"libelle"`
			Error     string    `json:"error"`
			Timestamp time.Time `json:"timestamp"`
		}

		file := libs.GetLastFileByDate("./apirome/version_4_0/data/listcompetence")
		fileData, err := os.Open(file)
		if err != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, err)
		}
		defer fileData.Close()

		var listCompetence []Competence
		err = json.NewDecoder(fileData).Decode(&listCompetence)
		if err != nil {
			logger.Ff("erreur lors du decodage json du fichier %s : %v", file, err)
		}

		var allBodies []json.RawMessage
		var errors []ErrorEntry
		successCount := 0
		errorCount := 0

		logger.If("Début du téléchargement de %d compétences...", len(listCompetence))

		for i, competence := range listCompetence {
			token := libs.GetToken()
			time.Sleep(2 * time.Second) // contrainte du serveur api, limite de temops par requête

			logger.If("Progression: %d/%d - Traitement de la compétence: %s", i+1, len(listCompetence), competence.Code)

			req := metiers_4_0.RequestGetOneCompetenceComplet(token, competence.Code)
			if req == nil {
				errorEntry := ErrorEntry{
					Code:      competence.Code,
					Libelle:   competence.Libelle,
					Error:     "erreur lors de la création de la requête HTTP",
					Timestamp: time.Now(),
				}
				errors = append(errors, errorEntry)
				errorCount++
				logger.Ef("Erreur requête pour %s: %s", competence.Code, errorEntry.Error)
				continue
			}

			body := libs.ExecuteRequest(req)
			if body == nil {
				errorEntry := ErrorEntry{
					Code:      competence.Code,
					Libelle:   competence.Libelle,
					Error:     "erreur lors de l'exécution de la requête (voir logs pour détails)",
					Timestamp: time.Now(),
				}
				errors = append(errors, errorEntry)
				errorCount++
				logger.Ef("Erreur execution pour %s", competence.Code)
				continue
			}

			// vérifie le contenu avant JSON
			bodyStr := string(body)
			if strings.HasPrefix(bodyStr, "<") {
				// C'est du HTML, pas du JSON !
				errorEntry := ErrorEntry{
					Code:      competence.Code,
					Libelle:   competence.Libelle,
					Error:     fmt.Sprintf("API retourne du HTML: %s", bodyStr[:min(200, len(bodyStr))]),
					Timestamp: time.Now(),
				}
				errors = append(errors, errorEntry)
				errorCount++
				logger.Ef("Réponse HTML pour %s: %s", competence.Code, bodyStr[:100])
				continue
			}

			// test si body est du json
			var testJson json.RawMessage
			if err := json.Unmarshal(body, &testJson); err != nil {
				errorEntry := ErrorEntry{
					Code:      competence.Code,
					Libelle:   competence.Libelle,
					Error:     fmt.Sprintf("réponse invalide (pas de JSON valide): %v", err),
					Timestamp: time.Now(),
				}
				errors = append(errors, errorEntry)
				errorCount++
				logger.Ef("JSON invalide pour %s: %v", competence.Code, err)
				continue
			}

			allBodies = append(allBodies, body)
			successCount++
			logger.If("Competence complete bien ajoutée : %s", competence.Code)

		}

		// Sauvegarder les erreurs s'il y en a
		if len(errors) > 0 {
			timestamp := time.Now().Format("20060102_150405")
			errorFileName := fmt.Sprintf("%s_erreur.json", timestamp)
			errorFilePath := filepath.Join("./apirome/version_4_0/data/listcompetencecomplet", errorFileName)

			// Créer le répertoire des erreurs s'il n'existe pas
			if err := os.MkdirAll(filepath.Dir(errorFilePath), 0755); err != nil {
				logger.Ef("Erreur création répertoire erreurs: %v", err)
			} else {
				errorData, err := json.MarshalIndent(errors, "", "  ")
				if err != nil {
					logger.Ef("Erreur sérialisation des erreurs: %v", err)
				} else {
					if err := os.WriteFile(errorFilePath, errorData, 0644); err != nil {
						logger.Ef("Erreur écriture fichier erreurs: %v", err)
					} else {
						logger.Wf("Fichier d'erreurs sauvegardé: %s", errorFilePath)
					}
				}
			}
		}

		// Sauvegarder tous les body dans un fichier JSON
		if len(allBodies) > 0 {
			logger.If("Sauvegarde de %d compétences réussies...", len(allBodies))
			libs.PrintSliceBrutInFile("./apirome/version_4_0/data/listcompetencecomplet", allBodies)
			logger.If("Fichier des compétences sauvegardé avec succès")
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
