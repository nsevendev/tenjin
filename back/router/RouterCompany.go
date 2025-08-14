package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"tenjin/back/controller/companycontroller"
	"tenjin/back/internal/company"
)

func RegisterCompanyRoutes(v1 *gin.RouterGroup, db *mongo.Database, deps *Dependencies) {
	companyService := company.NewCompanyService(db, deps.MongoHelper)
	companyController := companycontroller.NewCompanyController(companyService)

	companyGroup := v1.Group("/company")
	{
		companyGroup.POST("/retrieve-infos", companyController.RetrieveCompanyInfo)
		companyGroup.POST("/register", companyController.Create)
	}
}
