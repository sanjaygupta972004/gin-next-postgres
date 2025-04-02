package email

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/savvy-bit/gin-react-postgres/config"
)

func SendEmail(sssClient *ses.Client, recipient, subject, body string) error {
	awsConfig := config.GetGlobalConfig().AWS

	sender := awsConfig.SesSenderEmail
	if sender == "" {
		return errors.New("SES sender email is not configured")
	}

	input := &ses.SendEmailInput{
		Source: &sender,
		Destination: &types.Destination{
			ToAddresses: []string{
				recipient,
			},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: &subject,
			},
			Body: &types.Body{
				Html: &types.Content{
					Data: &body,
				},
			},
		},
	}

	_, err := sssClient.SendEmail(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("Email sent successfully to %s\n", recipient)
	return nil
}
