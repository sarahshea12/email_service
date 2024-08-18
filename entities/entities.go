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
	From     string
	To       string
	Body     string
	HostData EmailHost
}

func (s *Email) SendEmail() error {
	auth := smtp.PlainAuth(
		"",
		s.HostData.Username,
		s.HostData.Password,
		s.HostData.Host,
	)
	address := fmt.Sprintf("%s:%d", s.HostData.Host, s.HostData.Port)

	err := smtp.SendMail(address, auth, s.From, []string{s.To}, []byte(s.Body))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}
