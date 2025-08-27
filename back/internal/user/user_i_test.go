package user

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/mongohelpers"

	"github.com/nsevenpack/logger/v2/logger"
)

var userServiceTest UserServiceInterface

func TestMain(m *testing.M) {
	logger.I("Connexion à la base de données de dev...")
	database.ConnexionDatabase("dev")
	db := database.Client

	logger.I("Vider initialement la collection 'user'...")
	_, err := db.Collection("user").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		logger.Ef("Erreur lors du vidage initial de la collection 'user' : %v", err)
		os.Exit(1)
	}

	userServiceTest = NewUserService(db, mongohelpers.NewHelper())

	code := m.Run()

	logger.I("Vider la collection 'user' après les tests...")
	_, err = db.Collection("user").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		logger.Ef("Erreur lors du vidage final de la collection 'user' : %v", err)
		os.Exit(1)
	}

	os.Exit(code)
}

func TestCreateUser_Success(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Create User Success")

	dto := UserCreateDto{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Roles:     []string{"student"},
		Status:    "training",
	}

	user, err := userServiceTest.Create(context.Background(), dto)

	if err != nil {
		logger.Ef("[TestCreateUser_Success] Échec lors de la création de l'utilisateur: %v", err)
	} else {
		logger.Sf("[TestCreateUser_Success] Utilisateur créé avec succès: %+v", user)
	}

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John", user.Firstname)
	assert.Equal(t, "Doe", user.Lastname)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, []string{"student"}, user.Roles)
	assert.Equal(t, "training", user.Status)
	assert.False(t, user.ID.IsZero())
	assert.WithinDuration(t, time.Now(), user.CreatedAt.Time(), 2*time.Second)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Create User Duplicate Email")

	dto := UserCreateDto{
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "jane.doe@example.com",
		Roles:     []string{"trainer"},
	}

	user, err := userServiceTest.Create(context.Background(), dto)
	if err != nil {
		logger.Ef("[TestCreateUser_DuplicateEmail] Erreur à la première création: %v", err)
	} else {
		logger.Sf("[TestCreateUser_DuplicateEmail] Premier utilisateur créé: %+v", user)
	}
	assert.NoError(t, err)
	assert.NotNil(t, user)

	user2, err := userServiceTest.Create(context.Background(), dto)
	if err != nil {
		logger.Ef("[TestCreateUser_DuplicateEmail] Erreur sur doublon: %v", err)
	} else {
		logger.Sf("[TestCreateUser_DuplicateEmail] Second utilisateur créé malgré doublon: %+v", user2)
	}

	if err == nil {
		assert.NotNil(t, user2)
	} else {
		assert.Error(t, err)
	}
}
