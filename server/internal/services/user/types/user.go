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
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type User struct {
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the user.
	Name string `json:"name,omitempty" example:"Jane Doe"`

	// URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
	Slug string `json:"slug" example:"jane-doe" binding:"required"`

	PasswordHash string `json:"passwordHash,omitempty"`

	ThumbnailFileId string `json:"thumbnailFileId,omitempty" example:"file_fQDyi1xdHK"`

	Labels map[string]string `json:"labels,omitempty"`
}

type UserRecord struct {
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// Full name of the user.
	Name string `json:"name,omitempty" example:"Jane Doe"`

	// URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
	Slug string `json:"slug" example:"jane-doe" binding:"required"`

	ContactIds []string `json:"contactIds,omitempty"`
	Roles      []string `json:"roles,omitempty"`
}

func (c *User) GetId() string {
	return c.Id
}

func (c *User) GetUpdatedAt() string {
	return c.Id
}

type ReadSelfRequest struct{}

type ReadSelfResponse struct {
	// The user who made the request.
	User *User `json:"user" binding:"required"`

	// Roles the token has that made this request.
	Roles []string `json:"roles,omitempty"`

	// Organizations of the caller user.
	Organizations []*Organization `json:"organizations,omitempty"`

	// Contacts of the caller user.
	Contacts []Contact `json:"contacts,omitempty"`

	// Active organization of the caller user, if it has any.
	ActiveOrganizationId string `json:"activeOrganizationId,omitempty"`
}

type RegisterRequest struct {
	Name string `json:"name,omitempty"`

	// Slug is a URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
	// Required due to its central role in the platform.
	// If your project has no use for a slug, just derive it from the email or similar.
	Slug string `json:"slug" binding:"required"`

	Contact Contact `json:"contact,omitempty"`

	Password string `json:"password,omitempty"`
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

type SaveUserRequest struct {
	Name            string `json:"name,omitempty"`
	ThumbnailFileId string `json:"thumbnailFileId,omitempty" example:"file_fQDxusW8og"`
}

type SaveUserResponse struct {
}

type SaveSelfRequest struct {
	Name            string             `json:"name,omitempty"`
	ThumbnailFileId string             `json:"thumbnailFileId,omitempty" example:"file_fQDxusW8og"`
	Labels          *map[string]string `json:"labels,omitempty"`
}

type SaveSelfResponse struct {
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
}

type ChangePasswordResponse struct{}

type ResetPasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required"`
}

type ResetPasswordResponse struct{}

type ListUsersOrderByField string

const (
	ListUsersOrderByFieldCreatedAt ListUsersOrderByField = "createdAt"
	ListUsersOrderByFieldUpdatedAt ListUsersOrderByField = "updatedAt"
)

type ListUsersRequest struct {
	// Ids of the users to list.
	Ids []string `json:"ids,omitempty"`

	// ContactId is the id of the contact the user is associated with.
	// Will return a user list with one element if set.
	ContactId string `json:"contactId,omitempty"`

	Limit  int64 `json:"limit,omitempty" example:"10"`
	Offset int64 `json:"offset,omitempty" example:"0"`

	// AfterTime is a time in RFC3339 format.
	// It is used to paginate the results when the `orderByField` is set to `createdAt` or `updatedAt`.
	// The results will be returned after this time.
	AfterTime time.Time `json:"afterTime,omitempty"`

	OrderByDesc  bool                  `json:"orderByDesc,omitempty" example:"false"`
	OrderByField ListUsersOrderByField `json:"orderByField,omitempty" example:"createdAt"`

	// Count is a flag that indicates if the count of the users should be returned.
	Count bool `json:"count,omitempty" example:"false"`
}

type ListUsersResponse struct {
	Users []*UserRecord `json:"users,omitempty"`
	After time.Time     `json:"after,omitempty"`
	Count int64         `json:"count"`
}

type CreateUserRequest struct {
	User     *User     `json:"user,omitempty"`
	Contacts []Contact `json:"contacts,omitempty"`
	Password string    `json:"password,omitempty"`
	RoleIds  []string  `json:"roleIds,omitempty"`
}

type CreateUserResponse struct {
}

type DeleteUserRequest struct {
	UserId string `json:"userId,omitempty"`
}

type DeleteUserResponse struct{}

type AssignRoleRequest struct {
}

type AssignRoleResponse struct {
}
