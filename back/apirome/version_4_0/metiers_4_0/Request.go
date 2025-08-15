package metiers_4_0

import (
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"net/http"
)

// RequestGetListMetier crée une requête HTTP pour récupérer la liste des métiers en résumé
// Cette requête est utilisée pour obtenir un aperçu des métiers disponibles, avec leur code et lib
func RequestGetListMetier(token string) *http.Request {
	apiUrl := "https://api.francetravail.io/partenaire/rome-metiers/v1/metiers/metier"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête Get list metier summary : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

// RequestGetOneMetier crée une requête HTTP pour récupérer les détails d'un métier spécifique
// en utilisant son code. Cette requête est utilisée pour obtenir des informations détaillées sur un
// métier particulier, comme sa description, ses compétences requises, etc.
func RequestGetOneMetier(token, codeMetier string) *http.Request {
	apiUrl := fmt.Sprintf("https://api.francetravail.io/partenaire/rome-metiers/v1/metiers/metier/%s", codeMetier)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête get one metier detail : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

// RequestGetListCompetence crée une requête HTTP pour récupérer la liste des compétences sans details
func RequestGetListCompetence(token string) *http.Request {
	apiUrl := "https://api.francetravail.io/partenaire/rome-competences/v1/competences/competence"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête Get list competence : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

// RequestGetOneCompetenceComplet crée une requête HTTP pour récupérer les détails complets d'une compétence
func RequestGetOneCompetenceComplet(token, codeCompetence string) *http.Request {
	apiUrl := fmt.Sprintf("https://api.francetravail.io/partenaire/rome-competences/v1/competences/competence/%v", codeCompetence)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Ff("erreur lors de la création de la requête get one comeptence complet: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}
