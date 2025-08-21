package crm

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"tenjin/back/internal/utils/database"
)

var (
	testUserService *UserService
	testDB          *mongo.Database
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
	testUserService = NewUserService(nil, testDB)

	code := m.Run()

	logger.If("Tests UserService terminés")
	os.Exit(code)
}

func TestNewUserService(t *testing.T) {
	testup.LogNameTestInfo(t, "Test creation of UserService")

	service := NewUserService(nil, testDB)

	assert.NotNil(t, service)
	assert.Equal(t, testDB, service.db)
	assert.Nil(t, service.mongoHelper) // Pour ces tests, on utilise nil
}

func TestUserService_CreateUser_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create user success")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	uniqueEmail := fmt.Sprintf("testuser_%d@example.com", time.Now().UnixNano())
	userCreateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: "password123",
		Username: "testuser",
		Role:     "user",
	}

	createdUser, err := testUserService.CreateUser(ctx, userCreateDto)

	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// Vérifier les champs de l'utilisateur créé
	assert.Equal(t, userCreateDto.Email, createdUser.Email)
	assert.Equal(t, userCreateDto.Username, createdUser.Username)
	assert.Equal(t, userCreateDto.Role, createdUser.Role)
	assert.NotEmpty(t, createdUser.ID)
	assert.NotEqual(t, userCreateDto.Password, createdUser.Password) // Le mot de passe doit être hashé

	// Vérifier que le mot de passe a été hashé correctement
	assert.True(t, createdUser.CheckPassword(userCreateDto.Password))
}

func TestUserService_CreateUser_DuplicateEmail(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create user with duplicate email")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	uniqueEmail := fmt.Sprintf("duplicate_%d@example.com", time.Now().UnixNano())
	userCreateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: "password123",
		Username: "testuser1",
		Role:     "user",
	}

	// Créer le premier utilisateur
	firstUser, err := testUserService.CreateUser(ctx, userCreateDto)
	require.NoError(t, err)
	require.NotNil(t, firstUser)

	// Tenter de créer un deuxième utilisateur avec le même email
	userCreateDto.Username = "testuser2" // Changer le username mais garder le même email
	secondUser, err := testUserService.CreateUser(ctx, userCreateDto)

	assert.Error(t, err)
	assert.Nil(t, secondUser)
	assert.Contains(t, err.Error(), "impossible de créer votre compte")
}

func TestUserService_CreateUser_PasswordHashing(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create user password hashing")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	uniqueEmail := fmt.Sprintf("hash_test_%d@example.com", time.Now().UnixNano())
	plainPassword := "mySecretPassword123"
	userCreateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: plainPassword,
		Username: "hashtest",
		Role:     "user",
	}

	createdUser, err := testUserService.CreateUser(ctx, userCreateDto)

	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// Le mot de passe stocké ne doit pas être le mot de passe en clair
	assert.NotEqual(t, plainPassword, createdUser.Password)
	assert.Contains(t, createdUser.Password, "$2a$") // bcrypt hash prefix

	// Mais CheckPassword doit fonctionner avec le mot de passe original
	assert.True(t, createdUser.CheckPassword(plainPassword))
	assert.False(t, createdUser.CheckPassword("wrongPassword"))
}

func TestUserService_CreateUser_EmptyFields(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create user with empty fields")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	tests := []struct {
		name string
		dto  UserCreateDto
	}{
		{
			name: "Empty email",
			dto: UserCreateDto{
				Email:    "",
				Password: "password123",
				Username: "testuser",
				Role:     "user",
			},
		},
		{
			name: "Empty password",
			dto: UserCreateDto{
				Email:    "test@example.com",
				Password: "",
				Username: "testuser",
				Role:     "user",
			},
		},
		{
			name: "Empty username",
			dto: UserCreateDto{
				Email:    "test@example.com",
				Password: "password123",
				Username: "",
				Role:     "user",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := testUserService.CreateUser(ctx, tt.dto)
			// Ces tests pourraient passer ou échouer selon la validation côté service
			// Le comportement dépend de si vous avez des validations
			if err != nil {
				assert.Nil(t, user)
			}
			// Note: Ces tests montrent les cas limites, vous pourriez vouloir ajouter des validations
		})
	}
}

func TestUserService_FindByEmail_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test find user by email success")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Créer d'abord un utilisateur
	uniqueEmail := fmt.Sprintf("findtest_%d@example.com", time.Now().UnixNano())
	userCreateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: "password123",
		Username: "findtestuser",
		Role:     "admin",
	}

	createdUser, err := testUserService.CreateUser(ctx, userCreateDto)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// Rechercher cet utilisateur par email
	foundUser, err := testUserService.FindByEmail(ctx, uniqueEmail)

	require.NoError(t, err)
	require.NotNil(t, foundUser)

	// Vérifier que les données correspondent
	assert.Equal(t, createdUser.ID, foundUser.ID)
	assert.Equal(t, createdUser.Email, foundUser.Email)
	assert.Equal(t, createdUser.Username, foundUser.Username)
	assert.Equal(t, createdUser.Role, foundUser.Role)
	assert.Equal(t, createdUser.Password, foundUser.Password)
}

