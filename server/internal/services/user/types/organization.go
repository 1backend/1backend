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
	Id        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the organization
	Name string `json:"name,omitempty" example:"Acme Corporation"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug,omitempty" example:"acme-corporation"`
}

func (o *Organization) GetId() string {
	return o.Id
}

type CreateOrganizationRequest struct {
	Id string `json:"id,omitempty"`

	// Full name of the organization.
	Name string `json:"name,omitempty"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug,omitempty"`
}

type CreateOrganizationResponse struct{}

type AddUserToOrganizationRequest struct {
	UserId string `json:"userId,omitempty"`
}

type AddUserToOrganizationResponse struct {
}

type RemoveUserFromOrganizationRequest struct {
}

type RemoveUserFromOrganizationResponse struct {
}
