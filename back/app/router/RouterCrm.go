package router

import (
	"github.com/gin-gonic/gin"
	"tenjin/back/app/controller/crmcontroller"
	"tenjin/back/internal/crm"
)

func RegisterCrm(v1 *gin.RouterGroup, deps *dependencies) {
	userService := crm.NewUserService(deps.MongoHelper, deps.db)
	crmController := crmcontroller.NewCrmController(userService)

	crmGroup := v1.Group("/user")
	{
		crmGroup.POST("/register", crmController.RegisterUser)
	}
}
