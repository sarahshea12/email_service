package services

import (
	config "github.com/sarahshea12/email_service/config"
	entities "github.com/sarahshea12/email_service/entities"
)

func SendPrimaryEmail(recipient string, body string) error {
	hostData := config.GetProviderDetails("primary")
	sender := config.Sender
	email := entities.Email{
		From:         sender,
		To:           recipient,
		SMTPUsername: hostData.Username,
		SMTPPassword: hostData.Password,
		Host:         hostData.Host,
		Port:         hostData.Port,
		Body:         body,
	}

	return email.SendEmail()
}

func SendSecondaryEmail(recipient string, body string) error {
	hostData := config.GetProviderDetails("secondary")
	sender := config.Sender
	email := entities.Email{
		From:         sender,
		To:           recipient,
		SMTPUsername: hostData.Username,
		SMTPPassword: hostData.Password,
		Host:         hostData.Host,
		Port:         hostData.Port,
		Body:         body,
	}

	return email.SendEmail()
} // TODO: add helper to reduce duplication
