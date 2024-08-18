# Email service

## Description
Basic email service that accepts a message and recipient to send an email to.
First the email will attempt to be sent using mailtrap but if it fails will be sent
using Amazon SES.

mailtrap: https://mailtrap.io/
</br>
Amazon SES: https://docs.aws.amazon.com/ses/

## Usage
With docker installed:
```bash
docker build -t go-email-service .
docker run -it go-email-service
```
