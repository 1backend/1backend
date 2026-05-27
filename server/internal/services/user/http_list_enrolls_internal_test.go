package userservice

import (
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/stretchr/testify/require"
)

func TestHydrateEnrollContactIdsIncludesUserContacts(t *testing.T) {
	t.Parallel()

	contactStore, err := localstore.NewLocalStore(&user.Contact{}, "")
	require.NoError(t, err)

	now := time.Now()
	require.NoError(t, contactStore.Upsert(&user.Contact{
		Id:        "contact-user-email",
		CreatedAt: now,
		UpdatedAt: now,
		UserId:    "usr_1",
		Platform:  "email",
		Handle:    "contact-user-email",
	}))
	require.NoError(t, contactStore.Upsert(&user.Contact{
		Id:        "user-phone",
		CreatedAt: now,
		UpdatedAt: now,
		UserId:    "usr_1",
		Platform:  "phone",
		Handle:    "user-phone",
	}))

	service := &UserService{
		contactStore: contactStore,
	}
	enrolls := []user.Enroll{
		{
			Id:        "enr_1",
			ContactId: "contact-explicit",
			UserId:    "usr_1",
		},
	}

	require.NoError(t, service.hydrateEnrollContactIds(enrolls, map[string]struct{}{
		"usr_1": {},
	}))
	require.ElementsMatch(t,
		[]string{"contact-explicit", "contact-user-email", "user-phone"},
		enrolls[0].ContactIds,
	)
}

func TestHydrateEnrollContactIdsDoesNotIncludeHiddenUserContacts(t *testing.T) {
	t.Parallel()

	contactStore, err := localstore.NewLocalStore(&user.Contact{}, "")
	require.NoError(t, err)

	now := time.Now()
	require.NoError(t, contactStore.Upsert(&user.Contact{
		Id:        "contact-user-email",
		CreatedAt: now,
		UpdatedAt: now,
		UserId:    "usr_1",
		Platform:  "email",
		Handle:    "contact-user-email",
	}))

	service := &UserService{
		contactStore: contactStore,
	}
	enrolls := []user.Enroll{
		{
			Id:        "enr_1",
			ContactId: "contact-explicit",
			UserId:    "usr_1",
		},
	}

	require.NoError(t, service.hydrateEnrollContactIds(enrolls, nil))
	require.Equal(t, []string{"contact-explicit"}, enrolls[0].ContactIds)
}

func TestHydrateEnrollContactIdsSetsEmptySlice(t *testing.T) {
	t.Parallel()

	contactStore, err := localstore.NewLocalStore(&user.Contact{}, "")
	require.NoError(t, err)

	service := &UserService{
		contactStore: contactStore,
	}
	enrolls := []user.Enroll{
		{
			Id:     "enr_1",
			UserId: "usr_1",
		},
	}

	require.NoError(t, service.hydrateEnrollContactIds(enrolls, nil))
	require.NotNil(t, enrolls[0].ContactIds)
	require.Empty(t, enrolls[0].ContactIds)
}
