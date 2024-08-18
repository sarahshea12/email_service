package main

import (
	"errors"
	"net/smtp"
	"testing"

	entities "github.com/sarahshea12/email_service/entities"
	"github.com/sarahshea12/email_service/services"
	"github.com/stretchr/testify/assert"
)

// mock implementation of SMTPClient
type MockSMTPClient struct {
	Auth    smtp.Auth
	Address string
	SendErr error
}

func (m *MockSMTPClient) Authenticate(hostData entities.EmailHost) (smtp.Auth, string) {
	return m.Auth, m.Address
}

func (m *MockSMTPClient) SendMail(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	return m.SendErr
}

func TestSendEmail_Success(t *testing.T) {
	mockClient := &MockSMTPClient{
		Auth:    nil,
		Address: "smtp.example.com:587",
		SendErr: nil,
	}

	email := &entities.Email{
		From: "from@example.com",
		To:   "to@example.com",
		Body: "Test message",
		HostData: entities.EmailHost{
			Username: "user",
			Password: "password",
			Host:     "smtp.example.com",
			Port:     587,
		},
		Client: mockClient,
	}

	err := email.SendEmail()
	assert.NoError(t, err)
}

func TestSendCascadingEmail_Success(t *testing.T) {
	services.SendPrimaryEmail = func(recipient string, message string) error {
		return nil
	}
	services.SendSecondaryEmail = func(recipient string, message string) error {
		return errors.New("secondary service should not be called")
	}

	err := sendCascadingEmail("to@example.com", "Test message")
	assert.NoError(t, err)
}

func TestSendCascadingEmail_PrimaryFails_SecondarySucceeds(t *testing.T) {
	services.SendPrimaryEmail = func(recipient string, message string) error {
		return errors.New("primary service down")
	}
	services.SendSecondaryEmail = func(recipient string, message string) error {
		return nil
	}

	err := sendCascadingEmail("to@example.com", "Test message")
	assert.NoError(t, err)
}

func TestSendCascadingEmail_BothFail(t *testing.T) {
	services.SendPrimaryEmail = func(recipient string, message string) error {
		return errors.New("primary service down")
	}
	services.SendSecondaryEmail = func(recipient string, message string) error {
		return errors.New("secondary service down")
	}

	err := sendCascadingEmail("to@example.com", "Test message")
	assert.Error(t, err)
}
