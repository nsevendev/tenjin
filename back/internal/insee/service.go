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
	"time"

	"github.com/nsevenpack/env/env"
)

var (
	token     string
	tokenFile = "internal/insee/token.txt"
)

// token

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
		return "", fmt.Errorf("erreur lors de la recuperation du token: %s", string(body))
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

// check siret/siren + return basic company infos

func buildAddressFromSireneData(a sireneAdresseEtablissement) string {
	parts := []string{
		a.NumeroVoieEtablissement,
		a.TypeVoieEtablissement,
		a.LibelleVoieEtablissement,
		a.ComplementAdresseEtablissement,
	}
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	addr := strings.Join(out, " ")
	return strings.Join(strings.Fields(addr), " ")
}

func deriveSector(cj string) string {
	cj = strings.TrimSpace(cj)
	if cj != "" && strings.HasPrefix(cj, "7") {
		return "public"
	}
	return "private"
}

func mapAPEtoCompType(ape string) string {
	ape = strings.TrimSpace(ape)
	if len(ape) < 2 {
		return "company"
	}
	prefix := ape[:2]
	switch prefix {
	case "85":
		return "training_center"
	case "78":
		return "recruiting_agency"
	default:
		return "company"
	}
}

func findCompanyBySiretAndSiren(siret string, siren string) (*CompanyInfo, error) {
	url := fmt.Sprintf("https://api.insee.fr/entreprises/sirene/V3.11/siret/%s", siret)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+GetToken())
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		var sr sireneResponse
		if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
			return nil, err
		}
		etab := sr.Etablissement

		if siren != "" && etab.UniteLegale.Siren != "" && etab.UniteLegale.Siren != siren {
			return nil, fmt.Errorf("siren mismatch: attendu %s, trouvé %s", siren, etab.UniteLegale.Siren)
		}

		name := strings.TrimSpace(etab.UniteLegale.DenominationUniteLegale)
		if name == "" {
			name = strings.TrimSpace(etab.Enseigne1Etablissement)
		}

		addr := buildAddressFromSireneData(etab.AdresseEtablissement)
		zip := strings.TrimSpace(etab.AdresseEtablissement.CodePostalEtablissement)
		city := strings.TrimSpace(etab.AdresseEtablissement.LibelleCommuneEtablissement)

		ape := strings.TrimSpace(etab.UniteLegale.ActivitePrincipaleUniteLegale)
		cj := strings.TrimSpace(etab.UniteLegale.CategorieJuridiqueUniteLegale)

		ci := &CompanyInfo{
			BusinessName:       name,
			Siret:              strings.TrimSpace(etab.Siret),
			Siren:              siren,
			Address:            addr,
			ZipCode:            zip,
			City:               city,
			Ape:                ape,
			CategorieJuridique: cj,
			Sector:             deriveSector(cj),
			CompType:           mapAPEtoCompType(ape),
		}
		return ci, nil

	case 404:
		return nil, nil
	case 401:
		return nil, fmt.Errorf("unauthorized")
	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("erreur API Sirene: %d - %s", resp.StatusCode, string(body))
	}
}


func CheckSiretExists(siret string, siren string) (*CompanyInfo, error) {
	companyInfo, err := findCompanyBySiretAndSiren(siret, siren)
	if err == nil {
		return companyInfo, nil
	}

	if strings.Contains(err.Error(), "unauthorized") {
		_, refreshErr := RefreshToken()
		if refreshErr != nil {
			return nil, fmt.Errorf("echec du refresh token apres 401: %w", refreshErr)
		}

		return findCompanyBySiretAndSiren(siret, siren)
	}

	return companyInfo, err
}
