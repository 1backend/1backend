/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type AuthToken struct {
	Id        string     `json:"id" binding:"required"`
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

	// Active tokens contain the most up-to-date information.
	// When a user's role changes—due to role assignment, organization
	// creation/assignment, etc.—all existing tokens are marked inactive.
	// Active tokens are reused during login, while inactive tokens
	// are retained for historical reference.
	Active bool `json:"active,omitempty"`
}

func (c *AuthToken) GetId() string {
	return c.Id
}

type RefreshTokenRequest struct {
}

type RefreshTokenResponse struct {
	Token *AuthToken `json:"token" binding:"required"`
}
