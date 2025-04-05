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

	// ContactId is the the recipient of the invite.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	ContactId string `json:"contactId,omitempty"`

	// UserId is the recipient of the invite.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	UserId string `json:"userId,omitempty"`

	// Role specifies the role to be assigned to the ContactId.
	// Callers can only assign roles they own, identified by their service slug
	// (e.g., if "my-service" creates an invite, the role must be "my-service:admin").
	// Dynamic organization roles can also be assigned
	// (e.g., "user-svc:org:{%orgId}:admin" or "user-svc:org:{%orgId}:user"),
	// but in this case, the caller must be an admin of the target organization.
	Role string `json:"role" binding:"required"`

	// OwnerIds specifies the users who created the invite.
	// If you create an invite that already exists for a given role and contact ID,
	// you get added to the list of owners.
	OwnerIds []string `json:"ownerIds" binding:"required"`
}

func (i Invite) GetId() string {
	return i.Id
}

// NewInvite is used to create a new invite.
// It is a subset of the Invite struct, as some fields are set automatically.
type NewInvite struct {
	Id string `json:"id,omitempty" example:"inv_fIYPbMHIcI"`

	// ContactId is the the recipient of the invite.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	ContactId string `json:"contactId,omitempty"`

	// UserId is the recipient of the invite.
	// If the user is already registered, the role is assigned immediately;
	// otherwise, it is applied upon registration.
	UserId string `json:"userId,omitempty"`

	Role string `json:"role" binding:"required"`
}

type SaveInvitesRequest struct {
	Invites []NewInvite `json:"invites" binding:"required"`
}

type SaveInvitesResponse struct {
	Invites []Invite `json:"invites" binding:"required"`
}

type ListInvitesRequest struct {
	UserId    string `json:"userId,omitempty"`
	ContactId string `json:"contactId,omitempty"`
	Role      string `json:"role,omitempty"`
}

type ListInvitesResponse struct {
	Invites []Invite `json:"invites" binding:"required"`
}
