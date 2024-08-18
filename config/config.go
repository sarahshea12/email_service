package config

import (
	"os"
	"strings"

	entities "github.com/sarahshea12/email_service/entities"
)

func GetProviderDetails(provider string) entities.EmailHost {
	// fetch credentials
	passwordVarName := strings.ToUpper(provider) + "_SMTP_PASSWORD"
	usernameVarName := strings.ToUpper(provider) + "_SMTP_USERNAME"
	smtpPassword := os.Getenv(passwordVarName)
	smtpUsername := os.Getenv(usernameVarName)

	host := ProviderData[provider].(map[string]interface{})["host"].(string)
	port := ProviderData[provider].(map[string]interface{})["port"].(int)

	return entities.EmailHost{
		Host:     host,
		Port:     port,
		Username: smtpUsername,
		Password: smtpPassword,
	}
}
