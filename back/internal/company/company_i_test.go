//go:build integration

package company

import (
	"context"
	"os"
	"testing"

	"github.com/nsevenpack/testup"

	"tenjin/back/internal/insee"

	"github.com/stretchr/testify/assert"
)

var companyServiceTest *companyService

func TestMain(m *testing.M) {
	tmpFile, err := os.CreateTemp("", "token_test_*.txt")
	if err != nil {
		panic("Erreur lors de la création du fichier temporaire : " + err.Error())
	}

	dummyToken := []byte("dummy_test_token")
	if _, err = tmpFile.Write(dummyToken); err != nil {
		panic("Erreur lors de l’écriture dans le fichier temporaire : " + err.Error())
	}
	tmpFile.Close()

	insee.SetTokenFile(tmpFile.Name())

	companyServiceTest = &companyService{}

	code := m.Run()

	_ = os.Remove(tmpFile.Name())

	os.Exit(code)
}

func TestRetrieveCompanyInfo_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Retrieve Company Info Success")

	siret := "94503764600011"
	siren := "945037646"

	info, err := companyServiceTest.RetrieveCompanyInfo(context.Background(), siret, siren)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
	assert.Equal(t, siren, info.Siren)
}

func TestRetrieveCompanyInfo_NotFound(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Retrieve Company Info Not Found")

	siret := "00000000000000"
	siren := "000000000"

	info, err := companyServiceTest.RetrieveCompanyInfo(context.Background(), siret, siren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "aucune entreprise trouvee")
}

func TestRetrieveCompanyInfo_MissingSiret(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Retrieve Company Info Missing Siret")

	info, err := companyServiceTest.RetrieveCompanyInfo(context.Background(), "", "123456789")

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "le SIRET est requis")
}

func TestRetrieveCompanyInfo_MissingSiren(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Retrieve Company Info Missing Siren")

	info, err := companyServiceTest.RetrieveCompanyInfo(context.Background(), "12345678900000", "")

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "le SIREN est requis")
}

func TestRetrieveCompanyInfo_InvalidSiret(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Retrieve Company Info Invalid Siret")

	siret := "abc"
	siren := "123456789"

	info, err := companyServiceTest.RetrieveCompanyInfo(context.Background(), siret, siren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "echec lors de la recuperation des donnees INSEE")
}
