//go:build integration

package crmcontroller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"tenjin/back/internal/crm"
	"tenjin/back/internal/emailverification"
	"tenjin/back/internal/utils/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/ginresponse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	testCrmController            *CrmController
	testUserService              *crm.UserService
	testEmailVerificationService *emailverification.EmailVerificationService
	testDB                       *mongo.Database
)

func TestMain(m *testing.M) {
	appEnv := env.Get("APP_ENV")

	database.ConnexionDatabase(appEnv)
	testDB = database.Client

	_ = testDB.Collection("users").Drop(nil)
	_ = testDB.Collection("email_verifications").Drop(nil)

	testUserService = crm.NewUserService(nil, testDB)
	testEmailVerificationService = emailverification.NewEmailVerificationService(nil, testDB)
	testCrmController = &CrmController{
		userService:             testUserService,
		emailVerificationService: testEmailVerificationService,
	}

	ginresponse.SetFormatter(&ginresponse.JsonFormatter{})

	code := m.Run()
	os.Exit(code)
}

func TestVerifyUserEmail_Success(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	dto := testUserService.CreateDtoFaker()
	createdUser, err := testUserService.CreateUser(ctx, dto)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	token, err := testEmailVerificationService.GenerateToken(createdUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	req := httptest.NewRequest("GET", fmt.Sprintf("/user/verify-email?token=%s", token), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	testCrmController.VerifyUserEmail(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Email vérifié avec succès")
}

func TestVerifyUserEmail_MissingToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/user/verify-email", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	testCrmController.VerifyUserEmail(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Token manquant")
}

func TestVerifyUserEmail_InvalidToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/user/verify-email?token=invalidtoken", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	testCrmController.VerifyUserEmail(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Échec de la vérification de l'email")
}

func TestVerifyUserEmail_ExpiredToken(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	dto := testUserService.CreateDtoFaker()
	createdUser, err := testUserService.CreateUser(ctx, dto)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	expiredToken := uuid.New().String()
	expiredVerification := emailverification.EmailVerification{
		UserID:    createdUser.ID,
		Token:     expiredToken,
		Expiry:    time.Now().Add(-1 * time.Hour),
		CreatedAt: time.Now().Add(-2 * time.Hour),
	}

	_, err = testDB.Collection("email_verifications").InsertOne(ctx, expiredVerification)
	require.NoError(t, err)

	req := httptest.NewRequest("GET", fmt.Sprintf("/user/verify-email?token=%s", expiredToken), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	testCrmController.VerifyUserEmail(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Échec de la vérification de l'email")
}
