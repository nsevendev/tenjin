package jobs

import (
	"context"
	"fmt"

	"tenjin/back/internal/mailer"
	"tenjin/back/internal/utils/constantes"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleSendMail(job Job, mailerInstance *mailer.Mailer, mu *mailer.MailUploader, jobsProcessed chan Job) error {
	if mailerInstance == nil {
		return fmt.Errorf("MailerInstance non initialisé")
	}

	to := job.Payload["email"]
	subject := job.Payload["subject"]
	body := job.Payload["body"]
	mailType := constantes.MailRegister
	userIDHex := job.Payload["user_id"]
	var userID primitive.ObjectID
	if userIDHex != "" {
		userID, _ = primitive.ObjectIDFromHex(userIDHex)
	}

	logger.Sf("debut du stockage du mail ...")
	_, err := mu.StoreAndCreate(context.Background(), userID, to, subject, body, mailType)
	if err != nil {
		logger.Ef("Erreur stockage mail pour job %s: %v", job.Name, err)
		return err
	}
	logger.Sf("Stockage du mail reussi")

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

	logger.Sf("✅ Mail envoyé et stocké pour job %s à %s", job.Name, to)

	if jobsProcessed != nil {
		select {
		case jobsProcessed <- job:
		default:
			logger.Wf("JobsProcessed channel plein, job %s ignoré dans le test", job.Name)
		}
	}

	return nil
}
