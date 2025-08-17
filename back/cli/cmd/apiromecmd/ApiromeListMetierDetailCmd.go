package apiromecmd

import (
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"os"
	"tenjin/back/cli/cliutils"
	"tenjin/back/cli/internal/apirome"
	"time"
)

type errorEntry struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// ApiromeListMetierDetailCmd est la commande pour télécharger les détails de chaque métier et les enregistrer en JSON dans un fichier
// Cette commande est plus longue car elle télécharge les détails de chaque métier un par un avec un delais de 2 secondes entre chaque requête
// pour respecter la contrainte du serveur API france travail
var ApiromeListMetierDetailCmd = &cobra.Command{
	Use:   "list-metier-detail",
	Short: "Télécharge metier par metier avec leur detail et enregistre en JSON dans un fichier (long)",
	Run: func(cmd *cobra.Command, args []string) {
		successCount := 0
		errorCount := 0
		apiromeAuth := apirome.NewAuth()
		fileExplorer := cliutils.NewFileExplorer("./cli/data/apirome/list-metier-summary")
		fileManager := cliutils.NewFileManager("./cli/data/apirome/list-metier-detail", cliutils.FileTypeJSON)
		fileManagerLog := cliutils.NewFileManager("./cli/data/apirome/log", cliutils.FileTypeLOG)
		fileManagerError := cliutils.NewFileManager("./cli/data/apirome/log/error", cliutils.FileTypeLOG)
		httpclient := cliutils.NewHTTPClient(60 * time.Second)
		var errorsEntry []errorEntry

		file, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtJSON)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		fileData, err := os.Open(file)
		if err != nil {
			logger.Ff("erreur lors de l'ouverture du fichier %s : %v", file, err)
			return
		}
		defer fileData.Close()

		var listMetier []apirome.MetierSummary
		err = json.NewDecoder(fileData).Decode(&listMetier)
		if err != nil {
			logger.Ff("erreur lors du decodage json du fichier %s : %v", file, err)
			return
		}

		var allBodies []json.RawMessage
		for _, metier := range listMetier {
			token, err := apiromeAuth.GetToken()
			if err != nil {
				logger.Ff("%v", err)
				return
			}
			time.Sleep(2 * time.Second) // contrainte du serveur api, limite de temops par requête
			req, err := apirome.RequestGetOneMetier(token, metier.Code)
			if err != nil {
				errString := fmt.Sprintf("erreur lors de la création de la requête pour le métier %v : %v", metier.Code, err)
				logger.Ef("%v", errString)
				errorCount++
				errorsEntry = append(errorsEntry, errorEntry{
					Code:      metier.Code,
					Message:   errString,
					Timestamp: time.Now(),
				})
				continue
			}
			body, err := httpclient.ExecuteRequest(req)
			if err != nil {
				errString := fmt.Sprintf("erreur lors de l'appel pour le métier %v : %v", metier.Code, err)
				logger.Ef("%v", errString)
				errorCount++
				errorsEntry = append(errorsEntry, errorEntry{
					Code:      metier.Code,
					Message:   errString,
					Timestamp: time.Now(),
				})
				continue
			}

			// on ajoute le body en cours dans le tableau allBodies
			allBodies = append(allBodies, json.RawMessage(body))
			logger.If("MetierSummary bien ajouter : %s", metier.Code)
			successCount++
		}

		metierTraitesString := fmt.Sprintf("Total métiers traités: %d, succès: %d, erreurs: %d", len(listMetier), successCount, errorCount)
		bodieString := fmt.Sprintf("Total métiers avec détails récupérés: %d", len(allBodies))

		if len(errorsEntry) > 0 {
			errorFile, err := fileManagerError.WriteErrors(errorsEntry)
			if err != nil {
				logger.Ef("%v", err)
			} else {
				logger.If("Fichier d'erreurs créé: %s", errorFile)
			}
		}

		if err := fileManagerLog.AppendLog(metierTraitesString, nil); err != nil {
			logger.Wf("Erreur lors de l'écriture du log: %v", err)
		}
		if err := fileManagerLog.AppendLog(bodieString, nil); err != nil {
			logger.Wf("Erreur lors de l'écriture du log: %v", err)
		}
		if _, err := fileManager.WriteData(allBodies, nil); err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.If("%v", metierTraitesString)
		logger.If("%v", bodieString)
	},
}
