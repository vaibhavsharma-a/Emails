package notifications

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailCreds struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SendEmail(emailCredentials EmailCreds) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("could not load the env: %w", err)
	}

	email := os.Getenv("EMAIL_ADDR")
	pass := os.Getenv("EMAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", emailCredentials.From)
	m.SetHeader("To", emailCredentials.To)
	m.SetHeader("Subject", emailCredentials.Subject)
	m.SetBody("text/plain", emailCredentials.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)
	err := d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("fail to send email: %w", err)

	}
	return nil
}
