package user

import (
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/stretchr/testify/require"
)

func TestEnrollIdsForUserRole(t *testing.T) {
	t.Parallel()

	userId := "usr_target"
	otherUserId := "usr_other"

	require.Equal(t,
		[]string{"enr_1"},
		enrollIdsForUserRole([]openapi.UserSvcEnroll{
			{Id: "enr_1", UserId: &userId, Role: "role:a"},
			{Id: "enr_2", UserId: &userId, Role: "role:b"},
			{Id: "enr_3", UserId: &otherUserId, Role: "role:a"},
			{Id: "enr_4", Role: "role:a"},
		}, userId, "role:a"),
	)
}
