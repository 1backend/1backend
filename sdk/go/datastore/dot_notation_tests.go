package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDotNotation(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "Alice", Value: 10, Age: 25,
		Friend:        Friend{Name: "AliceFriend", Age: 26},
		FriendPointer: &Friend{Name: "AliceFriendP", Age: 27},
	}
	obj2 := TestObject{Name: "Bob", Value: 20, Age: 30,
		Friend:        Friend{Name: "BobFriend", Age: 31},
		FriendPointer: &Friend{Name: "BobFriendP", Age: 32},
	}
	obj3 := TestObject{Name: "Charlie", Value: 30, Age: 35,
		Friend:        Friend{Name: "CharlieFriend", Age: 36},
		FriendPointer: &Friend{Name: "CharlieFriendP", Age: 37},
	}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(IsInList(Field("Friend.Name"), "AliceFriend", "BobFriend")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	results, err = store.Query(IsInList(Field("friendPointer.name"), "AliceFriendP", "BobFriendP")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(IsInList(Field("Friend.Age"), 26, 36)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(Equals(Field("Friend.Age"), []any{})).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test Ordering
	results, err = store.Query().OrderBy(OrderByField("Friend.Age", false)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Alice", results[0].(TestObject).Name)
	require.Equal(t, "Bob", results[1].(TestObject).Name)
	require.Equal(t, "Charlie", results[2].(TestObject).Name)

	results, err = store.Query().OrderBy(OrderByField("Friend.Age", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Charlie", results[0].(TestObject).Name)
	require.Equal(t, "Bob", results[1].(TestObject).Name)
	require.Equal(t, "Alice", results[2].(TestObject).Name)

	// Test IN clause with one element slice
	results, err = store.Query(IsInList(Field("FriendPointer.Age"), 32)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(Equals(Field("Friend.Name"), "AliceFriend")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Friend.Name"), "BobFriendP")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Friend.Name"), "CharlieFriend")).Delete()
	require.NoError(t, err)
}

func TestPointerDotNotation(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "Alice", Value: 10, Age: 25,
		Friend:        Friend{Name: "AliceFriend", Age: 26},
		FriendPointer: &Friend{Name: "AliceFriendP", Age: 27},
	}
	obj2 := &TestObject{Name: "Bob", Value: 20, Age: 30,
		Friend:        Friend{Name: "BobFriend", Age: 31},
		FriendPointer: &Friend{Name: "BobFriendP", Age: 32},
	}
	obj3 := &TestObject{Name: "Charlie", Value: 30, Age: 35,
		Friend:        Friend{Name: "CharlieFriend", Age: 36},
		FriendPointer: &Friend{Name: "CharlieFriendP", Age: 37},
	}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(IsInList(Field("Friend.Name"), "AliceFriend", "BobFriend")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	results, err = store.Query(IsInList(Field("friendPointer.name"), "AliceFriendP", "BobFriendP")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(IsInList(Field("Friend.Age"), 26, 36)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(IsInList(Field("Friend.Age"))).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test Ordering
	results, err = store.Query().OrderBy(OrderByField("Friend.Age", false)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Alice", results[0].(*TestObject).Name)
	require.Equal(t, "Bob", results[1].(*TestObject).Name)
	require.Equal(t, "Charlie", results[2].(*TestObject).Name)

	results, err = store.Query().OrderBy(OrderByField("Friend.Age", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "Charlie", results[0].(*TestObject).Name)
	require.Equal(t, "Bob", results[1].(*TestObject).Name)
	require.Equal(t, "Alice", results[2].(*TestObject).Name)

	// Test IN clause with one element slice
	results, err = store.Query(IsInList(Field("FriendPointer.Age"), 32)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(Equals(Field("Friend.Name"), "AliceFriend")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Friend.Name"), "BobFriendP")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Friend.Name"), "CharlieFriend")).Delete()
	require.NoError(t, err)
}
