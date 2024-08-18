package entities

import (
	"fmt"
	"net/smtp"
)

type SMTPClient interface {
	authenticate(hostData EmailHost) (smtp.Auth, string)
	sendMail(addr string, auth smtp.Auth, from string, to []string, msg []byte) error
}

type SMTPClientSender struct{}

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
	Client   SMTPClient
}

func (s *SMTPClientSender) authenticate(hostData EmailHost) (smtp.Auth, string) {
	auth := smtp.PlainAuth(
		"",
		hostData.Username,
		hostData.Password,
		hostData.Host,
	)
	address := fmt.Sprintf("%s:%d", hostData.Host, hostData.Port)
	return auth, address
}

func (s *SMTPClientSender) sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

func (s *Email) SendEmail() error {
	auth, address := s.Client.authenticate(s.HostData)

	err := s.Client.sendMail(address, auth, s.From, []string{s.To}, []byte(s.Body))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}
