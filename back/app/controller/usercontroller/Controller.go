package usercontroller

import (
	"tenjin/back/internal/user"
)

type userController struct {
	userService user.UserServiceInterface
}

type UserControllerInterface interface {
}

func NewUserController(userService user.UserServiceInterface) UserControllerInterface {
	return &userController{
		userService: userService,
	}
}
