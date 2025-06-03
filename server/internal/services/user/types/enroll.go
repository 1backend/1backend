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
)

// Enroll (enrollment) is a mechanism to give registered or not yet registered roles.
type Enroll struct {
	Id string `json:"id" example:"inv_fIYPbMHIcI" binding:"required"`

	// App of the enroll.
	// Use `*` to match all apps, such as when bootstrapping
	// in services.
	App string `json:"app" example:"unnamed,omitempty"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`

	// Role specifies the role to be assigned to the ContactId.
	// Callers can only assign roles they own, identified by their service slug
	// (e.g., if "my-service" creates an enroll, the role must be "my-service:admin").
	// Dynamic organization roles can also be assigned
	// (e.g., "user-svc:org:{%orgId}:admin" or "user-svc:org:{%orgId}:user"),
	// but in this case, the caller must be an admin of the target organization.
	Role string `json:"role" binding:"required"`

	// ContactId is the the recipient of the enroll.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	ContactId string `json:"contactId,omitempty"`

	// UserId is the recipient of the enroll.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	UserId string `json:"userId,omitempty"`

	// CreatedBy contains the ID of the user who created the Enroll.
	CreatedBy string `json:"createdBy,omitempty"`
}

func (i Enroll) GetId() string {
	return i.Id
}

// EnrollInput is the settable subset of Enroll, excluding system-managed fields.
type EnrollInput struct {
	Id  string `json:"id,omitempty" example:"inv_fIYPbMHIcI"`
	App string `json:"app" example:"unnamed,omitempty"`

	// ContactId is the the recipient of the enroll.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	ContactId string `json:"contactId,omitempty"`

	// UserId is the recipient of the enroll.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	UserId string `json:"userId,omitempty"`

	Role string `json:"role" binding:"required"`
}

type SaveEnrollsRequest struct {
	Enrolls []EnrollInput `json:"enrolls" binding:"required"`
}

type SaveEnrollsResponse struct {
	Enrolls []Enroll `json:"enrolls" binding:"required"`
}

type ListEnrollsRequest struct {
	UserId    string `json:"userId,omitempty"`
	ContactId string `json:"contactId,omitempty"`
	Role      string `json:"role,omitempty"`
}

type ListEnrollsResponse struct {
	Enrolls []Enroll `json:"enrolls" binding:"required"`
}
