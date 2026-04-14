package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestScalarPointers(t *testing.T, store DataStore) {
	userID := "usr-iNQeXfVVot"
	failures := 3
	active := true
	reviewedAt := time.Now().UTC().Round(time.Microsecond)

	obj := ScalarPointerTestObject{
		Name:          "scalar-pointers",
		StringPointer: &userID,
		IntPointer:    &failures,
		BoolPointer:   &active,
		TimePointer:   &reviewedAt,
		FriendPointer: &Friend{Name: "reviewer", Age: 42},
	}

	err := store.Upsert(obj)
	require.NoError(t, err)

	res, found, err := store.Query(Equals(Field("Name"), obj.Name)).FindOne()
	require.NoError(t, err)
	require.True(t, found)

	stored := res.(ScalarPointerTestObject)
	require.NotNil(t, stored.StringPointer)
	require.Equal(t, userID, *stored.StringPointer)
	require.NotNil(t, stored.IntPointer)
	require.Equal(t, failures, *stored.IntPointer)
	require.NotNil(t, stored.BoolPointer)
	require.Equal(t, active, *stored.BoolPointer)
	require.NotNil(t, stored.TimePointer)
	require.Equal(t, reviewedAt, *stored.TimePointer)
	require.NotNil(t, stored.FriendPointer)
	require.Equal(t, obj.FriendPointer.Name, stored.FriendPointer.Name)
	require.Equal(t, obj.FriendPointer.Age, stored.FriendPointer.Age)
}

func TestPointerScalarPointers(t *testing.T, store DataStore) {
	userID := "usr-iNQeXfVVot"
	failures := 3
	active := true
	reviewedAt := time.Now().UTC().Round(time.Microsecond)

	obj := &ScalarPointerTestObject{
		Name:          "pointer-scalar-pointers",
		StringPointer: &userID,
		IntPointer:    &failures,
		BoolPointer:   &active,
		TimePointer:   &reviewedAt,
		FriendPointer: &Friend{Name: "reviewer", Age: 42},
	}

	err := store.Upsert(obj)
	require.NoError(t, err)

	res, found, err := store.Query(Equals(Field("Name"), obj.Name)).FindOne()
	require.NoError(t, err)
	require.True(t, found)

	stored := res.(*ScalarPointerTestObject)
	require.NotNil(t, stored.StringPointer)
	require.Equal(t, userID, *stored.StringPointer)
	require.NotNil(t, stored.IntPointer)
	require.Equal(t, failures, *stored.IntPointer)
	require.NotNil(t, stored.BoolPointer)
	require.Equal(t, active, *stored.BoolPointer)
	require.NotNil(t, stored.TimePointer)
	require.Equal(t, reviewedAt, *stored.TimePointer)
	require.NotNil(t, stored.FriendPointer)
	require.Equal(t, obj.FriendPointer.Name, stored.FriendPointer.Name)
	require.Equal(t, obj.FriendPointer.Age, stored.FriendPointer.Age)
}
