package main

import (
	"fmt"

	services "github.com/sarahshea12/email_service/services"
)

func sendCascadingEmail(recipient string, message string) error {
	primaryErr := services.SendPrimaryEmail(recipient, message)
	if primaryErr != nil {
		secondaryErr := services.SendSecondaryEmail(recipient, message)
		if secondaryErr != nil {
			fmt.Println("Failed to send emails")
			return secondaryErr
		}
	}
	return nil
}

func main() {
	var message string
	var recipient string

	fmt.Print("Enter a message: ")
	_, err := fmt.Scanln(&message)
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}

	fmt.Print("Enter a recipient: ")
	_, err = fmt.Scanln(&recipient)
	if err != nil {
		fmt.Println("Error reading recipient:", err)
		return
	}

	sendCascadingEmail(recipient, message)
}
