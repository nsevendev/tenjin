package router

import (
	"github.com/gin-gonic/gin"
	"tenjin/back/app/controller/authcontroller"
	"tenjin/back/internal/crm"
)

func RegisterAuth(v1 *gin.RouterGroup, deps *dependencies) {
	userService := crm.NewUserService(deps.MongoHelper, deps.db)
	authController := authcontroller.NewAuthController(userService, deps.authService)

	authGroup := v1.Group("/your")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/logout", authController.Logout)
	}
}
