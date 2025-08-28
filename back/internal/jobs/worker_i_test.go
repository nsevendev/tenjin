//go:build integration

package jobs

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/nsevenpack/testup"
	"github.com/stretchr/testify/assert"

	"tenjin/back/internal/mail"
	"tenjin/back/internal/mailer"
	"tenjin/back/internal/utils/database"
	"tenjin/back/internal/utils/filestores"
	"tenjin/back/internal/utils/mongohelpers"
	"tenjin/back/internal/utils/s3adapter"
)

var (
	mailService *mail.MailService
	mu          *mailer.MailUploader
)

func TestMain(m *testing.M) {
	database.ConnexionDatabase("dev")
	db := database.Client

	if err := db.Collection("mails").Drop(context.Background()); err != nil {
		logger.Ef("Erreur suppression collection 'mails' : %v", err)
		os.Exit(1)
	}

	mailService = mail.NewMailService(mongohelpers.NewHelper(), db)

	fileStoreService := filestores.NewService(s3adapter.AdapterCloudflareR2(), filestores.FileStoreConfig{
		KeyPrefix:      "tests/",
		MaxSize:        0,
		AllowedMIMEs:   []string{},
		UseDateFolders: false,
	})

	mu = &mailer.MailUploader{
		FileStore: fileStoreService,
		MailSvc:   mailService,
	}

	code := m.Run()

	if err := db.Collection("mails").Drop(context.Background()); err != nil {
		logger.Ef("Erreur suppression finale collection 'mails' : %v", err)
	}

	os.Exit(code)
}

func TestWorkerIntegration(t *testing.T) {
	testup.LogNameTestInfo(t, "Test Worker Mail Integration")

	redisAddr := env.Get("REDIS_ADDR")
	if redisAddr == "" {
		t.Fatal("REDIS_ADDR non défini, impossible de tester le worker")
	}
	Redis(redisAddr)

	jobsProcessed := make(chan Job, 10)

	testMailer := mailer.NewMailer(
		env.Get("MAIL_HOST"),
		env.Get("MAIL_PORT"),
		env.Get("MAIL_USER"),
		env.Get("MAIL_PASS"),
		env.Get("MAIL_FROM"),
	)

	StartWorker(testMailer, mu, jobsProcessed)

	jobTest := Job{
		Name: "mail:send",
		Payload: map[string]string{
			"email":   "test@example.com",
			"subject": "Test d'intégration",
			"body":    "ceci est un mail de test",
		},
		Retry:    0,
		MaxRetry: 3,
		Created:  time.Now(),
	}

	ProcessJob(context.Background(), jobTest)

	select {
	case processedJob := <-jobsProcessed:
		assert.Equal(t, jobTest.Name, processedJob.Name)
		assert.Equal(t, jobTest.Payload["email"], processedJob.Payload["email"])
		logger.Sf("✅ Job '%s' traité correctement", processedJob.Name)
	case <-time.After(10 * time.Second):
		t.Fatal("❌ Timeout : le worker n'a pas traité le job dans les 10s")
	}

	logger.I("🔹 Test d’intégration du worker terminé avec succès")
}
