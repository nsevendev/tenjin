package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/env/env"
	"net/http/httptest"
	"os"
	"tenjin/back/internal/utils/database"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tenjin/back/internal/crm"
)

var (
	testAuthService *AuthService
	testDB          *mongo.Database
	testJWTKey      = "test-secret-key-for-jwt-testing-123456789"
)

func TestMain(m *testing.M) {
	appEnv := env.Get("APP_ENV")

	database.ConnexionDatabase(appEnv)
	testDB = database.Client
	
	err := testDB.Collection("users").Drop(context.Background())
	if err != nil {
		logger.Ef("Erreur lors de la suppression de la collection 'users' : %v", err)
		return
	}
	testAuthService = NewAuthService(testDB, testJWTKey)

	code := m.Run()

	logger.If("Tests Auth terminés")
	os.Exit(code)
}

func TestNewAuthService(t *testing.T) {
	testup.LogNameTestInfo(t, "Test creation of AuthService")

	service := NewAuthService(testDB, "test-key")

	assert.NotNil(t, service)
	assert.Equal(t, testDB, service.db)
	assert.Equal(t, []byte("test-key"), service.jwtKey)
}

func TestAuthService_NameCookie(t *testing.T) {
	testup.LogNameTestInfo(t, "Test cookie name")

	cookieName := testAuthService.NameCookie()

	expected := "1209_!&@_entok_!&@_9834"
	assert.Equal(t, expected, cookieName)
}

func TestAuthService_generateToken_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test generate token success")

	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "$2a$10$hashedPassword",
		Role:     "user",
		Username: "testuser",
	}

	token, err := testAuthService.generateToken(user)

	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Vérifier que le token peut être parsé
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testJWTKey), nil
	})

	require.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(*tokenClaims)
	require.True(t, ok)
	assert.Equal(t, user.ID.Hex(), claims.IdUser)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.Role, claims.Role)
}

func TestAuthService_generateToken_EmptyJWTKey(t *testing.T) {
	testup.LogNameTestInfo(t, "Test generate token with empty JWT key")

	serviceWithEmptyKey := NewAuthService(testDB, "")
	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "$2a$10$hashedPassword",
		Role:     "user",
		Username: "testuser",
	}

	token, err := serviceWithEmptyKey.generateToken(user)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "secret JWT manquant")
}

func TestAuthService_generateToken_TokenExpiration(t *testing.T) {
	testup.LogNameTestInfo(t, "Test generate token expiration")

	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "$2a$10$hashedPassword",
		Role:     "user",
		Username: "testuser",
	}

	token, err := testAuthService.generateToken(user)
	require.NoError(t, err)

	// Parser le token pour vérifier l'expiration
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testJWTKey), nil
	})

	require.NoError(t, err)
	claims, ok := parsedToken.Claims.(*tokenClaims)
	require.True(t, ok)

	// Vérifier que l'expiration est dans environ 24 heures
	expectedExpiration := time.Now().Add(24 * time.Hour)
	actualExpiration := claims.ExpiresAt.Time

	timeDiff := actualExpiration.Sub(expectedExpiration)
	assert.Less(t, timeDiff.Abs(), time.Minute) // Tolérance d'une minute
}

func TestAuthService_CreateToken_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create token success")

	// Créer un utilisateur avec un mot de passe hashé
	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "plainPassword",
		Role:     "user",
		Username: "testuser",
	}

	// Hasher le mot de passe
	err := user.HashPassword()
	require.NoError(t, err)

	loginDto := LoginDto{
		Email:    "test@example.com",
		Password: "plainPassword",
	}

	token, err := testAuthService.CreateToken(user, loginDto)

	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Vérifier que le token est valide
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testJWTKey), nil
	})

	require.NoError(t, err)
	assert.True(t, parsedToken.Valid)
}

func TestAuthService_CreateToken_NilUser(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create token with nil user")

	loginDto := LoginDto{
		Email:    "test@example.com",
		Password: "password",
	}

	token, err := testAuthService.CreateToken(nil, loginDto)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "identifiants invalides")
}

