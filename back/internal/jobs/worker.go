package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nsevenpack/logger/v2/logger"
)

var JobsProcessed chan Job

func StartWorker() {
	go func() {
		for {
			data, err := ClientRedis.RPop(context.Background(), "job:queue").Result()
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			}

			var job Job
			if err := json.Unmarshal([]byte(data), &job); err != nil {
				logger.Ef("Erreur de décodage job : %v", err)
				continue
			}

			if err := routeJob(job); err != nil {
				job.Retry++
				if job.Retry >= job.MaxRetry {
					saveFailedJob(job)
					continue
				}
				retryJob(job)
			}
		}
	}()
}

func routeJob(job Job) error {
	logger.Sf("📌 Job reçu : %s avec payload : %v", job.Name, job.Payload)

	// Envoi du job traité dans le channel pour test (si défini)
	if JobsProcessed != nil {
		select {
		case JobsProcessed <- job:
		default:
			// évite de bloquer si le channel est plein
			logger.Wf("JobsProcessed channel plein, job %s ignoré dans le test", job.Name)
		}
	}

	// Ici tu pourras ajouter la logique métier plus tard
	return nil
}

func retryJob(job Job) {
	data, _ := json.Marshal(job)
	ClientRedis.LPush(context.Background(), "job:queue", data)
}

func saveFailedJob(job Job) {
	key := fmt.Sprintf("job:failed:%s:%d", job.Name, time.Now().Unix())
	data, _ := json.Marshal(job)
	ClientRedis.Set(context.Background(), key, data, 0)
}
