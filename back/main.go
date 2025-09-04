package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/mignosql"
	"strings"
	"tenjin/back/app/router"
	"tenjin/back/docs"
	"tenjin/back/internal/jobs"
	"tenjin/back/internal/mail"
	"tenjin/back/internal/mailer"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/filestores"
	"tenjin/back/internal/utils/s3adapter"
	"tenjin/back/migration"
	"time"
)

func init() {
	appEnv := env.Get("APP_ENV")
	logger.Init(appEnv)
	initDbAndMigNosql(appEnv)

	ginresponse.SetFormatter(&ginresponse.JsonFormatter{})
	s3adapter.CreateAdapteur()

	jobsProcessed := make(chan jobs.Job, 100)

	mailerInstance := mailer.NewMailer(
		env.Get("MAILTRAP_HOST"),
		env.Get("MAILTRAP_PORT"),
		env.Get("MAILTRAP_USER"),
		env.Get("MAILTRAP_PASS"),
		env.Get("MAIL_FROM"),
	)

	mailService := mail.NewMailService(nil, database.Client)
	fileStoreService := filestores.NewService(
		s3adapter.AdapterCloudflareR2(),
		filestores.FileStoreConfig{
			KeyPrefix:      "mails/",
			MaxSize:        0,
			AllowedMIMEs:   []string{},
			UseDateFolders: true,
		},
	)
	mu := &mailer.MailUploader{
		FileStore: fileStoreService,
		MailSvc:   mailService,
	}

	jobs.InitJobs(mailerInstance, mu, jobsProcessed)
	mailer.InitMailer()
}

// @title tenjin api
// @version 1.0
// @description API service tenjin api
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in headers
// @name Authorization
func main() {
	s := gin.Default()
	host := "0.0.0.0"
	hostTraefikApi := extractStringInBacktick(env.Get("HOST_TRAEFIK_API"))
	hostTraefikDb := extractStringInBacktick(env.Get("HOST_TRAEFIK_DB"))
	port := env.Get("PORT")
	setSwaggerOpt(hostTraefikApi)             // config option swagger
	infoServer(hostTraefikApi, hostTraefikDb) // log info server
	setCors(s)

	router.Routes(s)

	if err := s.Run(host + ":" + port); err != nil {
		logger.Ef("Une erreur est survenue au lancement du serveur : %v", err)
	}
}

func setCors(s *gin.Engine) {
	s.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://tenjin-app.local",
			"https://tenjin-app.woopear.fr",
		}, // front autorisés
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		}, // méthodes autorisées
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"X-Requested-With",
		},                                           // headers autorisés
		ExposeHeaders:    []string{"X-Total-Count"}, // headers exposés au client
		AllowCredentials: true,                      // true si utilise cookies ou fetch withCredentials
		MaxAge:           12 * time.Hour,            // cache du preflight
	}))
}

func infoServer(hostTraefikApi string, hostTraefikDb string) {
	logger.If("Lancement du serveur : https://%v", hostTraefikApi)
	logger.If("Lancement du Swagger : https://%v/swagger/index.html", hostTraefikApi)
}

func extractStringInBacktick(s string) string {
	start := strings.Index(s, "`")
	end := strings.LastIndex(s, "`")

	if start == -1 || end == -1 || start == end {
		return ""
	}

	return s[start+1 : end]
}

func setSwaggerOpt(hostTraefikApi string) {
	docs.SwaggerInfo.Host = hostTraefikApi
}

func initDbAndMigNosql(appEnv string) {
	database.ConnexionDatabase(appEnv)
	migrator := mignosql.New(database.Client)
	// EXAMPLE => migrator.Add(migration.<namefile>)
	// ajouter les migrations ici ...
	migrator.Add(migration.CreateCompanyCollection)

	if err := migrator.Apply(); err != nil {
		logger.Ff("Erreur lors de l'application des migrations : %v", err)
	}
}
