package insee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/nsevenpack/env/env"
)

// fonction qui génére un token a partir des clefs et secret conssomateur
// fonction qui fait une requete api vers sirene pour vérifier que le siret donné correspond a une entreprise existante dans la bdd insee
// doit également appellé la premiere fonction si échec de l'appel a cause d'un token expiré puis mettre a jour le token de l'appli ?

func GetToken() (string, error) {
    clientID := env.Get("SIRENE_CLIENT_KEY")
    clientSecret := env.Get("SIRENE_CLIENT_SECRET")

    if clientID == "" || clientSecret == "" {
        return "", errors.New("SIRENE_CLIENT_KEY ou SIRENE_CLIENT_SECRET introuvable dans .env")
    }

    url := "https://api.insee.fr/token"
    data := []byte("grant_type=client_credentials")

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    if err != nil {
        return "", err
    }
    req.SetBasicAuth(clientID, clientSecret)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        body, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("erreur lors de la récupération du token: %s", string(body))
    }

    var tokenRes Token
    if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
        return "", err
    }

    return tokenRes.AccessToken, nil
}