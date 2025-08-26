package router

import (
	"tenjin/back/app/controller/usercontroller"
	"tenjin/back/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(v1 *gin.RouterGroup, deps *dependencies) {
	userService := user.NewUserService(deps.db, deps.MongoHelper)
	userController := usercontroller.NewUserController(userService)

	userGroup := v1.Group("/user")
	{
		userGroup.POST("/register", userController.Create)
	}
}
