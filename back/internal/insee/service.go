package insee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/nsevenpack/env/env"
)

var (
	token     string
	tokenFile = "internal/insee/token.txt"
)

func LoadToken() error {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		if os.IsNotExist(err) {
			token = ""
			return nil
		}
		return err
	}

	token = strings.TrimSpace(string(data))
	return nil
}

func SaveToken() error {
	return os.WriteFile(tokenFile, []byte(token), 0644)
}

func GetToken() string {
	return token
}

func RefreshToken() (string, error) {
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

	var tokenRes AccessToken
	if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		return "", err
	}

	token = tokenRes.AccessToken

	if err := SaveToken(); err != nil {
		return "", err
	}

	return token, nil
}

func findCompanyBySiret(siret string) (bool, error) {
	url := fmt.Sprintf("https://api.insee.fr/entreprises/sirene/V3.11/siret/%s", siret)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer " + GetToken())
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	case 401:
		return false, fmt.Errorf("unauthorized")
	default:
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("erreur API Sirene: %d - %s", resp.StatusCode, string(body))
	}
}

func CheckSiretExists(siret string) (bool, error) {
	exists, err := findCompanyBySiret(siret)
	if err == nil {
		return exists, nil
	}

	if strings.Contains(err.Error(), "unauthorized") {
		_, refreshErr := RefreshToken()
		if refreshErr != nil {
			return false, fmt.Errorf("échec du refresh token après 401: %w", refreshErr)
		}

		return findCompanyBySiret(siret)
	}

	return false, err
}
