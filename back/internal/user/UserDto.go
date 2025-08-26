package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCreateDto struct {
	Firstname     string               `bson:"firstname" json:"firstname" validate:"required,min=2,max=100"`
	Lastname      string               `bson:"lastname" json:"lastname" validate:"required,min=2,max=100"`
	Email         string               `bson:"email" json:"email" validate:"required,email"`
	Roles         []string             `bson:"roles" json:"roles" validate:"required,dive,oneof=student trainer manager admin recruiter"`
	Organizations []primitive.ObjectID `bson:"organizations" json:"organizations"`
	Sessions      []primitive.ObjectID `bson:"sessions" json:"sessions"`
	Status        string               `bson:"status" json:"status" validate:"omitempty,oneof=training employed jobseeker"`
}