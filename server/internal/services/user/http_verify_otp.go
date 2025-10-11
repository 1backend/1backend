/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// VerifyOTP validates an OTP and optionally issues an auth token when the OTP represents a login.
func (s *UserService) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	req := user.VerifyOTPRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode OTP verification request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	otpI, found, err := s.otpStore.Query(
		datastore.Equals(datastore.Field("id"), req.OtpId),
	).FindOne()
	if err != nil {
		logger.Error("Failed to query OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	otp := otpI.(*user.OTP)

	if req.ContactId != "" && req.ContactId != otp.ContactId {
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	if req.Purpose != "" && req.Purpose != otp.Purpose {
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	if otp.VerifiedAt != nil {
		endpoint.WriteString(w, http.StatusUnauthorized, "OTP already used")
		return
	}

	if time.Now().After(otp.ExpiresAt) {
		endpoint.WriteString(w, http.StatusUnauthorized, "OTP expired")
		return
	}

	valid := checkPasswordHash(req.Code, otp.CodeHash)
	if s.options.Test && otp.TestCode != "" {
		valid = valid || otp.TestCode == req.Code
	}

	if !valid {
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	now := time.Now()
	update := map[string]any{
		"verifiedAt": &now,
		"updatedAt":  now,
	}
	if err := s.otpStore.Query(
		datastore.Equals(datastore.Field("id"), otp.Id),
	).UpdateFields(update); err != nil {
		logger.Error("Failed to update OTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	otp.VerifiedAt = &now
	otp.UpdatedAt = now

	if otp.ContactId != "" {
		if err := s.contactStore.Query(
			datastore.Equals(datastore.Field("id"), otp.ContactId),
		).UpdateFields(map[string]any{
			"verified":  true,
			"updatedAt": now,
		}); err != nil {
			logger.Error("Failed to update contact after OTP verification", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
	}

	var token *user.Token
	if otp.AppId != "" {
		userI, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("id"), otp.UserId),
		).FindOne()
		if err != nil {
			logger.Error("Failed to load user for OTP login", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
		if !found {
			endpoint.InternalServerError(w)
			return
		}

		usr := userI.(*user.User)
		device := req.Device
		if device == "" {
			device = unknownDevice
		}

		token, err = s.issueToken(otp.AppId, usr, device)
		if err != nil {
			logger.Error("Failed to issue token for OTP login", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
	}

	endpoint.WriteJSON(w, http.StatusOK, user.VerifyOTPResponse{
		Verified: true,
		Token:    token,
	})
}
