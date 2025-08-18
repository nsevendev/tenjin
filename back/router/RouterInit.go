package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tenjin/back/internal/storage/r2"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/mongohelpers"
)

const pathApiV1 = "api/v1"

// Dependencies contient toutes les dépendances partagées
type Dependencies struct {
	MongoHelper mongohelpers.Helper
	R2Adapter   *r2.Adapter
	// Ajouter d'autres dépendances globales
}

func Routes(r *gin.Engine) {
	fromEnv, err := r2.NewFromEnv()
	if err != nil {
		logger.Ff("Erreur lors de la création de l'adapter R2 : %v", err)
		return
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	deps := &Dependencies{
		MongoHelper: mongohelpers.NewHelper(),
		R2Adapter:   fromEnv,
	}

	v1 := r.Group(pathApiV1)

	RegisterCompanyRoutes(v1, database.Client, deps)
	RegisterUploadFileTest(v1, deps)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
