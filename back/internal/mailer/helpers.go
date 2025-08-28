package mailer

import (
	"context"
	"fmt"
	"tenjin/back/internal/mail"
	"tenjin/back/internal/utils/constantes"
	"tenjin/back/internal/utils/filestores"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailUploader struct {
	fileStore *filestores.FileStoreService
	mailSvc   *mail.MailService
}

func (mu *MailUploader) StoreAndCreate(ctx context.Context, userID primitive.ObjectID, to, subject, body string, mailType constantes.TypeMail) (*mail.Mail, error) {
	filename := fmt.Sprintf("%s.html", filestores.RandHex(16))

	uploadRes, err := mu.fileStore.UploadBytes(ctx, "mails", filename, []byte(body))
	if err != nil {
		return nil, fmt.Errorf("erreur upload sur R2 : %w", err)
	}

	dto := mail.MailCreateDto{
		UserID:  userID,
		To:      to,
		Subject: subject,
		Body:    body,
		Type:    mailType,
		S3Path:  &uploadRes.StoredPath,
		MetaName: &filename,
	}

	return mu.mailSvc.Create(ctx, dto)
}
