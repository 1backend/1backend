/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	htmlTemplate "html/template"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	textTemplate "text/template"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

const (
	otpCodeLength      = 6
	otpExpirationDelta = 10 * time.Minute
)

// @ID sendOtp
// @Summary Send OTP
// @Description Generates and sends a one-time password (OTP) to the specified contact.
// @Description
// @Description The OTP can be used for contact verification or login depending on purpose.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SendOtpRequest true "Send OTP Request"
// @Success 200 {object} user.SendOtpResponse "OTP sent successfully"
// @Param Accept-Language header string false "Language preference for the email"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 404 {object} user.ErrorResponse "Contact not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/otp/send [post]
func (s *UserService) SendOTP(w http.ResponseWriter, r *http.Request) {
	req := user.SendOtpRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode SendOTP request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.ContactId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "contactId is required")
		return
	}
	if req.ContactPlatform == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "contactPlatform is required")
		return
	}

	if isEmail(req.ContactId) {
		req.ContactId = normalizeEmail(req.ContactId)
	}

	code, err := generateOtpCode(otpCodeLength)
	if err != nil {
		logger.Error("Failed to generate OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	hash, err := s.hashPassword(code)
	if err != nil {
		logger.Error("Failed to hash OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	appId, err := s.options.TokenExchanger.AppIdFromHost(
		r.Context(),
		req.AppHost,
	)
	if err != nil {
		logger.Error(
			"Failed to get app id from host",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}
	logger.Info("Got app id",
		slog.String("appId", appId),
		slog.String("host", req.AppHost),
	)

	now := time.Now()

	otpId := sdk.Id("otp")
	internalId, err := sdk.InternalId(appId, otpId)
	if err != nil {
		logger.Error(
			"Failed to create internal id",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	otp := &user.OTP{
		InternalId:      internalId,
		AppId:           appId,
		Id:              otpId,
		CreatedAt:       now,
		UpdatedAt:       now,
		ContactId:       req.ContactId,
		ContactPlatform: req.ContactPlatform,
		ExpiresAt:       now.Add(otpExpirationDelta),
		CodeHash:        hash,
	}

	if err := s.otpStore.Upsert(otp); err != nil {
		logger.Error("Failed to store OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	emailReq, err := s.dispatchOtp(
		r.Context(),
		appId,
		req.AppHost,
		getPreferredLanguage(r),
		otp,
		code,
	)

	if err != nil {
		logger.Error("Failed to send OTP email", slog.Any("error", err))
		if derr := s.otpStore.Query(
			datastore.Equals(datastore.Field("id"), otp.Id),
		).Delete(); derr != nil {
			logger.Error("Failed to delete OTP after email error", slog.Any("error", derr))
		}
		endpoint.InternalServerError(w)
		return
	}

	response := user.SendOtpResponse{
		OtpId: otp.Id,
	}
	if s.options.Test {
		response.Code = code
		response.Subject = emailReq.Subject
		response.Body = emailReq.Body
		if emailReq.ContentType != nil {
			response.ContentType = *emailReq.ContentType
		}
		if emailReq.FromName != nil {
			response.FromName = *emailReq.FromName
		}
		if emailReq.FromEmail != nil {
			response.FromEmail = *emailReq.FromEmail
		}
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func getPreferredLanguage(r *http.Request) string {
	acceptLang := r.Header.Get("Accept-Language")
	if acceptLang == "" {
		return "en"
	}

	// Accept-Language can be: "en-GB,en;q=0.9,fr;q=0.8"
	// We take the first part of the first preference
	firstPart := strings.Split(acceptLang, ",")[0]   // "en-GB"
	languageOnly := strings.Split(firstPart, "-")[0] // "en"

	return strings.ToLower(languageOnly)
}

func (s *UserService) dispatchOtp(
	ctx context.Context,
	appId string,
	appHost string,
	language string,
	otp *user.OTP,
	code string,
) (*openapi.EmailSvcSendEmailRequest, error) {
	if otp.ContactPlatform == "" {
		return nil, fmt.Errorf("contact platform is empty")
	}

	// @todo only email is supported now
	if otp.ContactPlatform != "email" {
		return nil, fmt.Errorf("unsupported contact platform '%s'", otp.ContactPlatform)
	}

	var (
		err            error
		exchangedToken string
	)

	if appHost != sdk.DefaultAppHost {
		exchangedToken, err = s.options.
			TokenExchanger.
			ExchangeToken(ctx, s.token, endpoint.ExchangeOptions{AppId: appId})
		if err != nil {
			return nil, err
		}
	}

	ts, err := s.getOtpTemplateSecrets(ctx, appHost, exchangedToken, language)
	if err != nil {
		return nil, err
	}

	// 2. Prepare Template Data
	expiresIn := time.Until(otp.ExpiresAt).Round(time.Second)
	if expiresIn < 0 {
		expiresIn = 0
	}

	data := map[string]interface{}{
		"Code":   code,
		"Expiry": expiresIn.String(),
	}

	// 3. Render Subject (using text/template)
	finalSubject := "Your one-time password" // Default
	if ts.Subject != "" {
		if renderedSub, err := renderTemplate("otp-sub", ts.Subject, data, false); err == nil {
			finalSubject = renderedSub
		} else {
			return nil, err
		}
	}

	// 4. Render Body (using html/template for safety if HTML, or text if plain)
	finalBody := fmt.Sprintf("Your verification code is %s. It expires in %s.", code, expiresIn)
	if ts.Body != "" {
		isHTML := ts.ContentType == "text/html"
		if renderedBody, err := renderTemplate("otp-body", ts.Body, data, isHTML); err == nil {
			finalBody = renderedBody
		} else {
			return nil, err
		}
	}

	// 5. Send Email
	emailReq := openapi.EmailSvcSendEmailRequest{
		To:          []string{otp.ContactId},
		Subject:     finalSubject,
		Body:        finalBody,
		ContentType: openapi.PtrString(ts.ContentType),
	}

	if ts.SenderName != "" {
		emailReq.FromName = &ts.SenderName
	}

	if ts.SenderEmail != "" {
		emailReq.FromEmail = &ts.SenderEmail
	}

	contentType := "text/plain"
	if ts.ContentType != "" {
		contentType = ts.ContentType
	}

	emailReq.ContentType = openapi.PtrString(contentType)

	token := s.token
	if exchangedToken != "" {
		token = exchangedToken
	}

	if !s.options.Test {
		_, _, err := s.options.ClientFactory.
			Client(client.WithToken(token)).
			EmailSvcAPI.
			SendEmail(ctx).
			Body(emailReq).
			Execute()
		if err != nil {
			return nil, errors.Wrap(err, "email svc send email")
		}
	}

	return &emailReq, nil
}

func renderTemplate(name, text string, data any, isHTML bool) (string, error) {
	var buf bytes.Buffer
	var err error

	if isHTML {
		tmpl, parseErr := htmlTemplate.New(name).Parse(text)
		if parseErr != nil {
			return "", parseErr
		}
		err = tmpl.Execute(&buf, data)
	} else {
		tmpl, parseErr := textTemplate.New(name).Parse(text)
		if parseErr != nil {
			return "", parseErr
		}
		err = tmpl.Execute(&buf, data)
	}

	return buf.String(), err
}

type otpTemplateSecrets struct {
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	ContentType string
}

func (s *UserService) getOtpTemplateSecrets(
	ctx context.Context,
	appHost string,
	exchangedToken, lang string,
) (*otpTemplateSecrets, error) {
	// Generate keys for both the requested language and English fallback
	keys := []string{
		"sender-name",
		"sender-email",
		"otp-email-subject-" + lang, "otp-email-subject",
		"otp-email-body-" + lang, "otp-email-body",
		"otp-email-type",
	}

	var err error
	primary := map[string]string{}

	if exchangedToken != "" {
		primary, err = s.fetchSecretMap(ctx, exchangedToken, keys)
		if err != nil {
			return nil, err
		}
	}

	backup, err := s.fetchSecretMap(ctx, s.token, keys)
	if err != nil {
		return nil, err
	}

	logger.Info("Sending OTP with",
		slog.Bool("exchanged", exchangedToken != ""),
		slog.Any("appHost", appHost),
	)

	// Resolve with priority: App(Lang) -> App(en) -> Global(Lang) -> Global(en)
	return &otpTemplateSecrets{
		SenderName: pick(
			primary["sender-name"],
			backup["sender-name"],
		),
		SenderEmail: pick(
			primary["sender-email"],
			backup["sender-email"],
		),
		Subject: pick(
			primary["otp-email-subject-"+lang],
			primary["otp-email-subject"],
			backup["otp-email-subject-"+lang],
			backup["otp-email-subject"],
		),
		Body: pick(
			primary["otp-email-body-"+lang],
			primary["otp-email-body"],
			backup["otp-email-body-"+lang],
			backup["otp-email-body"],
		),
		ContentType: pick(primary["otp-email-type"], backup["otp-email-type"], "text/plain"),
	}, nil
}

func (s *UserService) fetchSecretMap(
	ctx context.Context,
	token string,
	keys []string,
) (map[string]string, error) {
	res := make(map[string]string)

	resp, _, err := s.options.ClientFactory.Client(client.WithToken(token)).SecretSvcAPI.
		ListSecrets(ctx).Body(openapi.SecretSvcListSecretsRequest{Ids: keys}).Execute()
	if err != nil {
		return nil, err
	}

	for _, sec := range resp.Secrets {
		res[sec.Id] = sec.Value
	}

	return res, nil
}

func pick(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

func generateOtpCode(length int) (string, error) {
	if length <= 0 {
		length = 6
	}

	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code[i] = byte('0') + byte(n.Int64())
	}

	return string(code), nil
}
