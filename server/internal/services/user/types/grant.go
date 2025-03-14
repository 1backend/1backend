/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package user_svc

type Grant struct {
	Id           string `json:"id,omitempty"`
	PermissionId string `json:"permissionId,omitempty"`

	// Slugs who are granted the PermissionId
	Slugs []string `json:"slugs,omitempty"`
}

func (g Grant) GetId() string {
	return g.Id
}

type ListGrantsRequest struct {
	PermissionId string `json:"permissionId,omitempty"`
	Slug         string `json:"slug,omitempty"`
}

type ListGrantsResponse struct {
	Grants []*Grant `json:"grants,omitempty"`
}

type SaveGrantsRequest struct {
	Grants []*Grant `json:"grants,omitempty"`
}

type SaveGrantsResponse struct{}
