/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type Token struct {
	InternalId string `json:"internalId,omitempty" swaggerignore:"true"`

	AppId string `json:"appId" binding:"required"`

	Id string `json:"id" binding:"required"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	UserId string `json:"userId,omitempty" binding:"required"`

	// Token is a signed JWT used to authenticate the user without querying the User Svc.
	// You can verify it using the public key at `/user-svc/public-key`.
	//
	// The token is just a JSON object with fields like:
	// - "oui": the user ID (e.g., "usr_dC4K75Cbp6")
	// - "olu": the user slug (e.g., "test-user-slug-0")
	// - "oro": a list of roles, such as:
	//   - "user-svc:user"
	//   - "user-svc:org:{org_dC4K7NNDCG}:user"
	Token string `json:"token,omitempty" binding:"required"`

	// The device the token is associated with.
	// This in combination with LastRefreshedAt can be used to
	// determine if the token is still in use, and lets us prune unused tokens.
	Device string `json:"device" binding:"required"`

	// The last time the token was refreshed.
	// This is used to determine if the token is still in use.
	LastRefreshedAt *time.Time `json:"lastRefreshedAt,omitempty"`

	ExpiresAt time.Time `json:"expiresAt" binding:"required"`

	// Active tokens contain the most up-to-date information.
	// When a user's role changes—due to role assignment, organization
	// creation/assignment, etc.—all existing tokens are marked inactive.
	// Active tokens are reused during login, while inactive tokens
	// that have been recently refreshed (being used still) are kept for further refreshing
	// (unless `OB_TOKEN_AUTO_REFRESH_OFF` is set to true, old tokens can be refreshed indefinitely.)
	//
	// Active tokens contain the most up-to-date information.
	// When a user's role changes—due to role assignment, organization
	// creation/assignment, etc.—all existing tokens are marked inactive.
	// Active tokens are reused during login, while inactive tokens
	// that have been recently refreshed (see `lastRefreshedAt` field) and are still in use are retained for further refreshing.
	// (Unless `OB_TOKEN_AUTO_REFRESH_OFF` is set to true, in which case old tokens can be refreshed indefinitely.)
	Active bool `json:"active,omitempty"`

	App *App `json:"app,omitempty"`
}

func (c *Token) GetId() string {
	return c.Id
}

type RefreshTokenRequest struct {
}

type RefreshTokenResponse struct {
	Token *Token `json:"token" binding:"required"`
}

type RevokeTokensRequest struct {
	Device string `json:"device,omitempty"`

	// Only used by admins (or whoever has the `user-svc:token:revoke` permission
	// revoke tokens for other users
	UserId string `json:"userId,omitempty"`

	// If true, all tokens for the user will be revoked.
	AllTokens bool `json:"allTokens,omitempty"`
}

type RevokeTokensResponse struct{}

type ExchangeTokenRequest struct {
	// NewApp is the app of the new token that will be returned by this endpoint.
	NewAppHost string `json:"newAppHost,omitempty"`

	NewAppId string `json:"newAppId,omitempty"`

	// NewDevice. If not provided, the device of the original token will be used.
	NewDevice string `json:"newDevice,omitempty"`
}

type ExchangeTokenResponse struct {
	// Token is the new token that will be returned by this endpoint.
	Token *Token `json:"token" binding:"required"`
}