func TestAuthService_CreateToken_WrongPassword(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create token with wrong password")

	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "correctPassword",
		Role:     "user",
		Username: "testuser",
	}

	// Hasher le mot de passe
	err := user.HashPassword()
	require.NoError(t, err)

	loginDto := LoginDto{
		Email:    "test@example.com",
		Password: "wrongPassword",
	}

	token, err := testAuthService.CreateToken(user, loginDto)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "identifiants invalides")
}

func TestAuthService_CreateToken_Integration(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create token integration with real user flow")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Créer un service utilisateur pour les tests
	userService := crm.NewUserService(nil, testDB) // mongoHelper peut être nil pour ces tests

	// Créer un utilisateur
	userCreateDto := crm.UserCreateDto{
		Email:    "integration@example.com",
		Password: "testpassword123",
		Username: "integrationuser",
		Role:     "user",
	}

	createdUser, err := userService.CreateUser(ctx, userCreateDto)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// Récupérer l'utilisateur par email
	foundUser, err := userService.FindByEmail(ctx, userCreateDto.Email)
	require.NoError(t, err)
	require.NotNil(t, foundUser)

	// Créer un token pour cet utilisateur
	loginDto := LoginDto{
		Email:    userCreateDto.Email,
		Password: userCreateDto.Password,
	}

	token, err := testAuthService.CreateToken(foundUser, loginDto)

	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Vérifier le contenu du token
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testJWTKey), nil
	})

	require.NoError(t, err)
	claims, ok := parsedToken.Claims.(*tokenClaims)
	require.True(t, ok)

	assert.Equal(t, foundUser.ID.Hex(), claims.IdUser)
	assert.Equal(t, foundUser.Email, claims.Email)
	assert.Equal(t, foundUser.Role, claims.Role)
}

func TestTokenClaims_Structure(t *testing.T) {
	testup.LogNameTestInfo(t, "Test token claims structure")

	claims := &tokenClaims{
		IdUser: "507f1f77bcf86cd799439011",
		Email:  "test@example.com",
		Role:   "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	assert.Equal(t, "507f1f77bcf86cd799439011", claims.IdUser)
	assert.Equal(t, "test@example.com", claims.Email)
	assert.Equal(t, "admin", claims.Role)
	assert.NotNil(t, claims.ExpiresAt)
	assert.NotNil(t, claims.IssuedAt)
}

func TestLoginDto_Structure(t *testing.T) {
	testup.LogNameTestInfo(t, "Test login DTO structure")

	loginDto := LoginDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	assert.Equal(t, "test@example.com", loginDto.Email)
	assert.Equal(t, "password123", loginDto.Password)
}

func TestAuthService_TokenValidation_Algorithm(t *testing.T) {
	testup.LogNameTestInfo(t, "Test token uses correct signing algorithm")

	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
		Username: "testuser",
	}

	err := user.HashPassword()
	require.NoError(t, err)

	token, err := testAuthService.generateToken(user)
	require.NoError(t, err)

	// Parser sans vérifier la signature d'abord pour examiner l'algorithm
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, &tokenClaims{})
	require.NoError(t, err)

	// Vérifier l'algorithme utilisé
	assert.Equal(t, "HS256", parsedToken.Header["alg"])
	assert.Equal(t, "JWT", parsedToken.Header["typ"])
}

func TestAuthService_MultipleTokensGeneration(t *testing.T) {
	testup.LogNameTestInfo(t, "Test multiple tokens generation")

	user := &crm.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "password",
		Role:     "user",
		Username: "testuser",
	}

	// Générer plusieurs tokens
	var tokens []string
	for i := 0; i < 5; i++ {
		token, err := testAuthService.generateToken(user)
		require.NoError(t, err)
		tokens = append(tokens, token)

		// Petite pause pour s'assurer que les timestamps diffèrent
		time.Sleep(time.Second)
	}

	// Vérifier que tous les tokens sont différents (à cause des timestamps différents)
	for i := 0; i < len(tokens); i++ {
		for j := i + 1; j < len(tokens); j++ {
			assert.NotEqual(t, tokens[i], tokens[j])
		}
	}

	// Vérifier que tous les tokens sont valides
	for _, token := range tokens {
		parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(testJWTKey), nil
		})
		require.NoError(t, err)
		assert.True(t, parsedToken.Valid)
	}
}
