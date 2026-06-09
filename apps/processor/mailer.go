package main

import (
	"fmt"
	"log"
	"net/smtp"
)

type Mailer struct {
	config Config
}

func newMailer(config Config) *Mailer {
	return &Mailer{config}
}

func (m *Mailer) send(message AppointmentMessage) bool {
	serverURL := fmt.Sprintf("%s:%d", m.config.SMTPHost, m.config.SMTPPort)
	auth := smtp.PlainAuth("", m.config.SMTPUsername, m.config.SMTPPassword, m.config.SMTPHost)

	recipients := []string{message.ContactEmail}
	body := fmt.Sprintf("Olá %s, seu agendamento #%d foi confirmado.",
		message.ContactName,
		message.AppointmentId,
	)
	rawMessage := []byte(
		"From: noreply@test.com\r\n" +
			"To: " + message.ContactEmail + "\r\n" +
			"Subject: Confirmação de Appointment #" + fmt.Sprintf("%d", message.AppointmentId) + "\r\n" +
			"\r\n" +
			body,
	)

	err := smtp.SendMail(serverURL, auth, "noreply@test.com", recipients, rawMessage)

	if err != nil {
		log.Printf("Error on send email", err)
		return false
	}

	log.Printf("Email sent to %s", message.ContactEmail)
	return true
}
