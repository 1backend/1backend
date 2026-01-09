package emailservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
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

	// Pre-check: Make sure NO ONE in the list is over the limit
	for _, rawRecipient := range req.To {
		recipient := normalizeEmail(rawRecipient)

		if limiter, exists := s.recipientFilters.Get(recipient); exists {
			if limiter.Tokens() < 1 {
				endpoint.WriteString(w, http.StatusTooManyRequests,
					fmt.Sprintf("Rate limit exceeded for %s", recipient))
				return
			}
		}
	}

	// Execution: Now that we know everyone is likely OK, consume the tokens
	for _, rawRecipient := range req.To {
		recipient := normalizeEmail(rawRecipient)

		limiter, exists := s.recipientFilters.Get(recipient)
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Hour/2), 5)
			s.recipientFilters.Add(recipient, limiter)
		}
		limiter.Allow()
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

// normalizeEmail removes sub-addressing (the + part) for rate limiting purposes.
func normalizeEmail(emailAddr string) string {
	parts := strings.Split(emailAddr, "@")
	if len(parts) != 2 {
		return emailAddr // Not a valid email, return as is
	}

	localPart := parts[0]
	domain := parts[1]

	// Remove everything after '+' in the local part
	if plusIdx := strings.Index(localPart, "+"); plusIdx != -1 {
		localPart = localPart[:plusIdx]
	}

	return strings.ToLower(localPart + "@" + domain)
}

func (s *EmailService) dispatchLocalOrGlobal(
	app openapi.UserSvcApp, req email.SendEmailRequest,
) (*email.Email, error) {
	// This is a pretty way to solve the problem of whether we
	// should use an app specific secret or the default one.
	// Anyway...

	var (
		currentEmailSecrets emailSecrets
		unnamedEmailSecrets emailSecrets
		err                 error
	)

	// Only try to exchange into an app if that's not the default one
	if app.Host != sdk.DefaultAppHost {
		exchangedToken, err := s.options.TokenExchanger.
			ExchangeToken(context.Background(), s.token, endpoint.ExchangeOptions{
				AppId: app.Id,
			})
		if err != nil {
			return nil, errors.Wrap(err, "cannot exchange token")
		}

		t, err := s.emailSecrets(exchangedToken, app, req)
		if err != nil {
			return nil, errors.Wrap(err, "error getting specific email secrets")
		}
		currentEmailSecrets = *t
	}

	if currentEmailSecrets.senderEmail == "" ||
		currentEmailSecrets.senderName == "" || !hasKey(currentEmailSecrets) {
		t, err := s.emailSecrets(s.token, app, req)
		if err != nil {
			return nil, errors.Wrap(err, "error getting unnamed email secrets")
		}
		unnamedEmailSecrets = *t
	}

	finalEmailSecrets := fallbackEmailSecrets(currentEmailSecrets, unnamedEmailSecrets)

	err = s.dispatchEmail(s.token, app, req, finalEmailSecrets)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to send email in apps '%v' and unnamed", app.Host))
	}

	now := time.Now()

	emailId := sdk.Id("email")
	internalId, err := sdk.InternalId(app.Id, emailId)
	if err != nil {
		return nil, err
	}

	em := &email.Email{
		InternalId:  internalId,
		AppId:       app.Id,
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

func hasKey(es emailSecrets) bool {
	if es.mailerSendApiKey != "" {
		return true
	}

	if es.sendgridApiKey != "" {
		return true
	}

	if es.awsAccessKey != "" && es.awsRegion != "" && es.awsSecret != "" {
		return true
	}

	return false
}

type emailSecrets struct {
	senderEmail                        string
	senderName                         string
	sendgridApiKey                     string
	awsAccessKey, awsSecret, awsRegion string
	mailerSendApiKey                   string
}

func fallbackEmailSecrets(primary, backup emailSecrets) emailSecrets {
	return emailSecrets{
		senderEmail:      pick(primary.senderEmail, backup.senderEmail),
		senderName:       pick(primary.senderName, backup.senderName),
		sendgridApiKey:   pick(primary.sendgridApiKey, backup.sendgridApiKey),
		awsAccessKey:     pick(primary.awsAccessKey, backup.awsAccessKey),
		awsSecret:        pick(primary.awsSecret, backup.awsSecret),
		awsRegion:        pick(primary.awsRegion, backup.awsRegion),
		mailerSendApiKey: pick(primary.mailerSendApiKey, backup.mailerSendApiKey),
	}
}

func pick(val, fallback string) string {
	if val != "" {
		return val
	}
	return fallback
}

func (s *EmailService) emailSecrets(
	token string,
	app openapi.UserSvcApp,
	req email.SendEmailRequest,
) (*emailSecrets, error) {
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
		return nil, errors.Wrap(err, "failed to read email secrets")
	}

	if secretResp == nil {
		return nil, errors.New("email secrets are not configured")
	}

	var es emailSecrets

	for _, secret := range secretResp.Secrets {
		if secret.Id == "" || secret.Value == "" {
			logger.Error("Secret key or value is nil", slog.Any("secret", secret))
			continue
		}

		switch secret.Id {
		case "sender-email":
			es.senderEmail = secret.Value
		case "sender-name":
			es.senderName = secret.Value
		case "sendgrid-api-key":
			es.sendgridApiKey = secret.Value
		case "aws-access-key":
			es.awsAccessKey = secret.Value
		case "aws-secret":
			es.awsSecret = secret.Value
		case "aws-region":
			es.awsRegion = secret.Value
		case "mailersend-api-key":
			es.mailerSendApiKey = secret.Value
		}
	}

	return &es, nil
}

func (s *EmailService) dispatchEmail(
	token string,
	app openapi.UserSvcApp,
	req email.SendEmailRequest,
	es emailSecrets,
) error {
	if req.FromEmail != "" {
		es.senderEmail = req.FromEmail
	}

	if es.senderEmail == "" {
		return fmt.Errorf("'sender-email' is not configured")
	}

	if req.FromName != "" {
		es.senderName = req.FromName
	}

	if es.senderName == "" {
		return fmt.Errorf("'sender-name' is not configured")
	}

	if es.awsAccessKey != "" {
		err := s.awsSesSendEmail(
			es.senderEmail, es.senderName,
			es.awsAccessKey, es.awsSecret, es.awsRegion,
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

	if es.mailerSendApiKey != "" {
		err := s.mailerSendSendEmail(
			es.senderEmail, es.senderName,
			es.mailerSendApiKey,
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

	if es.sendgridApiKey != "" {
		err := s.sendgridSendEmail(
			es.senderEmail, es.senderName,
			es.sendgridApiKey, req,
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

	return fmt.Errorf("none of the email providers succeeded")
}
