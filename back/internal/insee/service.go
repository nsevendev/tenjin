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
	"tenjin/back/internal/addresses"
	"tenjin/back/internal/utils/constantes"
	"time"

	"github.com/nsevenpack/logger/v2/logger"

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
		if os.IsNotExist(err) {
			token = ""
			return nil
		}
		logger.Ef("Une erreur est survenue lors de la lecture du fichier %v: %v", tokenFile, err)
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
func buildAddressFromSireneData(a *sireneAdresseEtablissement, typeAdresse constantes.TypeAddress) addresses.Address {
    streetParts := []string{
        a.TypeVoieEtablissement,
        a.LibelleVoieEtablissement,
        a.ComplementAdresseEtablissement,
    }
    street := strings.Join(strings.Fields(strings.Join(streetParts, " ")), " ")

    country := "france"
    if a.LibellePaysEtrangerEtablissement != "" {
        country = strings.ToLower(strings.TrimSpace(a.LibellePaysEtrangerEtablissement))
    }

    return addresses.Address{
        Number:      strings.TrimSpace(a.NumeroVoieEtablissement),
        Street:      street,
        ZipCode:     strings.TrimSpace(a.CodePostalEtablissement),
        City:        strings.TrimSpace(a.LibelleCommuneEtablissement),
        Country:     constantes.Country(country),
        TypeAddress: constantes.TypeAddress(typeAdresse),
    }
}

func deriveType(cj string) constantes.TypeInstitute {
	cj = strings.TrimSpace(cj)

	if cj == "" {
		return constantes.InstitutePrivate
	}

	switch {
	case strings.HasPrefix(cj, "7"):
		return constantes.InstitutePublic
	case strings.HasPrefix(cj, "8"):
		return constantes.InstituteAssociation
	default:
		return constantes.InstitutePrivate
	}
}

func isEmptyAddress(addr addresses.Address) bool {
    return addr.Number == "" &&
        addr.Street == "" &&
        addr.ZipCode == "" &&
        addr.City == ""
}

func mapSireneStatusToState(sireneStatus string) constantes.StatusState {
	switch sireneStatus {
	case "A":
		return constantes.StateEnable
	case "C":
		return constantes.StateDisable
	case "S":
		return constantes.StateSuspended
	default:
		return constantes.StateArchived
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

	var addrs []addresses.Address
	addr1 := buildAddressFromSireneData(&etab.AdresseEtablissement, constantes.TypeAddress("headOffice"))
	if !isEmptyAddress(addr1) {
		addrs = append(addrs, addr1)
	}
	if etab.Adresse2Etablissement != nil {
		addr2 := buildAddressFromSireneData(etab.Adresse2Etablissement, constantes.TypeAddress("other"))
		if !isEmptyAddress(addr2) {
			addrs = append(addrs, addr2)
		}
	}

	cj := strings.TrimSpace(etab.UniteLegale.CategorieJuridiqueUniteLegale)
	etype := deriveType(cj)

	status := mapSireneStatusToState(etab.UniteLegale.StatutAdministratifUniteLegale)

	ci := &CompanyInfo{
		BusinessName: name,
		Siret:        strings.TrimSpace(etab.Siret),
		Addresses:    addrs,
		Status:       string(status),
		Type:         etype,
	}

	return ci, nil
}


func CheckSiretExists(siret string, siren string) (*CompanyInfo, error) {
	companyInfo, err := findCompanyBySiretAndSiren(siret, siren)

	if err == nil {
		if companyInfo != nil {
			logger.If("Entreprise trouvée: %s, SIRET: %s",
				companyInfo.BusinessName, companyInfo.Siret)
		}
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

	return nil, err
}
