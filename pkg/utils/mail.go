package utils

import (
	"advent-calendar/internal/config"
	"log"
	"net/smtp"
)

func SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", config.Config.SMTP_USER, config.Config.SMTP_PASSWORD, config.Config.SMTP_HOST)

	message := []byte("Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		body)

	addr := config.Config.SMTP_HOST + ":" + config.Config.SMTP_PORT

	err := smtp.SendMail(addr, auth, config.Config.SMTP_USER, []string{to}, message)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
