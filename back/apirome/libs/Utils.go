package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// PrintBrutInFile enregistre un slice dans un fichier JSON avec un nom basé sur la date actuelle
func PrintBrutInFile(path string, body []byte) {
	// 1. Créer le dossier path s'il n'existe pas
	err := os.MkdirAll(path, 0755)
	if err != nil {
		logger.Ff("erreur lors de la création du dossier %s : %v", path, err)
	}

	// 2. Définir le nom du fichier
	now := time.Now().Format("20060102_150405") // format : YYYYMMDD_HHMMSS
	filename := fmt.Sprintf("%s/%s.json", path, now)

	// 3. Sauvegarder le body dans le fichier
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		logger.Ff("erreur lors de l'écriture dans le fichier %s : %v", filename, err)
	} else {
		logger.If("fichier JSON sauvegardé : %s", filename)
	}
}

// PrintSliceBrutInFile sauvegarde un slice en JSON dans un fichier
func PrintSliceBrutInFile(path string, slice any) {
	// 1. Créer le dossier path s'il n'existe pas
	err := os.MkdirAll(path, 0755)
	if err != nil {
		logger.Ff("erreur lors de la création du dossier %s : %v", path, err)
	}

	// 2. Définir le nom du fichier
	now := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/%s.json", path, now)

	// ⬇️ Ici : Marshal le slice en []byte
	body, errMarchal := json.MarshalIndent(slice, "", "  ")
	if errMarchal != nil {
		logger.Ff("erreur lors du marshal JSON : %v", errMarchal)
	}

	// 3. Sauvegarder le body dans le fichier
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		logger.Ff("erreur lors de l'écriture dans le fichier %s : %v", filename, err)
	} else {
		logger.If("fichier JSON sauvegardé : %s", filename)
	}
}

// ExecuteRequest exécute une requête HTTP et retourne le corps de la réponse
func ExecuteRequest(req *http.Request) []byte {
	clientHTTP := http.Client{Timeout: 30 * time.Second}

	resp, err := clientHTTP.Do(req)
	if err != nil {
		logger.Ff("erreur lors de l'appel : %v", err)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			logger.Ff("erreur lors de la fermeture de l'appel : %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		logger.Ff("erreur retour status : %d %s", resp.StatusCode, resp.Status)
	}

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		logger.Ff("erreur lors de la lecture du corps de la réponse : %v", errBody)
	}

	return body
}

// GetToken récupère un token d'accès pour l'API ROME
func GetToken() string {
	tokenURL := "https://entreprise.pole-emploi.fr/connexion/oauth2/access_token?realm=/partenaire"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", env.Get("API_ROME_4_0_CLIENT_ID"))
	data.Set("client_secret", env.Get("API_ROME_4_0_CLIENT_SECRET"))
	data.Set("scope", env.Get("API_ROME_4_0_SCOPE"))

	req, _ := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Ff("erreur lors de l'appel de l'API pour récupérer le token : %v", err)
	}
	defer resp.Body.Close()

	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		logger.Ff("erreur lors du décodage de la réponse JSON : %v", err)
	}

	logger.If("Token récupéré avec succès : %s", res.AccessToken)
	return res.AccessToken
}

func DecodeJSONList[T any](body []byte) ([]T, error) {
	var result []T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetLastFileByDate Récupère le chemin du dernier fichier JSON dans le dossier donné, trié par date dans le nom
func GetLastFileByDate(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		logger.Ff("Erreur lors de la lecture du dossier %s : %v", dir, err)
	}

	// Regex pour extraire la date du nom de fichier (ex: metiers_20250726_123045.json)
	re := regexp.MustCompile(`(\d{8}_\d{6})\.json$`)

	type datedFile struct {
		name string
		date time.Time
	}

	var datedFiles []datedFile

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matches := re.FindStringSubmatch(file.Name())
		if len(matches) != 2 {
			continue
		}
		parsed, err := time.Parse("20060102_150405", matches[1])
		if err != nil {
			continue
		}
		datedFiles = append(datedFiles, datedFile{filepath.Join(dir, file.Name()), parsed})
	}

	if len(datedFiles) == 0 {
		logger.Wf("Aucun fichier JSON trouvé dans le dossier %s", dir)
		logger.Ff("Vérifiez que les fichiers sont nommés correctement avec le format YYYYMMDD_HHMMSS.json %v", os.ErrNotExist)
	}

	// Trie par date décroissante
	sort.Slice(datedFiles, func(i, j int) bool {
		return datedFiles[i].date.After(datedFiles[j].date)
	})

	return datedFiles[0].name
}

// GetLastXMLFileByDate Récupère le dernier fichier XML dans le dossier donné
func GetLastXMLFileByDate(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		logger.Ff("Erreur lors de la lecture du dossier %s : %v", dir, err)
	}

	// Regex pour fichiers XML avec date (ex: rncp_2024-01-15.xml ou rncp_20240115.xml)
	re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2}|\d{8})\.xml$`)

	type datedFile struct {
		name string
		date time.Time
	}

	var datedFiles []datedFile

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matches := re.FindStringSubmatch(file.Name())
		if len(matches) != 2 {
			continue
		}

		var parsed time.Time
		var err error

		// Essayer les deux formats
		if strings.Contains(matches[1], "-") {
			parsed, err = time.Parse("2006-01-02", matches[1])
		} else {
			parsed, err = time.Parse("20060102", matches[1])
		}

		if err != nil {
			continue
		}
		datedFiles = append(datedFiles, datedFile{filepath.Join(dir, file.Name()), parsed})
	}

	if len(datedFiles) == 0 {
		logger.Wf("Aucun fichier XML trouvé dans le dossier %s", dir)
		logger.Ff("Vérifiez que les fichiers sont nommés correctement avec une date")
	}

	// Trie par date décroissante
	sort.Slice(datedFiles, func(i, j int) bool {
		return datedFiles[i].date.After(datedFiles[j].date)
	})

	return datedFiles[0].name
}
