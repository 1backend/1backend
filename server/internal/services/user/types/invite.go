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

// Invite is a mechanism to grant roles to users who are not yet registered.
// If the user is already registered, an invite is still created, but the role
// is applied immediately by calling `Add Role to User`. It's important to note
// that while `Add Role to User` works by UserId, the invite is based on ContactId.
//
// To extend an expired invite, update the `expiresAt` field with a new date.
// If a user is deleted and later re-registers, the invite remains valid unless explicitly revoked.
//
// While admittedly informal, the term Invite is preferred over Invitation for brevity.
type Invite struct {
	Id string `json:"id" example:"inv_fIYPbMHIcI" binding:"required"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`

	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	AppliedAt time.Time `json:"appliedAt,omitempty"`

	// ContactId represents the recipient of the invite.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	ContactId string `json:"contactId" binding:"required"`

	// RoleId specifies the role to be assigned to the ContactId.
	// Callers can only assign roles they own, identified by their service slug
	// (e.g., if "my-service" creates an invite, the role must be "my-service:admin").
	// Dynamic organization roles can also be assigned
	// (e.g., "user-svc:org:{%orgId}:admin" or "user-svc:org:{%orgId}:user"),
	// but in this case, the caller must be an admin of the target organization.
	RoleId string `json:"roleId" binding:"required"`
}

func (i Invite) GetId() string {
	return i.Id
}

type PartialInvite struct {
	Id        string `json:"id,omitempty" example:"inv_fIYPbMHIcI"`
	ContactId string `json:"contactId" binding:"required"`
	RoleId    string `json:"roleId" binding:"required"`
}

type SaveInvitesRequest struct {
	Invites []Invite `json:"invites" binding:"required"`
}

type SaveInvitesResponse struct {
	Invites []Invite `json:"invites" binding:"required"`
}

type ListInvitesRequest struct {
	ContactId string `json:"contactId,omitempty"`
	RoleId    string `json:"roleId,omitempty"`
}

type ListInvitesResponse struct {
	Invites []Invite `json:"invites" binding:"required"`
}
