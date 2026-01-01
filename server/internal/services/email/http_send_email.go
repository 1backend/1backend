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
	if !isAuthRsp.Authorized {
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

	em, err := s.dispatchLocalOrGlobal(isAuthRsp.App, *req)
	if err != nil {
		logger.Error(
			"Failed to send email",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	response := email.SendEmailResponse{
		EmailId: em.Id,
		Status:  "sent",
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func (s *EmailService) dispatchLocalOrGlobal(
	app openapi.UserSvcApp, req email.SendEmailRequest,
) (*email.Email, error) {
	// This is a pretty way to solve the problem of whether we
	// should use an app specific secret or the default one.
	// Anyway...

	exchangedToken, err := s.options.TokenExchanger.ExchangeToken(context.Background(), s.token, endpoint.ExchangeOptions{
		AppHost: app.Id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot exchange token")
	}

	err = s.dispatchEmail(exchangedToken, app, req)
	if err != nil {
		logger.Warn(
			"Failed to send email in specific app",
			slog.String("appHost", app.Host),
			slog.String("appId", app.Host),
			slog.Any("error", err),
		)
	}

	err = s.dispatchEmail(s.token, app, req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to send email in apps '%v' and unnamed", app.Host))
	}

	now := time.Now()

	emailId := sdk.Id("email")
	em := &email.Email{
		Id:          emailId,
		To:          req.To,
		CC:          req.CC,
		BCC:         req.BCC,
		Subject:     req.Subject,
		Body:        req.Body,
		ContentType: req.ContentType,
		CreatedAt:   now,
		FromName:    req.FromName,
	}
	err = s.emailStore.Create(em)
	if err != nil {
		return nil, errors.Wrap(err, "failed to store email")
	}

	return em, nil
}

func (s *EmailService) dispatchEmail(
	token string,
	app openapi.UserSvcApp,
	req email.SendEmailRequest,
) error {
	secretClient := s.options.ClientFactory.Client(client.WithToken(token)).SecretSvcAPI
	secretResp, _, err := secretClient.ListSecrets(context.Background()).Body(
		openapi.SecretSvcListSecretsRequest{
			Ids: []string{
				"sender-email",
				"sender-name",
				"sendgrid-api-key",                           // SendGrid
				"aws-access-key", "aws-secret", "aws-region", // SES
				"mailersend-api-key", // MailerSend
			},
		},
	).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to read email secrets")
	}

	if secretResp == nil {
		return errors.New("email secrets are not configured")
	}

	var (
		senderEmail                        string
		senderName                         string
		sendgridApiKey                     string
		awsAccessKey, awsSecret, awsRegion string
		mailerSendApiKey                   string
	)

	for _, secret := range secretResp.Secrets {
		if secret.Id == "" || secret.Value == "" {
			logger.Error("Secret key or value is nil", slog.Any("secret", secret))
			continue
		}

		switch secret.Id {
		case "sender-email":
			senderEmail = secret.Value
		case "sender-name":
			senderName = secret.Value
		case "sendgrid-api-key":
			sendgridApiKey = secret.Value
		case "aws-access-key":
			awsAccessKey = secret.Value
		case "aws-secret":
			awsSecret = secret.Value
		case "aws-region":
			awsRegion = secret.Value
		case "mailersend-api-key":
			mailerSendApiKey = secret.Value
		}
	}

	if req.FromEmail != "" {
		senderEmail = req.FromEmail
	}

	if senderEmail == "" {
		return fmt.Errorf("'sender-email' is not configured")
	}

	if req.FromName != "" {
		senderName = req.FromName
	}

	if senderName == "" {
		return fmt.Errorf("'sender-name' is not configured")
	}

	if awsAccessKey != "" {
		err = s.awsSesSendEmail(
			senderEmail, senderName,
			awsAccessKey, awsSecret, awsRegion,
			req,
		)
		if err != nil {
			logger.Warn("AWS email sending failed",
				slog.Any("error", err),
				slog.String("appId", app.Id),
				slog.String("appHost", app.Host),
			)
		} else {
			return nil
		}
	}

	if mailerSendApiKey != "" {
		err = s.mailerSendSendEmail(
			senderEmail, senderName,
			mailerSendApiKey,
			req,
		)
		if err != nil {
			logger.Warn("MailerSend email sending failed",
				slog.Any("error", err),
				slog.String("appId", app.Id),
				slog.String("appHost", app.Host),
			)
		} else {
			return nil
		}
	}

	if sendgridApiKey != "" {
		err = s.sendgridSendEmail(
			senderEmail, senderName,
			sendgridApiKey, req,
		)
		if err != nil {
			logger.Warn("Sendgrid email sending failed",
				slog.Any("error", err),
				slog.String("appId", app.Id),
				slog.String("appHost", app.Host),
			)
		} else {
			return nil
		}
	}

	if err == nil {
		return fmt.Errorf("no email provider is configured")
	}

	return errors.Wrap(err, "none of the email providers succeeded")
}
