package utils

import (
	"advent-calendar/internal/config"
	"net/smtp"
)

func SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", config.Config.SMTP_USER, config.Config.SMTP_PASSWORD, config.Config.SMTP_HOST)

	htmlBody := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>` + subject + `</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                background-color: #f5f5f5;
                margin: 0;
                padding: 0;
            }
            .container {
                max-width: 600px;
                margin: 0 auto;
                padding: 20px;
                background-color: #fff;
                border-radius: 5px;
            }
            .header {
                text-align: center;
                padding: 20px;
                background-color: #333;
                color: #fff;
            }
            .content {
                padding: 20px;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <h1>` + subject + `</h1>
            </div>
            <div class="content">
                ` + body + `
            </div>
        </div>
    </body>
    </html>
    `

	message := []byte("Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody)

	addr := config.Config.SMTP_HOST + ":" + config.Config.SMTP_PORT

	err := smtp.SendMail(addr, auth, config.Config.SMTP_USER, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
