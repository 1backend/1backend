/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package user_svc

import (
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetUsersOptions struct {
	Query *datastore.Query `json:"query"`
}

type User struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the organization.
	Name string `json:"name,omitempty" example:"Jane Doe"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `user`.
	Slug string `json:"slug,omitempty" example:"jane-doe"`

	// Contacts are used for login and identification purposes.
	Contacts []Contact `json:"contacts,omitempty"`

	PasswordHash string `json:"passwordHash,omitempty"`
}

type UserRoleLink struct {
	// userId:roleId
	Id string `json:"id,omitempty"`

	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	RoleId string `json:"roleId,omitempty"`
	UserId string `json:"userId,omitempty"`
}

func (u *UserRoleLink) GetId() string {
	return u.Id
}

func (c *User) GetId() string {
	return c.Id
}

func (c *User) GetUpdatedAt() string {
	return c.Id
}

type ReadUserByTokenRequest struct{}

type ReadUserByTokenResponse struct {
	User                 *User           `json:"user,omitempty"`
	Organizations        []*Organization `json:"organizations,omitempty"`
	ActiveOrganizationId string          `json:"activeOrganizationId,omitempty"`
}

type RegisterRequest struct {
	Name     string  `json:"name,omitempty"`
	Slug     string  `json:"slug,omitempty"`
	Contact  Contact `json:"contact,omitempty"`
	Password string  `json:"password,omitempty"`
}

type RegisterResponse struct {
	Token *AuthToken `json:"token,omitempty"`
}

type LoginRequest struct {
	Slug     string `json:"slug,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token *AuthToken `json:"token,omitempty"`
}

type SaveProfileRequest struct {
	Slug string `json:"slug,omitempty"`
	Name string `json:"name,omitempty"`
}

type SaveProfileResponse struct {
}

type ChangePasswordRequest struct {
	Slug            string `json:"slug,omitempty"`
	CurrentPassword string `json:"currentPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
}

type ChangePasswordResponse struct{}

type ResetPasswordRequest struct {
	Slug        string `json:"slug,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type ResetPasswordResponse struct{}

type GetUsersRequest struct {
	Query *datastore.Query `json:"query"`
}

type GetUsersResponse struct {
	Users []*User   `json:"users,omitempty"`
	After time.Time `json:"after,omitempty"`
	Count int64     `json:"count"`
}

type CreateUserRequest struct {
	User     *User    `json:"user,omitempty"`
	Password string   `json:"password,omitempty"`
	RoleIds  []string `json:"roleIds,omitempty"`
}

type CreateUserResponse struct {
}

type DeleteUserRequest struct {
	UserId string `json:"userId,omitempty"`
}

type DeleteUserResponse struct{}
