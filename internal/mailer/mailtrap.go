package mailer

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	subject, body, err := buildTemplate(templateFile, data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject)
	message.AddAlternative("text/html", body)

	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)

	var retryErr error
	for i := 0; i < maxRetries; i++ {
		retryErr := dialer.DialAndSend(message)
		if retryErr != nil {
			// exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		log.Printf("Email sent with status code %v", 200)
		return 200, nil
	}

	return -1, fmt.Errorf("failed to send email after %d attempts, error: %v", maxRetries, retryErr)
}
