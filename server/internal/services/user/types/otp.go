/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import (
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
)

const (
	// OTPPurposeLogin is used for issuing authentication tokens without a password.
	OTPPurposeLogin = "login"
	// OTPPurposeVerifyContact is used for verifying a contact method such as an email address.
	OTPPurposeVerifyContact = "verify-contact"
)

// OTP represents a time-bound one-time password that can be delivered to a user contact.
type OTP struct {
	// Unique identifier of the OTP.
	Id string `json:"id" binding:"required"`

	CreatedAt  time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt  time.Time  `json:"updatedAt" binding:"required"`
	VerifiedAt *time.Time `json:"verifiedAt,omitempty"`

	// ContactId identifies the contact (e.g. email) to which the OTP was sent.
	ContactId string `json:"contactId" binding:"required"`

	// UserId is the identifier of the user owning the contact.
	UserId string `json:"userId" binding:"required"`

	// Purpose distinguishes the intent of the OTP (e.g. login or contact verification).
	Purpose string `json:"purpose" binding:"required"`

	// AppId associates the OTP with an app for login purposes.
	AppId string `json:"appId,omitempty"`
	// AppHost is stored for convenience when displaying or auditing OTP usage.
	AppHost string `json:"appHost,omitempty"`

	// ExpiresAt defines when the OTP becomes invalid.
	ExpiresAt time.Time `json:"expiresAt" binding:"required"`

	// CodeHash stores the hashed OTP code.
	CodeHash string `json:"-"`

	// TestCode stores the plain OTP code in test environments to aid automated verification.
	TestCode string `json:"testCode,omitempty"`
}

func (o OTP) GetId() string {
	return o.Id
}

// Indexes defines helpful indexes for querying OTPs.
func (o OTP) Indexes() []datastore.Index {
	return []datastore.Index{
		{
			Fields:    []string{"contactId", "purpose"},
			Ascending: true,
		},
		{
			Fields:    []string{"createdAt"},
			Ascending: false,
		},
	}
}

// RequestOTPRequest represents an OTP generation request.
type RequestOTPRequest struct {
	// ContactId is the identifier of the contact that should receive the OTP.
	ContactId string `json:"contactId" binding:"required"`
	// Purpose differentiates the OTP intent. Defaults to `login`.
	Purpose string `json:"purpose,omitempty"`
	// AppHost is required for login OTPs so that a token can be issued for the correct app.
	AppHost string `json:"appHost,omitempty"`
}

// RequestOTPResponse contains metadata about the generated OTP.
type RequestOTPResponse struct {
	OtpId     string    `json:"otpId"`
	ExpiresAt time.Time `json:"expiresAt"`
	// Code is only returned in test environments to facilitate end-to-end testing.
	Code string `json:"code,omitempty"`
}

// VerifyOTPRequest represents a request to validate an OTP.
type VerifyOTPRequest struct {
	// OtpId identifies the OTP that is being verified.
	OtpId string `json:"otpId" binding:"required"`
	// ContactId, when provided, must match the OTP contact to avoid leaking identifiers.
	ContactId string `json:"contactId,omitempty"`
	// Purpose optionally asserts the expected purpose of the OTP.
	Purpose string `json:"purpose,omitempty"`
	// Code is the one-time password received by the user.
	Code string `json:"code" binding:"required"`
	// Device identifies the device requesting a login token. Defaults to "unknown".
	Device string `json:"device,omitempty"`
}

// VerifyOTPResponse returns whether the OTP was successfully validated and, if applicable, a session token.
type VerifyOTPResponse struct {
	Verified bool   `json:"verified"`
	Token    *Token `json:"token,omitempty"`
}
