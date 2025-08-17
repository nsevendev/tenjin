package insee

import (
	"os"
	"tenjin/back/internal/utils/constantes"
	"testing"

	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
)

var (
	tempTokenFile     string
	originalTokenFile string
	testToken         = "token-test"
)

func TestMain(m *testing.M) {
	originalTokenFile = tokenFile

	tmpFile, err := os.CreateTemp("", "token_test_*.txt")
	if err != nil {
		logger.Ef("Erreur lors de la creation du fichier temporaire : %v", err)
		panic("Erreur lors de la creation du fichier temporaire : " + err.Error())
	}
	tempTokenFile = tmpFile.Name()
	tmpFile.Close()

	tokenFile = tempTokenFile

	code := m.Run()

	_ = os.Remove(tempTokenFile)
	tokenFile = originalTokenFile

	os.Exit(code)
}

func Test_buildAddressFromSireneData(t *testing.T) {
	testup.LogNameTestInfo(t, "Test build address from sirene data")

	a := sireneAdresseEtablissement{
		NumeroVoieEtablissement:          "10",
		TypeVoieEtablissement:            "rue",
		LibelleVoieEtablissement:         "des Écoles",
		ComplementAdresseEtablissement:   "Bât A",
		CodePostalEtablissement:          "75001",
		LibelleCommuneEtablissement:      "Paris",
		LibellePaysEtrangerEtablissement: "",
	}

	addr := buildAddressFromSireneData(&a, constantes.TypeAddress("headOffice"))

	assert.Equal(t, "10", addr.Number)
	assert.Equal(t, "rue des Écoles Bât A", addr.Street)
	assert.Equal(t, "75001", addr.ZipCode)
	assert.Equal(t, "Paris", addr.City)
	assert.Equal(t, "france", string(addr.Country))
	assert.Equal(t, "headOffice", string(addr.TypeAddress))
}

func Test_deriveType(t *testing.T) {
	testup.LogNameTestInfo(t, "Test derive type from code juridique")

	assert.Equal(t, constantes.InstitutePublic, deriveType("7210"))
	assert.Equal(t, constantes.InstitutePrivate, deriveType("5498"))
	assert.Equal(t, constantes.InstitutePrivate, deriveType(""))
	assert.Equal(t, constantes.InstituteAssociation, deriveType("851"))
}

func Test_mapSireneStatusToState(t *testing.T) {
	testup.LogNameTestInfo(t, "Test map sirene status to internal state")

	assert.Equal(t, constantes.StatusStateEnable, mapSireneStatusToState("A"))
	assert.Equal(t, constantes.StatusStateDisable, mapSireneStatusToState("C"))
	assert.Equal(t, constantes.StatusStateSuspended, mapSireneStatusToState("S"))
	assert.Equal(t, constantes.StatusStateArchived, mapSireneStatusToState("X"))
}
