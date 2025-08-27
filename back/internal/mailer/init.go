package mailer

import (
	"os"
)

var MailService *Mailer

func InitMailer() {
	MailService = NewMailer(
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_PORT"),
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASS"),
		os.Getenv("MAIL_FROM"),
	)
}
