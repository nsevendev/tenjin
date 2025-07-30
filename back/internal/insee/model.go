package insee

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type CompanyInfo struct {
    BusinessName            string `json:"business_name"`
    Siret                   string `json:"siret"`
    Sector                  string `json:"sector"`
	CompType                string `json:"comp_type"`
    Address                 string `json:"adress"`
    ZipCode                 string `json:"zipCode"`
    City                    string `json:"city"`
    Ape                     string `json:"ape"`
    CategorieJuridique      string `json:"categorie_juridique"`
}