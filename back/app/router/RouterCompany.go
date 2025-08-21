package router

import (
	"github.com/gin-gonic/gin"
	"tenjin/back/app/controller/companycontroller"
	"tenjin/back/internal/company"
)

func RegisterCompanyRoutes(v1 *gin.RouterGroup, deps *dependencies) {
	companyService := company.NewCompanyService(deps.db, deps.MongoHelper)
	companyController := companycontroller.NewCompanyController(companyService)

	companyGroup := v1.Group("/company")
	{
		companyGroup.POST("/retrieve-infos", companyController.RetrieveCompanyInfo)
		companyGroup.POST("/register", companyController.Create)
	}
}
