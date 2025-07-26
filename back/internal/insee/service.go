package insee

// fonction qui génére un token a partir des clefs et secret conssomateur
// fonction qui fait une requete api vers sirene pour vérifier que le siret donné correspond a une entreprise existante dans la bdd insee
// doit également appellé la premiere fonction si échec de l'appel a cause d'un token expiré puis mettre a jour le token de l'appli ?

import (
	"os"
	"strings"
)

var (
	token     string
	tokenFile = "token.txt"
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

