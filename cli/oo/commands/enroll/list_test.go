package enroll

import (
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/stretchr/testify/require"
)

func TestContactIdsForEnrollPrefersAPIContactIds(t *testing.T) {
	t.Parallel()

	userId := "usr_1"
	fallback := map[string]string{
		userId: "contact-fallback",
	}

	require.Equal(t,
		"contact-api, contact-phone",
		contactIdsForEnroll(openapi.UserSvcEnroll{
			UserId:     &userId,
			ContactIds: []string{"contact-api", "contact-phone"},
		}, fallback),
	)

	require.Equal(t,
		"",
		contactIdsForEnroll(openapi.UserSvcEnroll{
			UserId:     &userId,
			ContactIds: []string{},
		}, fallback),
	)
}

func TestContactIdsForEnrollFallsBackWhenAPIContactIdsMissing(t *testing.T) {
	t.Parallel()

	userId := "usr_1"
	contactId := "contact-direct"
	fallback := map[string]string{
		userId: "contact-fallback",
	}

	require.Equal(t,
		contactId,
		contactIdsForEnroll(openapi.UserSvcEnroll{
			UserId:    &userId,
			ContactId: &contactId,
		}, fallback),
	)

	require.Equal(t,
		"contact-fallback",
		contactIdsForEnroll(openapi.UserSvcEnroll{
			UserId: &userId,
		}, fallback),
	)

	emptyContactId := ""
	require.Equal(t,
		"contact-fallback",
		contactIdsForEnroll(openapi.UserSvcEnroll{
			UserId:    &userId,
			ContactId: &emptyContactId,
		}, fallback),
	)
}

func TestUserIdsMissingContactIdsFromEnrolls(t *testing.T) {
	t.Parallel()

	missingUserId := "usr_missing"
	apiUserId := "usr_api"
	directUserId := "usr_direct"
	directContactId := "contact-direct"

	require.Equal(t,
		[]string{missingUserId},
		userIdsMissingContactIdsFromEnrolls([]openapi.UserSvcEnroll{
			{UserId: &missingUserId},
			{UserId: &missingUserId},
			{UserId: &apiUserId, ContactIds: []string{}},
			{UserId: &directUserId, ContactId: &directContactId},
		}),
	)
}