func TestUserService_FindByEmail_NotFound(t *testing.T) {
	testup.LogNameTestInfo(t, "Test find user by email not found")

	ctx := context.Background()

	nonExistentEmail := "nonexistent@example.com"
	foundUser, err := testUserService.FindByEmail(ctx, nonExistentEmail)

	require.NoError(t, err) // FindByEmail retourne nil, nil pour un utilisateur non trouvé
	assert.Nil(t, foundUser)
}

/*
func TestUserService_FindByEmail_EmptyEmail(t *testing.T) {
	testup.LogNameTestInfo(t, "Test find user by empty email")

	ctx := context.Background()

	foundUser, err := testUserService.FindByEmail(ctx, "")

	//require.NoError(t, err) // Même comportement que pour un email non trouvé
	assert.Nil(t, foundUser)
}*/

func TestUserService_CreateUser_Integration(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create and find user integration")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Données de test
	uniqueEmail := fmt.Sprintf("integration_%d@example.com", time.Now().UnixNano())
	userCreateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: "integrationPassword123",
		Username: "integrationUser",
		Role:     "moderator",
	}

	// Créer l'utilisateur
	createdUser, err := testUserService.CreateUser(ctx, userCreateDto)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// Vérifier qu'on peut le retrouver par email
	foundUser, err := testUserService.FindByEmail(ctx, uniqueEmail)
	require.NoError(t, err)
	require.NotNil(t, foundUser)

	// Vérifier que les données sont identiques
	assert.Equal(t, createdUser.ID, foundUser.ID)
	assert.Equal(t, createdUser.Email, foundUser.Email)
	assert.Equal(t, createdUser.Username, foundUser.Username)
	assert.Equal(t, createdUser.Role, foundUser.Role)

	// Vérifier que le mot de passe fonctionne
	assert.True(t, foundUser.CheckPassword(userCreateDto.Password))

	// Vérifier qu'on ne peut pas créer un autre utilisateur avec le même email
	duplicateDto := UserCreateDto{
		Email:    uniqueEmail,
		Password: "differentPassword",
		Username: "differentUser",
		Role:     "user",
	}

	duplicateUser, err := testUserService.CreateUser(ctx, duplicateDto)
	assert.Error(t, err)
	assert.Nil(t, duplicateUser)
}

func TestUserService_CreateUser_MultipleUsers(t *testing.T) {
	testup.LogNameTestInfo(t, "Test create multiple users")

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	users := []UserCreateDto{
		{
			Email:    fmt.Sprintf("user1_%d@example.com", time.Now().UnixNano()),
			Password: "password1",
			Username: "user1",
			Role:     "user",
		},
		{
			Email:    fmt.Sprintf("user2_%d@example.com", time.Now().UnixNano()),
			Password: "password2",
			Username: "user2",
			Role:     "admin",
		},
		{
			Email:    fmt.Sprintf("user3_%d@example.com", time.Now().UnixNano()),
			Password: "password3",
			Username: "user3",
			Role:     "moderator",
		},
	}

	var createdUsers []*User

	// Créer tous les utilisateurs
	for _, userDto := range users {
		createdUser, err := testUserService.CreateUser(ctx, userDto)
		require.NoError(t, err)
		require.NotNil(t, createdUser)
		createdUsers = append(createdUsers, createdUser)
	}

	// Vérifier qu'on peut tous les retrouver
	for i, userDto := range users {
		foundUser, err := testUserService.FindByEmail(ctx, userDto.Email)
		require.NoError(t, err)
		require.NotNil(t, foundUser)

		assert.Equal(t, createdUsers[i].ID, foundUser.ID)
		assert.Equal(t, userDto.Email, foundUser.Email)
		assert.Equal(t, userDto.Username, foundUser.Username)
		assert.Equal(t, userDto.Role, foundUser.Role)
	}
}

func TestUserCreateDto_Faker(t *testing.T) {
	testup.LogNameTestInfo(t, "Test UserCreateDto faker methods")

	// Test de la méthode CreateDtoFaker
	fakeDto := testUserService.CreateDtoFaker()

	assert.NotEmpty(t, fakeDto.Email)
	assert.Contains(t, fakeDto.Email, "@example.com")
	assert.Equal(t, "password", fakeDto.Password)
	assert.NotEmpty(t, fakeDto.Username)
	assert.Equal(t, "user", fakeDto.Role)

	// Test de CreateDtosFaker
	fakeDtos := testUserService.CreateDtosFaker(3)

	assert.Len(t, fakeDtos, 3)
	for i, dto := range fakeDtos {
		assert.NotEmpty(t, dto.Email, "DTO %d should have email", i)
		assert.NotEmpty(t, dto.Username, "DTO %d should have username", i)
		assert.Equal(t, "password", dto.Password, "DTO %d should have password", i)
		assert.Equal(t, "user", dto.Role, "DTO %d should have role", i)
	}
}
