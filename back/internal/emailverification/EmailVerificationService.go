package emailverification

import (
	"context"
	"tenjin/back/internal/utils/mongohelpers"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailVerificationService struct {
	mongoHelper mongohelpers.Helper
	db          *mongo.Database
}

func NewEmailVerificationService(mongoHelper mongohelpers.Helper, db *mongo.Database) *EmailVerificationService {
	return &EmailVerificationService{
		mongoHelper: mongoHelper,
		db:          db,
	}
}

func (s *EmailVerificationService) GenerateToken(userID primitive.ObjectID) (string, error) {
	token := uuid.New().String()
	expiry := time.Now().Add(1 * time.Hour)

	ev := EmailVerification{
		UserID:    userID,
		Token:     token,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}

	_, err := s.db.Collection("email_verifications").InsertOne(context.Background(), ev)
	if err != nil {
		return "", err
	}

	return token, nil
}