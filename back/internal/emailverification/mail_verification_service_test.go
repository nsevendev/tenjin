package emailverification

import (
	"context"
	"os"
	"testing"
	"time"

	"tenjin/back/internal/utils/database"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	testEmailVerificationService *EmailVerificationService
	testDB                       *mongo.Database
)

func TestMain(m *testing.M) {
	appEnv := env.Get("APP_ENV")

	database.ConnexionDatabase(appEnv)
	testDB = database.Client

	err := testDB.Collection("email_verifications").Drop(context.Background())
	if err != nil {
		logger.Ef("Erreur lors de la suppression de la collection 'email_verifications' : %v", err)
		return
	}

	testEmailVerificationService = NewEmailVerificationService(nil, testDB)

	code := m.Run()

	logger.If("Tests EmailVerificationService terminés")
	os.Exit(code)
}

func TestEmailVerificationService_GenerateToken(t *testing.T) {
	testup.LogNameTestInfo(t, "Test GenerateToken")

	userID := primitive.NewObjectID()
	token, err := testEmailVerificationService.GenerateToken(userID)

	require.NoError(t, err)
	assert.NotEmpty(t, token, "Le token généré ne doit pas être vide")

	var ev EmailVerification
	err = testDB.Collection("email_verifications").FindOne(context.Background(), map[string]interface{}{
		"user_id": userID,
		"token":   token,
	}).Decode(&ev)

	require.NoError(t, err)
	assert.Equal(t, userID, ev.UserID)
	assert.Equal(t, token, ev.Token)

	now := time.Now()
	assert.True(t, ev.Expiry.After(now), "Expiry should be after current time")
	assert.InDelta(t, now.Add(1*time.Hour).Unix(), ev.Expiry.Unix(), 5, "Expiry should be approximately 1 hour from now")

	assert.WithinDuration(t, time.Now(), ev.CreatedAt, 5*time.Second, "CreatedAt should be close to now")
}