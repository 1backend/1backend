/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice_test

import (
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/stretchr/testify/require"
)

func TestOwnsRole(t *testing.T) {
	t.Run("prefixed slug user owns role", func(t *testing.T) {
		owns := sdk.OwnsRole(&sdk.Claims{
			Slug: "user-svc",
		}, "user-svc:admin")

		require.Equal(t, true, owns)
	})

	t.Run("org user does not own org user role", func(t *testing.T) {
		owns := sdk.OwnsRole(&sdk.Claims{
			RoleIds: []string{"user-svc:org:{abc}:user"},
			Slug:    "some-svc",
		}, "user-svc:org:{abc}:user")

		require.Equal(t, false, owns)
	})

	t.Run("org admin owns org user role", func(t *testing.T) {
		owns := sdk.OwnsRole(&sdk.Claims{
			RoleIds: []string{"user-svc:org:{abc}:admin"},
			Slug:    "some-svc",
		}, "user-svc:org:{abc}:user")

		require.Equal(t, true, owns)
	})

	t.Run("test for prefix logic error", func(t *testing.T) {
		owns := sdk.OwnsRole(&sdk.Claims{
			RoleIds: []string{"user-svc:org:{abc}:admin"},
			Slug:    "some-svc",
		}, "user-svc:org:{abcd}:user")

		require.Equal(t, false, owns)
	})
}
