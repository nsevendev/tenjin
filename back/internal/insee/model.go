package insee

import (
	"tenjin/back/internal/addresses"
	"tenjin/back/internal/utils/constantes"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type CompanyInfo struct {
    BusinessName            string `json:"business_name"`
    Siret                   string `json:"siret"`
    Addresses               []addresses.Address `json:"adresses"`
	Status                	string `json:"status"`
	Type                	constantes.TypeInstitute `json:"type"`
}

type sireneResponse struct {
	Etablissement sireneEtablissement `json:"etablissement"`
}

type sireneEtablissement struct {
	Siret                  string                     `json:"siret"`
	UniteLegale            sireneUniteLegale          `json:"uniteLegale"`
	AdresseEtablissement   sireneAdresseEtablissement `json:"adresseEtablissement"`
	Adresse2Etablissement  *sireneAdresseEtablissement `json:"adresse2Etablissement"`
	Enseigne1Etablissement string                     `json:"enseigne1Etablissement"`
}

type sireneUniteLegale struct {
	Siren                          string `json:"siren"`
	DenominationUniteLegale        string `json:"denominationUniteLegale"`
	ActivitePrincipaleUniteLegale  string `json:"activitePrincipaleUniteLegale"`
	CategorieJuridiqueUniteLegale  string `json:"categorieJuridiqueUniteLegale"`
	StatutAdministratifUniteLegale string `json:"statutAdministratifUniteLegale"`
}

type sireneAdresseEtablissement struct {
	NumeroVoieEtablissement          string `json:"numeroVoieEtablissement"`
	TypeVoieEtablissement            string `json:"typeVoieEtablissement"`
	LibelleVoieEtablissement         string `json:"libelleVoieEtablissement"`
	ComplementAdresseEtablissement   string `json:"complementAdresseEtablissement"`
	CodePostalEtablissement          string `json:"codePostalEtablissement"`
	LibelleCommuneEtablissement      string `json:"libelleCommuneEtablissement"`
	LibellePaysEtrangerEtablissement string `json:"LibellePaysEtrangerEtablissement"`
}