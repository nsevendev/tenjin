package insee

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type CompanyInfo struct {
    BusinessName            string `json:"business_name"`
    Siret                   string `json:"siret"`
	Siren                   string `json:"siren"`
    Sector                  string `json:"sector"`
	CompType                string `json:"comp_type"`
    Address                 string `json:"adress"`
    ZipCode                 string `json:"zipCode"`
    City                    string `json:"city"`
    Ape                     string `json:"ape"`
    CategorieJuridique      string `json:"categorie_juridique"`
}

type sireneResponse struct {
	Etablissement sireneEtablissement `json:"etablissement"`
}

type sireneEtablissement struct {
	Siret                 string                   `json:"siret"`
	UniteLegale           sireneUniteLegale        `json:"uniteLegale"`
	AdresseEtablissement  sireneAdresseEtablissement `json:"adresseEtablissement"`
	Enseigne1Etablissement string                  `json:"enseigne1Etablissement"`
}

type sireneUniteLegale struct {
	Siren                          string `json:"siren"`
	DenominationUniteLegale        string `json:"denominationUniteLegale"`
	ActivitePrincipaleUniteLegale  string `json:"activitePrincipaleUniteLegale"`
	CategorieJuridiqueUniteLegale  string `json:"categorieJuridiqueUniteLegale"`
}

type sireneAdresseEtablissement struct {
	NumeroVoieEtablissement        string `json:"numeroVoieEtablissement"`
	TypeVoieEtablissement          string `json:"typeVoieEtablissement"`
	LibelleVoieEtablissement       string `json:"libelleVoieEtablissement"`
	ComplementAdresseEtablissement string `json:"complementAdresseEtablissement"`
	CodePostalEtablissement        string `json:"codePostalEtablissement"`
	LibelleCommuneEtablissement    string `json:"libelleCommuneEtablissement"`
}