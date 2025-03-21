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

	require.Equal(t, 1, len(roles))
	require.Equal(t, "admin", roles["org_dBZRCej3fo"][0])
	require.Equal(t, "member", roles["org_dBZRCej3fo"][1])
}
