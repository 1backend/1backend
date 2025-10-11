/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

// Password (password hash), is separate from the user record
// so that we avoid accidentally exposing or overwriting the password.
type Password struct {
	Id string `json:"id" binding:"required"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	UserId       string `json:"userId" binding:"required"`
	PasswordHash string `json:"passwordHash,omitempty"`
}

func (p *Password) GetId() string {
	return p.Id
}
