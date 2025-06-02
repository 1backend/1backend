/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

// Permit is a mechanism to give users or roles permissions to perform actions defined by the `Permission` field.
type Permit struct {
	Id  string `json:"id" example:"inv_fIYPbMHIcI" binding:"required"`
	App string `json:"app" example:"unnamed,omitempty"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`

	Permission string `json:"permission" binding:"required"`

	// Slugs that have been permited the specified permission.
	Slugs []string `json:"slugs,omitempty"`

	// Role IDs that have been permited the specified permission.
	//
	// Originally, permits were designed for slugs to facilitate service-to-service calls.
	// Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.
	Roles []string `json:"roles,omitempty"`
}

// PermitInput is the settable subset of Enroll, excluding system-managed fields.
type PermitInput struct {
	Id string `json:"id,omitempty" example:"inv_fIYPbMHIcI"`

	Permission string `json:"permission" binding:"required"`

	// Slugs that have been permited the specified permission.
	Slugs []string `json:"slugs,omitempty"`

	// Role IDs that have been permited the specified permission.
	//
	// Originally, permits were designed for slugs to facilitate service-to-service calls.
	// Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.
	Roles []string `json:"roles,omitempty"`
}

func (g Permit) GetId() string {
	return g.Id
}

type ListPermitsRequest struct {
	Permission string `json:"permission,omitempty"`
	Slug       string `json:"slug,omitempty"`
}

type ListPermitsResponse struct {
	Permits []*Permit `json:"permits,omitempty"`
}

type SavePermitsRequest struct {
	Permits []*PermitInput `json:"permits,omitempty"`
}

type SavePermitsResponse struct{}
