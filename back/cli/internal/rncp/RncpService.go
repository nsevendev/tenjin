package rncp

import (
	"encoding/xml"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"os"
	"time"
)

func ReadAndParseXMLToStruct(xmlFile string, xmlFiches *XMLFiches) error {
	xmlData, err := os.ReadFile(xmlFile)
	if err != nil {
		return fmt.Errorf("erreur de la lecture du fichier XML : %v", err)
	}

	if err = xml.Unmarshal(xmlData, &xmlFiches); err != nil {
		return fmt.Errorf("erreur lors du parsing XML : %v", err)
	}

	logger.If("XML parsé avec succès - Version flux: %s", xmlFiches.VersionFlux)
	logger.If("Nombre de fiches trouvées : %d", len(xmlFiches.Fiches))

	return nil
}

func ConvertXMLStructToStructMongo(xmlFiches XMLFiches, certifications *[]Certification) {
	for _, xmlFiche := range xmlFiches.Fiches {
		cert := convertXMLToMongo(xmlFiche)
		*certifications = append(*certifications, cert)
	}
}

// ConvertXMLToMongo Fonction de conversion XML vers structure finale
func convertXMLToMongo(xmlFiche XMLFiche) Certification {
	// Conversion des codes NSF
	var codesNSF []NSF
	for _, nsf := range xmlFiche.CodesNSF {
		codesNSF = append(codesNSF, NSF{
			Code:    nsf.Code,
			Libelle: nsf.Libelle,
		})
	}

	// Conversion des certificateurs
	var certificateurs []Certificateur
	for _, cert := range xmlFiche.Certificateurs {
		certificateurs = append(certificateurs, Certificateur{
			Nom:          cert.NomCertificateur,
			Etat:         cert.EtatCertificateur,
			SiteInternet: cert.SiteInternet,
		})
	}

	// Conversion des codes ROME
	var codesROME []CodeROME
	for _, rome := range xmlFiche.CodesROME {
		codesROME = append(codesROME, CodeROME{
			Code:    rome.Code,
			Libelle: rome.Libelle,
		})
	}

	return Certification{
		IDFiche:     xmlFiche.IDFiche,
		NumeroFiche: xmlFiche.NumeroFiche,
		Intitule:    xmlFiche.Intitule,
		Abrege: Abrege{
			Code:    xmlFiche.Abrege.Code,
			Libelle: xmlFiche.Abrege.Libelle,
		},
		EtatFiche:             xmlFiche.EtatFiche,
		Niveau:                xmlFiche.NomenclatureEurope.Niveau,
		NiveauLibelle:         xmlFiche.NomenclatureEurope.Libelle,
		CodesNSF:              codesNSF,
		Certificateurs:        certificateurs,
		ActivitesVisees:       xmlFiche.ActivitesVisees,
		CapacitesAttestees:    xmlFiche.CapacitesAttestees,
		SecteursActivite:      xmlFiche.SecteursActivite,
		TypeEmploiAccessibles: xmlFiche.TypeEmploiAccessibles,
		CodesROME:             codesROME,
		Actif:                 xmlFiche.Actif == "Oui",
		ImportedAt:            time.Now(),
		UpdatedAt:             time.Now(),
	}
}
