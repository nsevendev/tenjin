package crm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCreateDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"omitempty"`
}

type UserDeleteDto struct {
	Ids []primitive.ObjectID `json:"ids" binding:"required"`
}

func (u *UserService) CreateDtoFaker() UserCreateDto {
	return UserCreateDto{
		Email:    "user" + time.Now().Format("2006-01-02_15-04-05.000") + "@example.com",
		Password: "password",
		Username: "user" + time.Now().Format("2006-01-02_15-04-05.000"),
		Role:     "user",
	}
}

func (u *UserService) CreateDtosFaker(n int) []UserCreateDto {
	users := make([]UserCreateDto, n)
	for i := 0; i < n; i++ {
		users[i] = u.CreateDtoFaker()
	}
	return users
}
