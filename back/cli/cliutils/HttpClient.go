package cliutils

import (
	"fmt"
	"github.com/nsevenpack/logger/v2/logger"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	Client  *http.Client
	BaseURL string
	Timeout time.Duration
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		Client:  &http.Client{Timeout: timeout},
		Timeout: timeout,
	}
}

// ExecuteRequest exécute une requête HTTP et retourne le corps de la réponse
func (hc *HTTPClient) ExecuteRequest(req *http.Request) ([]byte, error) {
	resp, err := hc.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'appel : %v", err)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			logger.Ff("erreur lors de la fermeture de l'appel : %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur retour status : %d %s", resp.StatusCode, resp.Status)
	}

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du corps de la réponse : %v", errBody)
	}

	return body, nil
}
