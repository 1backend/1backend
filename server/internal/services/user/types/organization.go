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

type OrganizationUserLink struct {
	// organizationId:userId
	Id string `json:"id,omitempty"`

	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	OrganizationId string `json:"organizationId,omitempty"`
	UserId         string `json:"userId,omitempty"`

	// Active/default organization for a user.
	// There can only be one per user.
	Active bool `json:"active,omitempty"`
}

func (o *OrganizationUserLink) GetId() string {
	return o.Id
}

type Organization struct {
	Id        string     `json:"id" binding:"required"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the organization
	Name string `json:"name" example:"Acme Corporation" binding:"required"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug" example:"acme-corporation" binding:"required"`
}

func (o *Organization) GetId() string {
	return o.Id
}

type CreateOrganizationRequest struct {
	Id string `json:"id"`

	// Full name of the organization.
	Name string `json:"name" binding:"required"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug" binding:"required"`
}

type CreateOrganizationResponse struct {
	Organization Organization `json:"organization" binding:"required"`

	// Due to the nature of JWT tokens, the token must be refreshed after
	// creating an organization, as dynamic organization roles are embedded in it.
	Token AuthToken `json:"token" binding:"required"`
}

type AddUserToOrganizationRequest struct {
}

type AddUserToOrganizationResponse struct {
}

type RemoveUserFromOrganizationRequest struct {
}

type RemoveUserFromOrganizationResponse struct {
}
