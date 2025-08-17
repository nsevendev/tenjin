package rncpcmd

import (
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/spf13/cobra"
	"tenjin/back/cli/cliutils"
	"tenjin/back/cli/internal/rncp"
)

// RncpXmlToJsonCmd est la commande pour convertir le fichier XML RNCP en JSON
// Elle lit le dernier fichier XML trouvé dans le répertoire spécifié, le parse, le convertit en structures MongoDB,
// et sauvegarde le résultat en JSON dans le même répertoire avec un nom de fichier basé sur la date et l'heure actuelle.
var RncpXmlToJsonCmd = &cobra.Command{
	Use:   "xml-to-json",
	Short: "Convertit le fichier XML RNCP en format JSON",
	Run: func(cmd *cobra.Command, args []string) {
		var xmlFiches rncp.XMLFiches
		var certifications []rncp.Certification

		fileDirRncp := "./cli/data/rncp/"
		fileExplorer := cliutils.NewFileExplorer(fileDirRncp + "data-xml")
		fileManager := cliutils.NewFileManager(fileDirRncp+"data-json", cliutils.FileTypeJSON)
		fileManagerLog := cliutils.NewFileManager(fileDirRncp+"log", cliutils.FileTypeLOG)

		logger.If("recupération du dernier fichier XML dans le répertoire : %s", fileDirRncp+"data-xml")
		xmlFile, err := fileExplorer.GetLastFileByDateAndType(".", cliutils.ExtXML)
		if err != nil {
			logger.Ff("%v", err)
			return
		}

		logger.If("fichier XML trouvé : %v", xmlFile)

		logger.If("début de la conversion XML → JSON du fichier : %v", xmlFile)
		if err = rncp.ReadAndParseXMLToStruct(xmlFile, &xmlFiches); err != nil {
			logger.Ff("%v", err)
			return
		}

		rncp.ConvertXMLStructToStructMongo(xmlFiches, &certifications)
		logger.I("conversion terminée avec succès")
		logger.If("nombre de fiches convertis : %v", len(certifications))

		if _, err = fileManager.WriteData(certifications, nil); err != nil {
			logger.Ff("%v", err)
			return
		}

		totalCertificationString := fmt.Sprintf("certifications enregistrer xml -> json : %v", len(certifications))
		if err := fileManagerLog.AppendLog(totalCertificationString, nil); err != nil {
			logger.Wf("%v", err)
			return
		}
		logger.If("%v", totalCertificationString)
	},
}
