package router

import (
	"tenjin/back/internal/controller/companycontroller"
	"tenjin/back/internal/insee"

	initializer "tenjin/back/internal/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const pathApiV1 = "api/v1"

func Routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	companyService := insee.NewCompanyService(initializer.Db)
	companyController := companycontroller.NewCompanyController(companyService)

	v1 := r.Group(pathApiV1)

	v1Company := v1.Group("/company")
	v1Company.POST("/create", companyController.Create)
}
