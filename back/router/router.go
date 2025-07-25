package router

import (
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
}
