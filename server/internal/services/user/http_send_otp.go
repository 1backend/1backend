/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
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

	now := time.Now()
	otp := &user.OTP{
		Id:              sdk.Id("otp"),
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

	if err := s.dispatchOtp(r.Context(), otp, code); err != nil {
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
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func (s *UserService) dispatchOtp(
	ctx context.Context,
	otp *user.OTP,
	code string,
) error {
	if otp.ContactPlatform == "" {
		return fmt.Errorf("contact platform is empty")
	}

	// @todo only email is supported now
	if otp.ContactPlatform != "email" {
		return fmt.Errorf("unsupported contact platform '%s'", otp.ContactPlatform)
	}

	// @todo switch to app from request

	subject := "Your one-time password"

	expiresIn := time.Until(otp.ExpiresAt).Round(time.Second)
	if expiresIn < 0 {
		expiresIn = 0
	}

	body := fmt.Sprintf("Your verification code is %s. It expires in %s.", code, expiresIn)

	emailReq := openapi.EmailSvcSendEmailRequest{
		To:      []string{otp.ContactId},
		Subject: subject,
		Body:    body,
	}
	contentType := "text/plain"
	emailReq.ContentType = openapi.PtrString(contentType)

	if !s.options.Test {
		_, _, err := s.options.ClientFactory.
			Client(client.WithToken(s.token)).
			EmailSvcAPI.
			SendEmail(ctx).
			Body(emailReq).
			Execute()
		if err != nil {
			return errors.Wrap(err, "email svc send email")
		}
	}

	return nil
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
