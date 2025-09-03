package mailer

import (
	"fmt"
	"net/smtp"
)

func NewMailer(host, port, user, pass, from string) *Mailer {
	return &Mailer{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		From:     from,
	}
}

func (m *Mailer) Send(mail Mail) error {
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	to := []string{mail.To}
	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
			"%s",
		m.From, mail.To, mail.Subject, mail.Body,
	))

	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	return smtp.SendMail(addr, auth, m.From, to, msg)
}

/*func (m *Mailer) Send(mail Mail) error {
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	to := []string{mail.To}
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		m.From, mail.To, mail.Subject, mail.Body,
	))

	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	return smtp.SendMail(addr, auth, m.From, to, msg)
}*/
