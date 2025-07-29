package insee

import (
	"os"
	"testing"

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

func TestSaveToken_WriteToFile(t *testing.T) {
	token = testToken

	err := SaveToken()
	assert.Nil(t, err)

	content, err := os.ReadFile(tokenFile)
	assert.Nil(t, err)
	assert.Equal(t, testToken, string(content))
}

func TestLoadToken_ReadFromFile(t *testing.T) {
	expected := "token-du-fichier"
	err := os.WriteFile(tokenFile, []byte(expected), 0644)
	assert.Nil(t, err)

	err = LoadToken()
	assert.Nil(t, err)
	assert.Equal(t, expected, token)
}

func TestLoadToken_FileDoesNotExist(t *testing.T) {
	nonExistentFile := tempTokenFile + "_missing"

	oldTokenFile := tokenFile
	tokenFile = nonExistentFile
	defer func() { tokenFile = oldTokenFile }()

	err := LoadToken()
	assert.Nil(t, err)
	assert.Equal(t, "", token)
}

func TestGetToken_ReturnsInMemoryToken(t *testing.T) {
	token = "in-memory-token"
	actual := GetToken()
	assert.Equal(t, "in-memory-token", actual)
}

func TestRefreshToken_RefreshToken(t *testing.T) {
	clientID := os.Getenv("SIRENE_CLIENT_KEY")
	clientSecret := os.Getenv("SIRENE_CLIENT_SECRET")

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

func TestFindCompanyBySiretAndSiren_Success(t *testing.T) {
	err := LoadToken()
	assert.Nil(t, err)

	siret := "94503764600011"
	siren := "945037646"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
	assert.Equal(t, siren, info.Siren)
}

func TestFindCompanyBySiretAndSiren_NotFound(t *testing.T) {
	err := LoadToken()
	assert.Nil(t, err)

	siret := "00000000000000"
	siren := "000000000"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFindCompanyBySiretAndSiren_SirenMismatch(t *testing.T) {
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
	token = "invalid-token"

	siret := "94503764600011"
	siren := "945037646"

	info, err := findCompanyBySiretAndSiren(siret, siren)

	assert.Nil(t, info)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unauthorized")
}

func TestCheckSiretExists_Success(t *testing.T) {
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
	err := LoadToken()
	assert.Nil(t, err)

	siret := "00000000000000"
	siren := "000000000"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCheckSiretExists_UnauthorizedWithSuccessfulRefresh(t *testing.T) {
	token = "invalid-token"

	siret := "94503764600011"
	siren := "945037646"

	info, err := CheckSiretExists(siret, siren)

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, siret, info.Siret)
}

func TestCheckSiretExists_UnauthorizedWithFailedRefresh(t *testing.T) {
	token = "invalid-token"

	origClientID := os.Getenv("SIRENE_CLIENT_KEY")
	origClientSecret := os.Getenv("SIRENE_CLIENT_SECRET")
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
