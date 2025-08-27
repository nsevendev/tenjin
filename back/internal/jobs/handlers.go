package jobs

import (
	"fmt"

	"tenjin/back/internal/mailer"

	"github.com/nsevenpack/logger/v2/logger"
)

func HandleSendMail(job Job, mailerInstance *mailer.Mailer, jobsProcessed chan Job) error {
	if mailerInstance == nil {
		return fmt.Errorf("MailerInstance non initialisé")
	}

	to := job.Payload["email"]
	subject := job.Payload["subject"]
	body := job.Payload["body"]

	if subject == "" {
		subject = fmt.Sprintf("Notification %s", job.Name)
	}
	if body == "" {
		body = fmt.Sprintf("Vous avez un nouveau message pour le job %s", job.Name)
	}

	m := mailer.Mail{
		To:      to,
		Subject: subject,
		Body:    body,
		Type:    job.Name,
		Context: job.Payload,
	}

	if err := mailerInstance.Send(m); err != nil {
		logger.Ef("Erreur envoi mail pour job %s: %v", job.Name, err)
		return err
	}

	logger.Sf("✅ Mail envoyé pour job %s à %s", job.Name, to)

	if jobsProcessed != nil {
		select {
		case jobsProcessed <- job:
		default:
			logger.Wf("JobsProcessed channel plein, job %s ignoré dans le test", job.Name)
		}
	}

	return nil
}
