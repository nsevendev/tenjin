//go:build integration

package jobs

import (
	"context"
	"testing"
	"time"

	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	"github.com/stretchr/testify/assert"
)

func TestWorkerIntegration(t *testing.T) {
	logger.I("🔹 Démarrage du test d’intégration du worker...")

	redisAddr := env.Get("REDIS_ADDR")
	if redisAddr == "" {
		t.Fatal("REDIS_ADDR non défini, impossible de tester le worker")
	}

	Redis(redisAddr)

	JobsProcessed = make(chan Job, 10)

	StartWorker()

	jobTest := Job{
		Name: "test:integration",
		Payload: map[string]string{
			"message": "Hello Redis Worker!",
		},
		Retry:    0,
		MaxRetry: 3,
		Created:  time.Now(),
	}

	ProcessJob(context.Background(), jobTest)

	select {
	case processedJob := <-JobsProcessed:
		assert.Equal(t, jobTest.Name, processedJob.Name)
		assert.Equal(t, jobTest.Payload["message"], processedJob.Payload["message"])
		logger.Sf("✅ Job '%s' traité correctement", processedJob.Name)
	case <-time.After(5 * time.Second):
		t.Fatal("❌ Timeout : le worker n'a pas traité le job dans les 5s")
	}

	logger.I("🔹 Test d’intégration du worker terminé avec succès")
}
