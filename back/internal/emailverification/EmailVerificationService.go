package emailverification

import (
	"context"
	"errors"
	"fmt"
	"tenjin/back/internal/utils/mongohelpers"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *EmailVerificationService) VerifyToken(token string) (*primitive.ObjectID, error) {
	var ev EmailVerification

	err := s.db.Collection("email_verifications").
		FindOne(context.Background(), bson.M{"token": token}).
		Decode(&ev)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("token invalide")
		}
		return nil, fmt.Errorf("erreur lors de la recherche du token : %v", err)
	}

	if time.Now().After(ev.Expiry) {
		_, _ = s.db.Collection("email_verifications").
			DeleteOne(context.Background(), bson.M{"_id": ev.ID})
		return nil, errors.New("token expiré")
	}

	_, _ = s.db.Collection("email_verifications").
		DeleteOne(context.Background(), bson.M{"_id": ev.ID})

	return &ev.UserID, nil
}

func (s *EmailVerificationService) DeleteTokensByUserID(userID primitive.ObjectID) error {
	_, err := s.db.Collection("email_verifications").DeleteMany(
		context.Background(),
		map[string]interface{}{
			"user_id": userID,
		},
	)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression des tokens de l'utilisateur %v : %w", userID.Hex(), err)
	}
	return nil
}