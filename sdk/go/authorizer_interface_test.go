package sdk_test

import (
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/stretchr/testify/require"
)

func TestOrgExtraction(t *testing.T) {
	var auth sdk.Authorizer
	auth = sdk.AuthorizerImpl{}
	require.Equal(t, true, auth != nil)

	roleIds := []string{
		"user-svc:org:{org_dBZRCej3fo}:admin",
		"user-svc:org:{org_dBZRCej3fo}:member",
	}

	roles := sdk.ExtractOrganizationRoles(roleIds)

	require.Equal(t, 2, len(roles))
	require.Equal(t, "admin", roles[0])
	require.Equal(t, "member", roles[1])
}
