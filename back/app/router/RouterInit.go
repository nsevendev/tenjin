package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"tenjin/back/internal/auth"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/mongohelpers"
	s3adapter2 "tenjin/back/internal/utils/s3adapter"
)

const pathApiV1 = "api/v1"

// Dependencies contient toutes les dépendances partagées
type dependencies struct {
	MongoHelper mongohelpers.Helper
	R2Adapter   s3adapter2.AdapterInterface
	db          *mongo.Database
	authService *auth.AuthService
}

func Routes(r *gin.Engine) {
	deps := &dependencies{
		MongoHelper: mongohelpers.NewHelper(),
		R2Adapter:   s3adapter2.AdapterCloudflareR2(),
		db:          database.Client,
		authService: auth.NewAuthService(database.Client, env.Get("JWT_SECRET_KEY")),
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group(pathApiV1)

	RegisterCompanyRoutes(v1, deps)
	RegisterUploadFileTest(v1, deps)
	RegisterAuth(v1, deps)
	RegisterCrm(v1, deps)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
