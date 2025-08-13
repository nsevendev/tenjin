package insee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
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

func SetTokenFile(path string) {
	tokenFile = path
}

func LoadToken() error {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		logger.Ef("Une erreur est survenue lors de la lecture du fichier %v: %v", tokenFile, err)
		if os.IsNotExist(err) {
			token = ""
			return nil
		}
		return err
	}

	logger.Sf("Token chargé depuis le fichier %v", tokenFile)

	token = strings.TrimSpace(string(data))
	return nil
}

func SaveToken() error {
	err := os.WriteFile(tokenFile, []byte(token), 0644)

	if err != nil {
		logger.Ef("Une erreur est survenu au moment du chargement du fichier%v", err)
		return err
	}

	logger.Sf("Token enregistré dans %s", tokenFile)

	return nil
}

func GetToken() string {
	return token
}

func RefreshToken() (string, error) {
	clientID := env.Get("SIRENE_CLIENT_KEY")
	clientSecret := env.Get("SIRENE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		logger.Ef("SIRENE_CLIENT_KEY ou SIRENE_CLIENT_SECRET introuvable dans .env")
		return "", errors.New("SIRENE_CLIENT_KEY ou SIRENE_CLIENT_SECRET introuvable dans .env")
	}

	url := "https://api.insee.fr/token"
	data := []byte("grant_type=client_credentials")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		logger.Ef("Erreur lors de la création de la requête HTTP: %v", err)
		return "", err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Ef("Erreur lors de l'envoi de la requête HTTP: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		logger.Ef("Erreur lors de la récupération du token: %d - %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("erreur lors de la recuperation du token: %s", string(body))
	}

	var tokenRes AccessToken
	if err = json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		logger.Ef("Erreur lors du décodage de la réponse JSON: %v", err)
		return "", err
	}

	token = tokenRes.AccessToken

	if err = SaveToken(); err != nil {
		logger.Ef("Erreur lors de l'enregistrement du token: %v", err)
		return "", err
	}

	logger.Sf("Nouveau token enregistré: %s", token)

	return token, nil
}

// check siret/siren + return basic company infos
func buildAddressFromSireneData(a *sireneAdresseEtablissement) string {
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
		logger.Ef("Erreur lors de la création de la requête HTTP: %v", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+GetToken())
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Ef("Erreur lors de l'envoi de la requête HTTP: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			logger.If("Aucune entreprise trouvée pour le SIRET %s", siret)
			return nil, nil
		}
		if resp.StatusCode == 401 {
			logger.Ef("Erreur d'authentification: 401 Unauthorized")
			return nil, fmt.Errorf("unauthorized")
		}
		body, _ := io.ReadAll(resp.Body)
		logger.Ef("Erreur API Sirene: %d - %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("erreur API Sirene: %d - %s", resp.StatusCode, string(body))
	}

	var sr sireneResponse
	if err = json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		logger.Ef("Erreur lors du décodage de la réponse JSON: %v", err)
		return nil, err
	}
	etab := sr.Etablissement

	if siren != "" && etab.UniteLegale.Siren != "" && etab.UniteLegale.Siren != siren {
		logger.Ef("Siren mismatch: attendu %s, trouvé %s", siren, etab.UniteLegale.Siren)
		return nil, fmt.Errorf("siren mismatch: attendu %s, trouvé %s", siren, etab.UniteLegale.Siren)
	}

	apiSiren := strings.TrimSpace(etab.UniteLegale.Siren)

	if apiSiren == "" && len(siret) >= 9 {
		apiSiren = siret[:9]
	}

	if siren != "" && apiSiren != "" && apiSiren != siren {
		logger.Ef("Siren mismatch: attendu %s, trouvé %s", siren, apiSiren)
		return nil, fmt.Errorf("siren mismatch: attendu %s, trouvé %s", siren, apiSiren)
	}

	name := strings.TrimSpace(etab.UniteLegale.DenominationUniteLegale)
	if name == "" {
		name = strings.TrimSpace(etab.Enseigne1Etablissement)
	}

	addr := buildAddressFromSireneData(&etab.AdresseEtablissement)
	zip := strings.TrimSpace(etab.AdresseEtablissement.CodePostalEtablissement)
	city := strings.TrimSpace(etab.AdresseEtablissement.LibelleCommuneEtablissement)
	ape := strings.TrimSpace(etab.UniteLegale.ActivitePrincipaleUniteLegale)
	cj := strings.TrimSpace(etab.UniteLegale.CategorieJuridiqueUniteLegale)

	logger.If("adresse: %s, zip: %s, city: %s, ape: %s, cj: %s", addr, zip, city, ape, cj)

	ci := &CompanyInfo{
		BusinessName:       name,
		Siret:              strings.TrimSpace(etab.Siret),
		Siren:              apiSiren,
		Address:            addr,
		ZipCode:            zip,
		City:               city,
		Ape:                ape,
		CategorieJuridique: cj,
		Sector:             deriveSector(cj),
		CompType:           mapAPEtoCompType(ape),
	}
	return ci, nil
}

func CheckSiretExists(siret string, siren string) (*CompanyInfo, error) {
	companyInfo, err := findCompanyBySiretAndSiren(siret, siren)
	if err == nil {
		logger.If("Entreprise trouvée: %s, SIRET: %s, SIREN: %s", companyInfo.BusinessName, companyInfo.Siret, companyInfo.Siren)
		return companyInfo, nil
	}

	if strings.Contains(err.Error(), "unauthorized") {
		_, refreshErr := RefreshToken()
		if refreshErr != nil {
			logger.Ef("Echec du refresh token apres 401: %v", refreshErr)
			return nil, fmt.Errorf("echec du refresh token apres 401: %w", refreshErr)
		}

		return findCompanyBySiretAndSiren(siret, siren)
	}

	return companyInfo, err
}
