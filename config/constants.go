package config

var ProviderData = map[string]interface{}{
	// Mailtrap
	"primary": map[string]interface{}{
		"host": "sandbox.smtp.mailtrap.io",
		"port": 2525,
	},

	// Amazon
	"secondary": map[string]interface{}{
		"host": "email-smtp.us-west-2.amazonaws.com",
		"port": 587,
	},
}

const (
	Sender string = "from@example.com"
)
