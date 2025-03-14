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
	Id        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	UserId string `json:"userId,omitempty"`
	Token  string `json:"token,omitempty"`

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

func (c *AuthToken) GetUpdatedAt() string {
	return c.Id
}
