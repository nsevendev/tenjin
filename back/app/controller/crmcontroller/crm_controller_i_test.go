//go:build integration

package crmcontroller

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"tenjin/back/internal/crm"
	"tenjin/back/internal/emailverification"
	"tenjin/back/internal/jobs"
	"tenjin/back/internal/mail"
	"tenjin/back/internal/mailer"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/filestores"
	"tenjin/back/internal/utils/mongohelpers"
	"tenjin/back/internal/utils/s3adapter"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/ginresponse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
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
	_ = testDB.Collection("mails").Drop(nil)

	testUserService = crm.NewUserService(nil, testDB)
	testEmailVerificationService = emailverification.NewEmailVerificationService(nil, testDB)
	testCrmController = &CrmController{
		userService:              testUserService,
		emailVerificationService: testEmailVerificationService,
	}

	ginresponse.SetFormatter(&ginresponse.JsonFormatter{})

	redisAddr := env.Get("REDIS_ADDR")
	jobs.Redis(redisAddr)

	mailerInstance := mailer.NewMailer(
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_PORT"),
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASS"),
		os.Getenv("MAIL_FROM"),
	)

	mailService := mail.NewMailService(mongohelpers.NewHelper(), testDB)
	s3adapter.CreateAdapteur()

	fileStore := filestores.NewService(
		s3adapter.AdapterCloudflareR2(),
		filestores.FileStoreConfig{
			KeyPrefix:      "mails/",
			MaxSize:        0,
			AllowedMIMEs:   []string{},
			UseDateFolders: true,
		},
	)
	mu := &mailer.MailUploader{
		FileStore: fileStore,
		MailSvc:   mailService,
	}

	jobsProcessed := make(chan jobs.Job, 10)
	jobs.StartWorker(mailerInstance, mu, jobsProcessed)

	_jobsProcessed = jobsProcessed

	code := m.Run()
	os.Exit(code)
}

var _jobsProcessed chan jobs.Job

func TestRegisterUser_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	dto := testUserService.CreateDtoFaker()
	body := fmt.Sprintf(
		`{"firstname":"%s","lastname":"%s","email":"%s","username":"%s","password":"%s"}`,
		dto.Firstname, dto.Lastname, dto.Email, dto.Username, dto.Password,
	)
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	testCrmController.RegisterUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Utilisateur créé avec succès")

	user, err := testUserService.FindByEmail(context.Background(), dto.Email)
	require.NoError(t, err)
	require.NotNil(t, user)
	assert.False(t, user.EmailVerified, "L'email doit être non vérifié à la création")

	var ev emailverification.EmailVerification
	err = testDB.Collection("email_verifications").FindOne(context.Background(), bson.M{"user_id": user.ID}).Decode(&ev)
	require.NoError(t, err)
	assert.NotEmpty(t, ev.Token, "Un token de vérification doit être généré")
	assert.True(t, ev.Expiry.After(time.Now()), "Le token doit être valide (non expiré)")

	select {
	case job := <-_jobsProcessed:
		assert.Equal(t, "mail:send", job.Name)
		assert.Equal(t, dto.Email, job.Payload["email"])

		var mailRecord mail.Mail
		err := testDB.Collection("mails").FindOne(
			context.Background(),
			bson.M{"to": dto.Email},
		).Decode(&mailRecord)
		require.NoError(t, err)
		assert.NotNil(t, mailRecord.S3Path, "Le mail doit être stocké en R2")

	case <-time.After(10 * time.Second):
		t.Fatal("Le worker n'a pas traité le job à temps")
	}
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
