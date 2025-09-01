package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"tenjin/back/internal/mailer"
	"time"

	"github.com/nsevenpack/logger/v2/logger"
)

func StartWorker(mailerInstance *mailer.Mailer, mu *mailer.MailUploader, jobsProcessed chan Job) {
	logger.If("👂 Worker en écoute sur Redis...")
	go func() {
		for {
			data, err := ClientRedis.RPop(context.Background(), "job:queue").Result()
			logger.Sf("🔹 Lecture job Redis: data=%v, err=%v", data, err)
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			}

			var job Job
			if err := json.Unmarshal([]byte(data), &job); err != nil {
				logger.Ef("Erreur de décodage job : %v", err)
				continue
			}

			logger.If("👀 Worker a reçu un job : %+v", job)
			if err := routeJob(job, mailerInstance, mu, jobsProcessed); err != nil {
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

func routeJob(job Job, mailerInstance *mailer.Mailer, mu *mailer.MailUploader, jobsProcessed chan Job) error {
	switch job.Name {
	case "mail:send":
		return HandleSendMail(job, mailerInstance, mu, jobsProcessed)
	default:
		logger.Wf("Job inconnu : %s", job.Name)
		return nil
	}
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
