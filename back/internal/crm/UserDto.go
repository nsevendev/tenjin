package crm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCreateDto struct {
	Firstname     string               `json:"firstname" validate:"required,min=2,max=100"`
	Lastname      string               `json:"lastname" validate:"required,min=2,max=100"`
	Email         string               `json:"email" validate:"required,email"`
	Password      string               `json:"password" validate:"required,min=6"`
	Username      string               `json:"username" validate:"required"`
	Roles         []string             `json:"roles" validate:"required,dive,oneof=student trainer manager admin recruiter"`
	Status        string               `json:"status" validate:"omitempty,oneof=training employed jobseeker"`
	Organizations []primitive.ObjectID `json:"organizations"`
	Sessions      []primitive.ObjectID `json:"sessions"`
}

type UserDeleteDto struct {
	Ids []primitive.ObjectID `json:"ids" binding:"required"`
}

func (u *UserService) CreateDtoFaker() UserCreateDto {
	timestamp := time.Now().Format("2006-01-02_15-04-05.000")
	return UserCreateDto{
		Firstname:     "Firstname" + timestamp,
		Lastname:      "Lastname" + timestamp,
		Email:         "user" + timestamp + "@example.com",
		Password:      "password123",
		Username:      "user" + timestamp,
		Roles:         []string{"user"},
		Status:        "employee",
		Organizations: []primitive.ObjectID{},
		Sessions:      []primitive.ObjectID{},
	}
}


func (u *UserService) CreateDtosFaker(n int) []UserCreateDto {
	users := make([]UserCreateDto, n)
	for i := 0; i < n; i++ {
		users[i] = u.CreateDtoFaker()
	}
	return users
}
