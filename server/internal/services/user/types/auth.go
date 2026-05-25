/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type AuthState struct {
	Id        string    `json:"id" binding:"required"`
	Provider  string    `json:"provider" binding:"required"`
	AppHost   string    `json:"appHost" binding:"required"`
	Device    string    `json:"device,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	ReturnTo  string    `json:"returnTo,omitempty"`
	Nonce     string    `json:"nonce,omitempty"`
	Used      bool      `json:"used,omitempty"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	ExpiresAt time.Time `json:"expiresAt" binding:"required"`
}

func (a *AuthState) GetId() string {
	return a.Id
}

type ContactAuthLoginRequest struct {
	AppHost string `json:"appHost" binding:"required" example:"shoes.com"`

	// Token is the provider token. For OIDC providers this is an ID token.
	// For Facebook this is a user access token.
	Token string `json:"token,omitempty"`

	// IDToken is accepted as an alias for Token for OIDC client integrations.
	IDToken string `json:"idToken,omitempty"`

	// AccessToken is accepted as an alias for Token for OAuth2 providers.
	AccessToken string `json:"accessToken,omitempty"`

	// Optional slug used only when the verified contact has no 1backend user yet.
	// If omitted, 1backend derives one from the contact id.
	Slug string `json:"slug,omitempty"`

	Device string `json:"device,omitempty"`

	TOTPCode string `json:"totpCode,omitempty"`
}

type ContactAuthLoginResponse struct {
	Token *Token `json:"token,omitempty"`
}

type ContactAuthProviderInfo struct {
	Id       string   `json:"id" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Kind     string   `json:"kind" binding:"required"`
	Scopes   []string `json:"scopes,omitempty"`
	StartUrl string   `json:"startUrl,omitempty"`
	LoginUrl string   `json:"loginUrl,omitempty"`
}

type ListContactAuthProvidersResponse struct {
	Providers []ContactAuthProviderInfo `json:"providers,omitempty"`
}
