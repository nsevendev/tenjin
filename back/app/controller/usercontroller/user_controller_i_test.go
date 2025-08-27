package usercontroller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"tenjin/back/internal/user"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/mongohelpers"
)

var (
	userServiceTest   user.UserServiceInterface
	userControllerTest UserControllerInterface
	router             *gin.Engine
)

func TestMain(m *testing.M) {
	ginresponse.SetFormatter(&ginresponse.JsonFormatter{})
	database.ConnexionDatabase("dev")
	db := database.Client

	if _, err := db.Collection("user").DeleteMany(context.Background(), bson.M{}); err != nil {
		logger.Ef("Erreur lors du vidage initial de la collection 'user' : %v", err)
		os.Exit(1)
	}

	userServiceTest = user.NewUserService(db, mongohelpers.NewHelper())
	userControllerTest = NewUserController(userServiceTest)

	gin.SetMode(gin.TestMode)
	router = gin.New()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userControllerTest.Create)
	}

	code := m.Run()

	if _, err := db.Collection("user").DeleteMany(context.Background(), bson.M{}); err != nil {
		logger.Ef("Erreur lors du vidage final de la collection 'user' : %v", err)
		os.Exit(1)
	}

	os.Exit(code)
}

func TestCreateUser_Success(t *testing.T) {
	dto := user.UserCreateDto{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Roles:     []string{"student"},
		Status:    "training",
	}

	body, _ := json.Marshal(dto)
	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Utilisateur cree avec succes", responseBody["message"])

	data := responseBody["data"].(map[string]interface{})
	assert.Equal(t, dto.Firstname, data["firstname"])
	assert.Equal(t, dto.Lastname, data["lastname"])
	assert.Equal(t, dto.Email, data["email"])
	assert.Equal(t, dto.Status, data["status"])
	assert.ElementsMatch(t, dto.Roles, toStringSlice(data["roles"]))
}

func TestCreateUser_InvalidPayload(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Erreur de validation", responseBody["message"])
}

func toStringSlice(input interface{}) []string {
	raw, ok := input.([]interface{})
	if !ok {
		return nil
	}
	out := make([]string, len(raw))
	for i, v := range raw {
		out[i] = v.(string)
	}
	return out
}
