package apirome

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"net/http"
	"net/url"
	"tenjin/back/cli/cliutils"
	"time"
)

type Auth struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
	Scope        string
	httpClient   *cliutils.HTTPClient
}

func NewAuth() *Auth {
	return &Auth{
		TokenURL:     "https://entreprise.pole-emploi.fr/connexion/oauth2/access_token?realm=/partenaire",
		ClientID:     env.Get("API_ROME_4_0_CLIENT_ID"),
		ClientSecret: env.Get("API_ROME_4_0_CLIENT_SECRET"),
		Scope:        env.Get("API_ROME_4_0_SCOPE"),
		httpClient:   cliutils.NewHTTPClient(30 * time.Second),
	}
}

// GetToken récupère un token d'accès pour l'API ROME
func (am *Auth) GetToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", am.ClientID)
	data.Set("client_secret", am.ClientSecret)
	data.Set("scope", am.Scope)

	req, _ := http.NewRequest("POST", am.TokenURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := am.httpClient.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'appel de l'API pour récupérer le token : %v", err)
	}
	defer resp.Body.Close()

	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("erreur lors du décodage de la réponse JSON : %v", err)
	}

	logger.If("Token récupéré avec succès : %s", res.AccessToken)
	return res.AccessToken, nil
}
