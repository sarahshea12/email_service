package services

import (
	config "github.com/sarahshea12/email_service/config"
	entities "github.com/sarahshea12/email_service/entities"
)

var SendPrimaryEmail = func(recipient string, body string) error {
	hostData := config.GetProviderDetails("primary")
	email := entities.Email{
		From:     config.Sender,
		To:       recipient,
		Body:     body,
		HostData: hostData,
		Client:   &entities.SMTPClientSender{},
	}

	return email.SendEmail()
}

var SendSecondaryEmail = func(recipient string, body string) error {
	hostData := config.GetProviderDetails("secondary")
	email := entities.Email{
		From:     config.Sender,
		To:       recipient,
		Body:     body,
		HostData: hostData,
		Client:   &entities.SMTPClientSender{},
	}

	return email.SendEmail()
}
