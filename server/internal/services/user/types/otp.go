/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type OTP struct {
	Id         string    `json:"id"  binding:"required"`
	CreatedAt  time.Time `json:"createdAt" binding:"required"`
	UpdatedAt  time.Time `json:"updatedAt" binding:"required"`
	ExpiresAt  time.Time `json:"expiresAt" binding:"required"`
	VerifiedAt time.Time `json:"verifiedAt,omitempty"`
	Used       bool      `json:"used,omitempty"`
	Attempts   int       `json:"attempts,omitempty"`

	CodeHash string `json:"codeHash,omitempty"`

	ContactId       string `json:"contactId" binding:"required"`
	ContactPlatform string `json:"contactPlatform" binding:"required"`
	ContactHandle   string `json:"contactHandle,omitempty"`
}

func (o *OTP) GetId() string {
	return o.Id
}

type OTPInput struct {
	ContactId       string `json:"contactId" binding:"required"`
	ContactPlatform string `json:"contactPlatform" binding:"required"`
	ContactHandle   string `json:"contactHandle,omitempty"`
}

type SendOtpRequest struct {
	AppHost string `json:"appHost" binding:"required"`

	OTPInput
}

type SendOtpResponse struct {
	OtpId string `json:"otpId" binding:"required"`

	// In test mode, the OTP code is returned in the response for easier testing.
	Code string `json:"code,omitempty"`
}
