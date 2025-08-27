package jobs

import (
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
)

func InitJobs() {
	redisAddr := env.Get("REDIS_ADDR")
	if redisAddr == "" {
		logger.Ef("❌ REDIS_ADDR non défini, impossible de démarrer le worker")
		return
	}

	Redis(redisAddr)
	logger.Sf("✅ Redis initialisé sur %s", redisAddr)

	StartWorker()
	logger.Sf("✅ Worker jobs démarré")
}
