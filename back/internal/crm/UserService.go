package crm

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tenjin/back/internal/utils/mongohelpers"
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
		Email:    userCreateDto.Email,
		Password: userCreateDto.Password,
		Username: userCreateDto.Username,
		Role:     userCreateDto.Role,
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
