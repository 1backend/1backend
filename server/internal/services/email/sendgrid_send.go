package emailservice

import (
	"fmt"

	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (s *EmailService) sendgridSendEmail(
	senderEmail, senderName,
	sendgridApiKey string,
	req email.SendEmailRequest,
) error {
	from := mail.NewEmail(senderName, senderEmail)

	subject := req.Subject

	// @todo For simplicity, only sending to the first recipient
	to := mail.NewEmail("Recipient", req.To[0])

	var content mail.Content
	if req.ContentType == "text/html" {
		content = *mail.NewContent("text/html", req.Body)
	} else {
		content = *mail.NewContent("text/plain", req.Body)
	}

	message := mail.NewV3MailInit(from, subject, to, &content)

	for _, attachment := range req.Attachments {
		if attachment.FileId != "" && attachment.Content != "" {
			return errors.New("only one of 'FileId' or 'Content' should be provided")
		}

		// @todo support FileSvc.serveUpload too

		file := mail.NewAttachment()
		file.SetContent(attachment.Content)   // File content as base64-encoded string
		file.SetType(attachment.ContentType)  // MIME type of the file
		file.SetFilename(attachment.Filename) // Filename for the attachment
		message.AddAttachment(file)
	}

	client := sendgrid.NewSendClient(sendgridApiKey)

	resp, err := client.Send(message)
	if err != nil {
		return errors.Wrap(err, "failed to send email with Sendgrid")
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error sending email with Sendgrid: %s", resp.Body)
	}

	return nil
}
