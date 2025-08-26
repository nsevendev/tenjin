package jobs

import (
	"context"
	"testing"
	"time"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/stretchr/testify/assert"

	"tenjin/back/internal/mailer"
)

func TestWorkerIntegration(t *testing.T) {

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
	StartWorker(testMailer, jobsProcessed)

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
