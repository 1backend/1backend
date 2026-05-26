/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type TOTP struct {
	InternalId string `json:"internalId,omitempty"`

	AppId string `json:"appId,omitempty"`

	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	UserId string `json:"userId" binding:"required"`

	Issuer          string `json:"issuer" binding:"required"`
	AccountName     string `json:"accountName" binding:"required"`
	Secret          string `json:"secret,omitempty"`
	ProvisioningURI string `json:"provisioningUri,omitempty"`

	Enabled    bool       `json:"enabled,omitempty"`
	EnabledAt  *time.Time `json:"enabledAt,omitempty"`
	DisabledAt *time.Time `json:"disabledAt,omitempty"`
}

func (t *TOTP) GetId() string {
	return t.Id
}

type BeginTOTPSetupRequest struct {
	// Issuer is the optional service name shown by authenticator apps.
	// If omitted, 1Backend is used.
	Issuer string `json:"issuer,omitempty" example:"app.example.com"`

	// AccountName is an optional account label template shown by authenticator apps.
	// It may include $name, $slug, $contactId, $contactIds, $email, $phone, or contact platform placeholders.
	// If omitted, $slug is used for backward compatibility.
	AccountName string `json:"accountName,omitempty" example:"$email"`
}

type BeginTOTPSetupResponse struct {
	TOTPId          string `json:"totpId" binding:"required"`
	Secret          string `json:"secret" binding:"required"`
	ProvisioningURI string `json:"provisioningUri" binding:"required"`
	QRImagePath     string `json:"qrImagePath" binding:"required"`
}

type ReadTOTPStatusRequest struct {
	Email string `json:"email" binding:"required"`
}

type ReadTOTPStatusResponse struct {
	TOTPEnabled bool `json:"totpEnabled" binding:"required"`
}

type EnableTOTPRequest struct {
	TOTPId string `json:"totpId,omitempty"`
	Code   string `json:"code" binding:"required"`
}

type EnableTOTPResponse struct {
	Enabled bool `json:"enabled" binding:"required"`
}

type DisableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type DisableTOTPResponse struct{}
