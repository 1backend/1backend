/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type Membership struct {
	Id  string `json:"id,omitempty"`
	App string `json:"app,omitempty" example:"unnamed"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	OrganizationId string `json:"organizationId,omitempty"`
	UserId         string `json:"userId,omitempty"`

	// Marks the active (default) organization for a user.
	// Only one organization can be active per user.
	//
	// The ID of this active organization will be stored in the JWT token
	// upon login or refresh under the JSON key `oao`
	// (`1Backend ActiveOrganizationId`).
	Active bool `json:"active,omitempty"`
}

func (o *Membership) GetId() string {
	return o.Id
}

type SaveMembershipsRequest struct {
}

type SaveMembershipsResponse struct {
}

type DeleteMembershipRequest struct {
}

type DeleteMembershipResponse struct {
}
