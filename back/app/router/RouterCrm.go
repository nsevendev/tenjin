package router

import (
	"tenjin/back/app/controller/crmcontroller"
	"tenjin/back/internal/crm"
	"tenjin/back/internal/emailverification"

	"github.com/gin-gonic/gin"
)

func RegisterCrm(v1 *gin.RouterGroup, deps *dependencies) {
	userService := crm.NewUserService(deps.MongoHelper, deps.db)
	emailVerificationService := emailverification.NewEmailVerificationService(deps.MongoHelper, deps.db)
	crmController := crmcontroller.NewCrmController(userService, emailVerificationService)

	crmGroup := v1.Group("/user")
	{
		crmGroup.POST("/register", crmController.RegisterUser)
	}
}
