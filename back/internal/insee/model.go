package insee

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type CompanyInfo struct {
	Siret   string `json:"siret"`
	Siren   string `json:"siren"`
	Adresse struct {
		NumeroVoie        string `json:"numeroVoie"`
		TypeVoie          string `json:"typeVoie"`
		LibelleVoie       string `json:"libelleVoie"`
		ComplementAdresse string `json:"complementAdresse"`
		CodePostal        string `json:"codePostal"`
		LibelleCommune    string `json:"libelleCommune"`
	} `json:"adresseEtablissement"`
}