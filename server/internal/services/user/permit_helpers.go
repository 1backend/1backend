/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"sort"
	"strings"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

func normalizePermissions(permission string, permissions []string) []string {
	ret := make([]string, 0, len(permissions)+1)
	seen := map[string]struct{}{}

	add := func(value string) {
		if value == "" {
			return
		}
		if _, ok := seen[value]; ok {
			return
		}
		seen[value] = struct{}{}
		ret = append(ret, value)
	}

	add(permission)
	for _, permission := range permissions {
		add(permission)
	}

	return ret
}

func normalizePermitInputPermissions(permit *user.PermitInput) []string {
	if permit == nil {
		return nil
	}

	return normalizePermissions(permit.Permission, permit.Permissions)
}

func permitPermissions(permit *user.Permit) []string {
	if permit == nil {
		return nil
	}

	return normalizePermissions(permit.Permission, permit.Permissions)
}

func canonicalPermissionFields(permissions []string) (string, []string) {
	if len(permissions) == 0 {
		return "", nil
	}
	if len(permissions) == 1 {
		return permissions[0], nil
	}

	return "", append([]string(nil), permissions...)
}

func permitIdentityKey(appId string, roles, slugs, permissions []string) string {
	return strings.Join([]string{
		"app=" + appId,
		"roles=" + stableSliceKey(roles),
		"slugs=" + stableSliceKey(slugs),
		"permissions=" + stableSliceKey(permissions),
	}, "\n")
}

func stableSliceKey(values []string) string {
	if len(values) == 0 {
		return ""
	}

	cpy := append([]string(nil), values...)
	sort.Strings(cpy)

	return strings.Join(cpy, "\x1f")
}
