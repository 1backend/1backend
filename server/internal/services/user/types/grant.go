/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package user_svc

// Grant represents a permission assignment that can be applied to slugs or role IDs.
// Originally designed for service-to-service authentication via slugs, grants were later
// extended to support roles due to their convenience in CLI and infrastructure-as-code workflows.
// While UserSvc.AssignPermissions is the standard method for assigning permissions to roles,
// grants provide an alternative, streamlined approach.
type Grant struct {
	Id         string `json:"id,omitempty"`
	Permission string `json:"permission" binding:"required"`

	// Slugs that have been granted the specified permission.
	Slugs []string `json:"slugs,omitempty"`

	// Role IDs that have been granted the specified permission.
	//
	// Originally, grants were designed for slugs to facilitate service-to-service calls.
	// Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.
	//
	// Alternatively, permissions can be assigned to roles using UserSvc.AssignPermissions.
	// Grants currently offer a more streamlined approach, though this may evolve over time.
	Roles []string `json:"roles,omitempty"`
}

func (g Grant) GetId() string {
	return g.Id
}

type ListGrantsRequest struct {
	Permission string `json:"permission,omitempty"`
	Slug       string `json:"slug,omitempty"`
}

type ListGrantsResponse struct {
	Grants []*Grant `json:"grants,omitempty"`
}

type SaveGrantsRequest struct {
	Grants []*Grant `json:"grants,omitempty"`
}

type SaveGrantsResponse struct{}
