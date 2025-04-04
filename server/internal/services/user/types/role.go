/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

var (
	RoleAdmin = "user-svc:admin"
	RoleUser  = "user-svc:user"
)

type PermissionRoleLink struct {
	// permissionId:roleId
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Permission string `json:"permission,omitempty"`
	Role       string `json:"role,omitempty"`
}

func (c *PermissionRoleLink) GetId() string {
	return c.Id
}

type UnassignRoleRequest struct {
	UserId string `json:"userId,omitempty"`
	Role   string `json:"role,omitempty"`
}

type UnassignRoleResponse struct{}

type ListRolesRequest struct {
}

type ListRolesResponse struct {
	Roles []string `json:"roles,omitempty"`
}
