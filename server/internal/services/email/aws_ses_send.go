package emailservice

import (
	"context"
	"fmt"

	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	sesTypes "github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/pkg/errors"
)

func (s *EmailService) awsSesSendEmail(
	senderEmail, senderName, accessKey, secretKey, region string,
	req email.SendEmailRequest,
) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return errors.Wrap(err, "unable to load AWS config")
	}

	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(fmt.Sprintf("%s <%s>", senderName, senderEmail)),
		Destination: &sesTypes.Destination{
			ToAddresses: req.To,
		},
		Content: &sesTypes.EmailContent{
			Simple: &sesTypes.Message{
				Subject: &sesTypes.Content{Data: aws.String(req.Subject)},
				Body:    &sesTypes.Body{},
			},
		},
	}

	if req.ContentType == "text/html" {
		input.Content.Simple.Body.Html = &sesTypes.Content{Data: aws.String(req.Body)}
	} else {
		input.Content.Simple.Body.Text = &sesTypes.Content{Data: aws.String(req.Body)}
	}

	_, err = client.SendEmail(context.TODO(), input)
	return err
}
