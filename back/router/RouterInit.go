package router

import (
	"tenjin/back/internal/utils/db"
	"tenjin/back/internal/utils/mongoapp"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const pathApiV1 = "api/v1"

// Dependencies contient toutes les dépendances partagées
type Dependencies struct {
	MongoHelper mongoapp.Helper
	// Ajouter d'autres dépendances globales
}

func Routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	deps := &Dependencies{
		MongoHelper: mongoapp.NewHelper(),
	}

	v1 := r.Group(pathApiV1)

	RegisterCompanyRoutes(v1, db.Client, deps)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
