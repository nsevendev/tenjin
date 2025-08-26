package user

import (
	"context"
	"fmt"

	"tenjin/back/internal/utils/mongohelpers"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	collection  *mongo.Collection
	mongoHelper mongohelpers.Helper
}

type UserServiceInterface interface {
	Create(ctx context.Context, dto UserCreateDto) (*User, error)
}

func NewUserService(db *mongo.Database, helper mongohelpers.Helper) UserServiceInterface {
	return &userService{
		collection:  db.Collection("user"),
		mongoHelper: helper,
	}
}

func (s *userService) Create(ctx context.Context, dto UserCreateDto) (*User, error) {
	user := &User{
		Firstname:     dto.Firstname,
		Lastname:      dto.Lastname,
		Email:         dto.Email,
		Roles:         dto.Roles,
		Organizations: dto.Organizations,
		Sessions:      dto.Sessions,
		Status:        dto.Status,
	}

	s.mongoHelper.SetTimestamps(user)

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Ef("erreur lors de la creation de l'utilisateur : %v", err)
		return nil, fmt.Errorf("erreur lors de la creation de l'utilisateur : %w", err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	logger.Sf("utilisateur cree avec succes : %s %s (%s)", user.Firstname, user.Lastname, user.Email)

	return user, nil
}
