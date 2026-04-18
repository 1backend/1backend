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

type Organization struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	AppId string `json:"appId,omitempty"`

	Id string `json:"id" binding:"required"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Full name of the organization
	Name string `json:"name" example:"Acme Corporation" binding:"required"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug" example:"acme-corporation" binding:"required"`

	ThumbnailFileId string `json:"thumbnailFileId,omitempty" example:"file_fQDxusW8og"`
}

func (o *Organization) GetId() string {
	if o.InternalId == "" {
		panic("organization has no internal id")
	}

	return o.InternalId
}

type SaveOrganizationRequest struct {
	Id string `json:"id"`

	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug" binding:"required"`

	// Full name of the organization.
	Name string `json:"name,omitempty"`

	ThumbnailFileId string `json:"thumbnailFileId,omitempty" example:"file_fQDxusW8og"`

	// If true, the organization becomes the active organization for the
	// caller's current device and a fresh token is issued.
	Activate bool `json:"activate" example:"true"`
}

type SaveOrganizationResponse struct {
	Organization Organization `json:"organization" binding:"required"`

	// A fresh token is returned only when the organization is activated.
	Token *Token `json:"token,omitempty"`
}

type ActivateOrganizationRequest struct {
	OrganizationId string `json:"organizationId" binding:"required"`
}

type ActivateOrganizationResponse struct {
	Token Token `json:"token" binding:"required"`
}

type DeactivateOrganizationResponse struct {
	Token Token `json:"token" binding:"required"`
}

type ListOrganizationsRequest struct {
	Ids []string `json:"ids"`

	Limit int `json:"limit"`

	All bool `json:"all,omitempty"`

	// Organizations by default come back ordered
	// desc by `createdAt` field.
	AfterTime time.Time `json:"afterTime"`
}

type ListOrganizationsResponse struct {
	Organizations []Organization `json:"organizations"`
}
