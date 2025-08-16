package rncp

import (
	"encoding/xml"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type XMLFiches struct {
	XMLName     xml.Name   `xml:"FICHES"`
	VersionFlux string     `xml:"VERSION_FLUX"`
	Fiches      []XMLFiche `xml:"FICHE"`
}

type XMLFiche struct {
	XMLName                  xml.Name              `xml:"FICHE"`
	IDFiche                  string                `xml:"ID_FICHE"`
	NumeroFiche              string                `xml:"NUMERO_FICHE"`
	Intitule                 string                `xml:"INTITULE"`
	Abrege                   XMLAbrege             `xml:"ABREGE"`
	EtatFiche                string                `xml:"ETAT_FICHE"`
	NomenclatureEurope       XMLNomenclatureEurope `xml:"NOMENCLATURE_EUROPE"`
	CodesNSF                 []XMLNSF              `xml:"CODES_NSF>NSF"`
	Certificateurs           []XMLCertificateur    `xml:"CERTIFICATEURS>CERTIFICATEUR"`
	ExistencePartenaires     string                `xml:"EXISTENCE_PARTENAIRES"`
	ActivitesVisees          string                `xml:"ACTIVITES_VISEES"`
	CapacitesAttestees       string                `xml:"CAPACITES_ATTESTEES"`
	SecteursActivite         string                `xml:"SECTEURS_ACTIVITE"`
	TypeEmploiAccessibles    string                `xml:"TYPE_EMPLOI_ACCESSIBLES"`
	CodesROME                []XMLCodeROME         `xml:"CODES_ROME>ROME"`
	ReglementationsActivites string                `xml:"REGLEMENTATIONS_ACTIVITES"`
	SiJuryFI                 XMLJury               `xml:"SI_JURY_FI"`
	SiJuryCA                 XMLJury               `xml:"SI_JURY_CA"`
	SiJuryFC                 XMLJury               `xml:"SI_JURY_FC"`
	SiJuryCQ                 XMLJury               `xml:"SI_JURY_CQ"`
	SiJuryCL                 XMLJury               `xml:"SI_JURY_CL"`
	SiJuryVAE                XMLJury               `xml:"SI_JURY_VAE"`
	AccessibleNouvCaledonie  string                `xml:"ACCESSIBLE_NOUVELLE_CALEDONIE"`
	AccessiblePolynesie      string                `xml:"ACCESSIBLE_POLYNESIE_FRANCAISE"`
	TypeEnregistrement       string                `xml:"TYPE_ENREGISTREMENT"`
	ObjectifsContexte        string                `xml:"OBJECTIFS_CONTEXTE"`
	Actif                    string                `xml:"ACTIF"`
	PrerequisEntreeFormation string                `xml:"PREREQUIS_ENTREE_FORMATION"`
}

type XMLAbrege struct {
	Code    string `xml:"CODE"`
	Libelle string `xml:"LIBELLE"`
}

type XMLNomenclatureEurope struct {
	Niveau  string `xml:"NIVEAU"`
	Libelle string `xml:"LIBELLE"`
}

type XMLNSF struct {
	Code    string `xml:"CODE"`
	Libelle string `xml:"LIBELLE"`
}

type XMLCertificateur struct {
	NomCertificateur  string `xml:"NOM_CERTIFICATEUR"`
	EtatCertificateur string `xml:"ETAT_CERTIFICATEUR"`
	SiteInternet      string `xml:"SITE_INTERNET"`
}

type XMLCodeROME struct {
	Code    string `xml:"CODE"`
	Libelle string `xml:"LIBELLE"`
}

type XMLJury struct {
	Actif       string `xml:"ACTIF"`
	Composition string `xml:"COMPOSITION"`
}

// Structure finale pour MongoDB (format simplifié et optimisé)
type Certification struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDFiche               string             `bson:"id_fiche" json:"id_fiche"`
	NumeroFiche           string             `bson:"numero_fiche" json:"numero_fiche"`
	Intitule              string             `bson:"intitule" json:"intitule"`
	Abrege                Abrege             `bson:"abrege" json:"abrege"`
	EtatFiche             string             `bson:"etat_fiche" json:"etat_fiche"`
	Niveau                string             `bson:"niveau" json:"niveau"`
	NiveauLibelle         string             `bson:"niveau_libelle" json:"niveau_libelle"`
	CodesNSF              []NSF              `bson:"codes_nsf" json:"codes_nsf"`
	Certificateurs        []Certificateur    `bson:"certificateurs" json:"certificateurs"`
	ActivitesVisees       string             `bson:"activites_visees" json:"activites_visees"`
	CapacitesAttestees    string             `bson:"capacites_attestees" json:"capacites_attestees"`
	SecteursActivite      string             `bson:"secteurs_activite" json:"secteurs_activite"`
	TypeEmploiAccessibles string             `bson:"type_emploi_accessibles" json:"type_emploi_accessibles"`
	CodesROME             []CodeROME         `bson:"codes_rome" json:"codes_rome"`
	Actif                 bool               `bson:"actif" json:"actif"`
	ImportedAt            time.Time          `bson:"imported_at" json:"imported_at"`
	UpdatedAt             time.Time          `bson:"updated_at" json:"updated_at"`
}

type Abrege struct {
	Code    string `bson:"code" json:"code"`
	Libelle string `bson:"libelle" json:"libelle"`
}

type NSF struct {
	Code    string `bson:"code" json:"code"`
	Libelle string `bson:"libelle" json:"libelle"`
}

type Certificateur struct {
	Nom          string `bson:"nom" json:"nom"`
	Etat         string `bson:"etat" json:"etat"`
	SiteInternet string `bson:"site_internet" json:"site_internet"`
}

type CodeROME struct {
	Code    string `bson:"code" json:"code"`
	Libelle string `bson:"libelle" json:"libelle"`
}

// Fonction de conversion XML vers structure finale
func ConvertXMLToMongo(xmlFiche XMLFiche) Certification {
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
