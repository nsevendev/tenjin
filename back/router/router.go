package router

import (
	"tenjin/back/controller/companycontroller"
	"tenjin/back/internal/company"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	initializer "tenjin/back/internal/utils/db"
)

const pathApiV1 = "api/v1"

func Routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	companyService := company.NewCompanyService(initializer.Db)
	companyController := companycontroller.NewCompanyController(companyService)

	v1 := r.Group(pathApiV1)

	v1Company := v1.Group("/company")
	v1Company.POST("/retrieve-infos", companyController.RetrieveCompanyInfo)
	v1Company.POST("/register", companyController.Create)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
