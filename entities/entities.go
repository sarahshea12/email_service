package entities

import (
	"fmt"
	"net/smtp"
)

type EmailHost struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Email struct {
	From         string
	To           string
	SMTPUsername string
	SMTPPassword string
	Host         string
	Port         int
	Body         string
} // TODO: make a field EmailHost instead of duplicating fields

func (s *Email) SendEmail() error {
	auth := smtp.PlainAuth("", s.SMTPUsername, s.SMTPPassword, s.Host)
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)

	err := smtp.SendMail(address, auth, s.From, []string{s.To}, []byte(s.Body))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}
