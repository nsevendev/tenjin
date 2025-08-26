package usercontroller

import (
	"tenjin/back/internal/user"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.UserServiceInterface
}

type UserControllerInterface interface {
	Create(c *gin.Context)
}

func NewUserController(userService user.UserServiceInterface) UserControllerInterface {
	return &userController{
		userService: userService,
	}
}
