package emailservice

import (
	"context"
	"time"

	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/mailersend/mailersend-go"
	"github.com/pkg/errors"
)

func (s *EmailService) mailerSendSendEmail(
	senderEmail, senderName, mailersendApiKey string,
	req email.SendEmailRequest,
) error {
	ms := mailersend.NewMailersend(mailersendApiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	subject := req.Subject
	text := ""
	html := ""

	if req.ContentType == "text/html" {
		html = req.Body
	} else {
		text = req.Body
	}

	from := mailersend.From{
		Name:  senderName,
		Email: senderEmail,
	}

	var recipients []mailersend.Recipient
	for _, toEmail := range req.To {
		recipients = append(recipients, mailersend.Recipient{
			Email: toEmail,
		})
	}

	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	// Handle Attachments
	for _, att := range req.Attachments {
		message.Attachments = []mailersend.Attachment{
			{
				Content:     att.Content, // Base64 string
				Filename:    att.Filename,
				Disposition: "attachment",
			},
		}
	}

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		return errors.Wrap(err, "mailersend api error")
	}

	if res.StatusCode >= 400 {
		return errors.Errorf("mailersend returned status: %d", res.StatusCode)
	}

	return nil
}
