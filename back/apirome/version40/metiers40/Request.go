package metiers40

import (
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"net/http"
)

func RequestGetListMetier(token string) *http.Request {
	apiUrl := "https://api.francetravail.io/partenaire/rome-metiers/v1/metiers/metier"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête Get list metier summary : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

func RequestGetOneMetier(token, codeMetier string) *http.Request {
	apiUrl := fmt.Sprintf("https://api.francetravail.io/partenaire/rome-metiers/v1/metiers/metier/%s", codeMetier)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête get one metier detail : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}
