package emailservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// @ID sendEmail
// @Summary Send an Email
// @Description Sends an email with optional attachments via a supported email provider.
// @Description
// @Description Currently, only SendGrid is supported. Additional providers may be added in the future.
// @Description
// @Description Required secrets from the Secret Svc for SendGrid:
// @Description - `sender-email`: Sender's email address.
// @Description - `sender-name`: Sender's display name.
// @Description - `sendgrid-api-key`: API key for SendGrid.
// @Tags Email Svc
// @Accept json
// @Produce json
// @Param body body email.SendEmailRequest true "Send Email Request"
// @Success 200 {object} email.SendEmailResponse "Successfully sent the email"
// @Failure 400 {object} email.ErrorResponse "Invalid JSON"
// @Failure 401 {object} email.ErrorResponse "Unauthorized"
// @Failure 500 {object} email.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /email-svc/email [post]
func (s *EmailService) SendEmail(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := s.options.PermissionChecker.HasPermission(
		r,
		email.PermissionSendEmail,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &email.SendEmailRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error("Failed to decode request body", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if len(req.To) == 0 {
		logger.Error("No recipients specified", slog.Any("request", req))
		endpoint.WriteString(w, http.StatusBadRequest, "At least one recipient is required")
		return
	}

	err = s.sendgridSendEmail(*req)
	if err != nil {
		logger.Error(
			"Failed to send email",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	response := email.SendEmailResponse{
		Status: "sent",
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func (s *EmailService) sendgridSendEmail(req email.SendEmailRequest) error {
	secretClient := s.options.ClientFactory.Client(client.WithToken(s.token)).SecretSvcAPI
	secretResp, _, err := secretClient.ListSecrets(context.Background()).Body(
		openapi.SecretSvcListSecretsRequest{
			Keys: []string{
				"sender-email",
				"sender-name",
				"sendgrid-api-key",
			},
		},
	).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to read SendGrid API key")
	}

	if secretResp == nil {
		return errors.New("SendGrid is not configured")
	}

	var (
		senderEmail    string
		senderName     string
		sendgridApiKey string
	)

	for _, secret := range secretResp.Secrets {
		if secret.Key == nil || secret.Value == nil {
			logger.Error("Secret key or value is nil", slog.Any("secret", secret))
			continue
		}

		switch *secret.Key {
		case "sender-email":
			senderEmail = *secret.Value
		case "sender-name":
			senderName = *secret.Value
		case "sendgrid-api-key":
			sendgridApiKey = *secret.Value
		}
	}

	if senderEmail == "" {
		return errors.New("sender email is not configured")
	}
	if senderName == "" {
		return errors.New("sender name is not configured")
	}
	if sendgridApiKey == "" {
		return errors.New("SendGrid API key is not configured")
	}

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
		return fmt.Errorf("failed to send email: %v", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error sending email: %s", resp.Body)
	}

	now := time.Now()

	emailId := sdk.Id("email")
	err = s.emailStore.Create(&email.Email{
		Id:          emailId,
		To:          req.To,
		CC:          req.CC,
		BCC:         req.BCC,
		Subject:     req.Subject,
		Body:        req.Body,
		ContentType: req.ContentType,
		CreatedAt:   now,
	})
	if err != nil {
		return errors.Wrap(err, "failed to store email")
	}

	return nil
}
