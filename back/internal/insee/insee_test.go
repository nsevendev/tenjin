package insee

import (
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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
		NumeroVoieEtablissement:        "10",
		TypeVoieEtablissement:          "rue",
		LibelleVoieEtablissement:       "des Écoles",
		ComplementAdresseEtablissement: "Bât A",
	}
	addr := buildAddressFromSireneData(&a)
	assert.Equal(t, "10 rue des Écoles Bât A", addr)
}

func Test_deriveSector(t *testing.T) {
	testup.LogNameTestInfo(t, "Test derive sector from code juridique")

	assert.Equal(t, "public", deriveSector("7210"))
	assert.Equal(t, "private", deriveSector("5498"))
	assert.Equal(t, "private", deriveSector(""))
}

func Test_mapAPEtoCompType(t *testing.T) {
	testup.LogNameTestInfo(t, "Test map APE to company type")

	assert.Equal(t, "training_center", mapAPEtoCompType("8510Z"))
	assert.Equal(t, "recruiting_agency", mapAPEtoCompType("7820Z"))
	assert.Equal(t, "company", mapAPEtoCompType("6202A"))
	assert.Equal(t, "company", mapAPEtoCompType(""))
}
