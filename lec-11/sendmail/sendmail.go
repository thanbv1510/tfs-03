package sendmail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	apiKey           = "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"
	fromName         = "Than Bui Van"
	fromAddress      = "thanbv1510@gmail.com"
	subject          = "Thank you for your order"
	plainTextContent = "Thank you so much <3"
	htmlContent      = `<p>Hello,</p>
					<p>thank you for shopping<p>
					<p>thanks again,</p>
					<strong>Than Bui Van</strong>`
)

func SendEmailThankYou(toUsername, toEmail string) (int, error) {
	from := mail.NewEmail(fromName, fromAddress)
	to := mail.NewEmail(toUsername, toEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return -1, err

	}

	return response.StatusCode, nil
}
