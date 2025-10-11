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
	"encoding/json"
	"fmt"
	"log/slog"
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

// RequestOTP generates an OTP for the specified contact and dispatches it via the Email Service.
func (s *UserService) RequestOTP(w http.ResponseWriter, r *http.Request) {
	req := user.RequestOTPRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode OTP request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.ContactId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "contactId is required")
		return
	}

	purpose := req.Purpose
	if purpose == "" {
		purpose = user.OTPPurposeLogin
	}

	if purpose == user.OTPPurposeLogin && req.AppHost == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "appHost is required for login OTPs")
		return
	}

	contactI, found, err := s.contactStore.Query(
		datastore.Equals(datastore.Field("id"), req.ContactId),
	).FindOne()
	if err != nil {
		logger.Error("Failed to query contact for OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.WriteString(w, http.StatusNotFound, "Contact not found")
		return
	}

	contact := contactI.(*user.Contact)
	if contact.Platform != string(user.ContactPlatformEmail) {
		endpoint.WriteString(w, http.StatusBadRequest, "Only email contacts support OTPs")
		return
	}

	userI, found, err := s.userStore.Query(
		datastore.Equals(datastore.Field("id"), contact.UserId),
	).FindOne()
	if err != nil {
		logger.Error("Failed to query user for OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.InternalServerError(w)
		return
	}
	usr := userI.(*user.User)

	appId := ""
	if req.AppHost != "" {
		app, err := s.getOrCreateApp(req.AppHost)
		if err != nil {
			logger.Error("Failed to get or create app for OTP", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
		appId = app.Id
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
		Id:        sdk.Id("otp"),
		CreatedAt: now,
		UpdatedAt: now,
		ContactId: contact.Id,
		UserId:    usr.Id,
		Purpose:   purpose,
		AppId:     appId,
		AppHost:   req.AppHost,
		ExpiresAt: now.Add(otpExpirationDelta),
		CodeHash:  hash,
	}
	if s.options.Test {
		otp.TestCode = code
	}

	if err := s.otpStore.Upsert(otp); err != nil {
		logger.Error("Failed to store OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	if err := s.dispatchOtp(r.Context(), contact, otp, code); err != nil {
		logger.Error("Failed to send OTP email", slog.Any("error", err))
		if derr := s.otpStore.Query(
			datastore.Equals(datastore.Field("id"), otp.Id),
		).Delete(); derr != nil {
			logger.Error("Failed to delete OTP after email error", slog.Any("error", derr))
		}
		endpoint.InternalServerError(w)
		return
	}

	response := user.RequestOTPResponse{
		OtpId:     otp.Id,
		ExpiresAt: otp.ExpiresAt,
	}
	if s.options.Test {
		response.Code = code
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func (s *UserService) dispatchOtp(
	ctx context.Context,
	contact *user.Contact,
	otp *user.OTP,
	code string,
) error {
	subject := "Your 1Backend verification code"
	if otp.Purpose == user.OTPPurposeLogin {
		subject = "Your 1Backend login code"
	}

	expiresIn := time.Until(otp.ExpiresAt).Round(time.Second)
	if expiresIn < 0 {
		expiresIn = 0
	}

	body := fmt.Sprintf("Your verification code is %s. It expires in %s.", code, expiresIn)

	emailReq := openapi.EmailSvcSendEmailRequest{
		To:      []string{contact.Id},
		Subject: subject,
		Body:    body,
	}
	contentType := "text/plain"
	emailReq.ContentType = openapi.PtrString(contentType)

	_, _, err := s.options.ClientFactory.
		Client(client.WithToken(s.token)).
		EmailSvcAPI.
		SendEmail(ctx).
		Body(emailReq).
		Execute()
	if err != nil {
		return errors.Wrap(err, "email svc send email")
	}

	return nil
}
