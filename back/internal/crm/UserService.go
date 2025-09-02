package crm

import (
	"context"
	"errors"
	"fmt"
	"tenjin/back/internal/emailverification"
	"tenjin/back/internal/utils/mongohelpers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongoHelper mongohelpers.Helper
	db          *mongo.Database
}

func NewUserService(mongoHelper mongohelpers.Helper, db *mongo.Database) *UserService {
	return &UserService{
		mongoHelper: mongoHelper,
		db:          db,
	}
}

func (u *UserService) CreateUser(ctx *gin.Context, userCreateDto UserCreateDto) (*User, error) {
	existingUser, err := u.FindByEmail(ctx, userCreateDto.Email)
	if err != nil {
		return nil, fmt.Errorf("impossible de créer votre compte : %v", err)
	}
	if existingUser != nil {
		return nil, errors.New("impossible de créer votre compte")
	}

	user := &User{
		Firstname:            userCreateDto.Firstname,
		Lastname:             userCreateDto.Lastname,
		Email:                userCreateDto.Email,
		Username:             userCreateDto.Username,
		Roles:                userCreateDto.Roles,
		Status:               userCreateDto.Status,
		Password:             userCreateDto.Password,
		Organizations:        userCreateDto.Organizations,
		Sessions:             userCreateDto.Sessions,
		CompetenceRecords:    []CompetenceRecord{},
		ExternalExperiences:  []ExternalExperience{},
		ReceivedOffers:       []primitive.ObjectID{},
		PendingShareRequests: []ShareRequest{},
		QuizResults:          []QuizResult{},
		Chats:                []primitive.ObjectID{},
		CompanyHistory:       []CompanyAssignment{},
		ParticipationHistory: []FormationParticipation{},
		EmailVerified:        false,
	}

	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("un problème est survenue à la création de l'utilisateur : %v", err)
	}

	resutl, err := u.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("un problème est survenue à la création de l'utilisateur : %v", err)
	}

	user.ID = resutl.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (u *UserService) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := u.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Wf("une erreur mongodb est survenue : %v", email)
			return nil, nil
		}
		return nil, fmt.Errorf("erreur à la recuperation du user : %v", err)
	}
	return &user, nil
}

func (u *UserService) VerifyEmail(ctx context.Context, token string, evService *emailverification.EmailVerificationService) error {
	userID, err := evService.VerifyToken(token)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *userID}
	update := bson.M{
		"$set": bson.M{
			"email_verified": true,
			"updated_at":     primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	res, err := u.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de l'utilisateur : %v", err)
	}
	if res.MatchedCount == 0 {
		return errors.New("utilisateur introuvable")
	}

	if err := evService.DeleteTokensByUserID(*userID); err != nil {
		logger.Wf("échec suppression des tokens après vérif email pour user %v : %v", userID.Hex(), err)
	}

	return nil
}