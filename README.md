# Email service

## Description
Basic email service that accepts a message and recipient to send an email to.
First the email will attempt to be sent using mailtrap but if it fails will be sent
using Amazon SES.

mailtrap: https://mailtrap.io/
</br>
Amazon SES: https://docs.aws.amazon.com/ses/

## Usage
```bash
go run main.go
```
