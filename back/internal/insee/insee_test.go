package insee

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tempTokenFile      string
	originalTokenFile  string
	testToken          = "token-test"
)

func TestMain(m *testing.M) {
	originalTokenFile = tokenFile

	tmpFile, err := os.CreateTemp("", "token_test_*.txt")
	if err != nil {
		panic("Erreur lors de la création du fichier temporaire : " + err.Error())
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
