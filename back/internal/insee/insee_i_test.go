//go:build integration

package insee

import (
	"fmt"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/testup"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveToken_WriteToFile(t *testing.T) {
	testup.LogNameTestInfo(t, "Test SaveToken writes to file")

	token = testToken

	err := SaveToken()
	assert.Nil(t, err, "OK")

	content, err := os.ReadFile(tokenFile)
	assert.Nil(t, err, "OK")
	assert.Equal(t, testToken, string(content))
}

func TestLoadToken_ReadFromFile(t *testing.T) {
	testup.LogNameTestInfo(t, "Test LoadToken reads from file")

	expected := "token-du-fichier"
	err := os.WriteFile(tokenFile, []byte(expected), 0644)
	assert.Nil(t, err)

	err = LoadToken()
	assert.Nil(t, err)
	assert.Equal(t, expected, token)
}

func TestLoadToken_FileDoesNotExist(t *testing.T) {
	testup.LogNameTestInfo(t, "Test LoadToken handles non-existent file")

	nonExistentFile := tempTokenFile + "_missing"

	oldTokenFile := tokenFile
	tokenFile = nonExistentFile
	defer func() { tokenFile = oldTokenFile }()

	err := LoadToken()
	assert.Nil(t, err, "OK")
	assert.Equal(t, "", token)
}

func TestGetToken_ReturnsInMemoryToken(t *testing.T) {
	testup.LogNameTestInfo(t, "Test GetToken returns in-memory token")

	token = "in-memory-token"
	actual := GetToken()
	assert.Equal(t, "in-memory-token", actual)
}

func TestRefreshToken_RefreshToken(t *testing.T) {
	testup.LogNameTestInfo(t, "Test RefreshToken refreshes token")

	// TODO : à changer avec package env
	clientID := env.Get("SIRENE_CLIENT_KEY")
	clientSecret := env.Get("SIRENE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		t.Fatal("SIRENE_CLIENT_KEY ou SIRENE_CLIENT_SECRET non définis")
	}

	newToken, err := RefreshToken()
	assert.Nil(t, err)
	assert.NotEmpty(t, newToken)
	assert.Equal(t, newToken, GetToken())

	content, err := os.ReadFile(tokenFile)
	assert.Nil(t, err)
	assert.Equal(t, newToken, string(content))
}

// siret/siren

func TestFindCompanyBySiretAndSiren_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test FindCompanyBySiretAndSiren Success")

	err := LoadToken()
	assert.Nil(t, err)

	siret := "67205008502051"
	siren := "672050085"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	fmt.Printf("Résultat CompanyInfo : %+v\n", info)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
	assert.Equal(t, siren, info.Siren)
}

func TestFindCompanyBySiretAndSiren_NotFound(t *testing.T) {
	testup.LogNameTestInfo(t, "Test FindCompanyBySiretAndSiren NotFound")

	err := LoadToken()
	assert.Nil(t, err)

	siret := "00000000000000"
	siren := "000000000"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFindCompanyBySiretAndSiren_SirenMismatch(t *testing.T) {
	testup.LogNameTestInfo(t, "Test FindCompanyBySiretAndSiren SirenMismatch")

	err := LoadToken()
	assert.Nil(t, err)

	siret := "94503764600011"
	wrongSiren := "123456789"

	info, err := findCompanyBySiretAndSiren(siret, wrongSiren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "siren mismatch")
}

func TestFindCompanyBySiretAndSiren_Unauthorized(t *testing.T) {
	testup.LogNameTestInfo(t, "Test FindCompanyBySiretAndSiren Unauthorized")

	token = "invalid-token"

	siret := "94503764600011"
	siren := "945037646"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unauthorized")
}

func TestCheckSiretExists_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test CheckSiretExists Success")

	err := LoadToken()
	assert.Nil(t, err)

	siret := "94503764600011"
	siren := "945037646"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
	assert.Equal(t, siren, info.Siren)
}

func TestCheckSiretExists_NotFound(t *testing.T) {
	testup.LogNameTestInfo(t, "Test CheckSiretExists NotFound")

	err := LoadToken()
	assert.Nil(t, err)

	siret := "00000000000000"
	siren := "000000000"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFindCompanyBySiretAndSiren_BadRequest_DefaultBranch(t *testing.T) {
	testup.LogNameTestInfo(t, "Test FindCompanyBySiretAndSiren BadRequest with default branch")

	err := LoadToken()
	assert.Nil(t, err)

	if GetToken() == "" {
		if env.Get("SIRENE_CLIENT_KEY") == "" || env.Get("SIRENE_CLIENT_SECRET") == "" {
			t.Skip("Pas de token ni de credentials pour obtenir un 400 réel")
		}
		_, _ = RefreshToken()
	}

	badSiret := "ABC"
	info, err := findCompanyBySiretAndSiren(badSiret, "")
	if err != nil && strings.Contains(err.Error(), "unauthorized") {
		t.Skip("Token invalide — test 400 ignoré")
	}
	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erreur API Sirene: 400")
}

func TestCheckSiretExists_UnauthorizedWithSuccessRefresh(t *testing.T) {
	testup.LogNameTestInfo(t, "Test CheckSiretExists Unauthorized with successful refresh")

	token = "invalid-token"

	siret := "94503764600011"
	siren := "945037646"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
}

func TestCheckSiretExists_UnauthorizedWithFailedRefresh(t *testing.T) {
	testup.LogNameTestInfo(t, "Test CheckSiretExists Unauthorized with failed refresh")

	token = "invalid-token"

	origClientID := env.Get("SIRENE_CLIENT_KEY")
	origClientSecret := env.Get("SIRENE_CLIENT_SECRET")
	os.Unsetenv("SIRENE_CLIENT_KEY")
	os.Unsetenv("SIRENE_CLIENT_SECRET")
	defer func() {
		os.Setenv("SIRENE_CLIENT_KEY", origClientID)
		os.Setenv("SIRENE_CLIENT_SECRET", origClientSecret)
	}()

	siret := "94503764600011"
	siren := "945037646"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "echec du refresh token")
}
