/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package auth_test

import (
	"testing"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/stretchr/testify/require"
)

func TestOrgExtraction(t *testing.T) {
	var authr auth.Authorizer
	authr = auth.AuthorizerImpl{}
	require.Equal(t, true, authr != nil)

	roleIds := []string{
		"user-svc:org:{org_dBZRCej3fo}:admin",
		"user-svc:org:{org_dBZRCej3fo}:member",
	}

	roles := auth.ExtractOrganizationRoles(roleIds)

	require.Equal(t, 1, len(roles))
	require.Equal(t, "admin", roles["org_dBZRCej3fo"][0])
	require.Equal(t, "member", roles["org_dBZRCej3fo"][1])
}

func TestOwnsRole(t *testing.T) {
	t.Run("prefixed slug user owns role", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			Slug: "user-svc",
		}, "user-svc:admin")

		require.Equal(t, true, owns)

		owns = auth.OwnsRole(&auth.Claims{
			Slug: "user-svc",
		}, "user-svc:custom")

		require.Equal(t, true, owns)
	})

	t.Run("org user does not own org user role", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			RoleIds: []string{"user-svc:org:{abc}:user"},
			Slug:    "some-svc",
		}, "user-svc:org:{abc}:user")

		require.Equal(t, false, owns)
	})

	t.Run("org admin owns org user role", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			RoleIds: []string{"user-svc:org:{abc}:admin"},
			Slug:    "some-svc",
		}, "user-svc:org:{abc}:user")

		require.Equal(t, true, owns)
	})

	t.Run("test for prefix logic error", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			RoleIds: []string{"user-svc:org:{abc}:admin"},
			Slug:    "some-svc",
		}, "user-svc:org:{abcd}:user")

		require.Equal(t, false, owns)
	})

	t.Run("static admin should own", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			RoleIds: []string{"a-role:admin"},
			Slug:    "some-svc",
		}, "a-role:user")

		require.Equal(t, true, owns)
	})

	t.Run("static non-admin should not own", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			RoleIds: []string{"a-role:user"},
			Slug:    "some-svc",
		}, "a-role:user")

		require.Equal(t, false, owns)
	})

	t.Run("admin owns any role", func(t *testing.T) {
		owns := auth.OwnsRole(&auth.Claims{
			Slug:    "does-not-matter",
			RoleIds: []string{"user-svc:admin"},
		}, "anything")

		require.Equal(t, true, owns)
	})
}
