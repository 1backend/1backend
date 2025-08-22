package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOr(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "FirstObject"}
	obj2 := TestObject{Name: "SecondObject"}
	obj3 := TestObject{Name: "ThirdObject"}

	// Create objects in the store
	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj2)
	require.NoError(t, err)

	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Or(
			Equals([]string{"name"}, "SecondObject"),
			Equals([]string{"name"}, "ThirdObject"),
		),
	).OrderBy(
		OrderByField("name", false),
	).
		Find()

	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "SecondObject", res[0].GetId())
	require.Equal(t, "ThirdObject", res[1].GetId())
}

func TestPointerOr(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "FirstObject"}
	obj2 := &TestObject{Name: "SecondObject"}
	obj3 := &TestObject{Name: "ThirdObject"}

	// Create objects in the store
	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj2)
	require.NoError(t, err)

	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Or(
			Equals([]string{"name"}, "SecondObject"),
			Equals([]string{"name"}, "ThirdObject"),
		),
	).OrderBy(
		OrderByField("name", false),
	).Find()

	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "SecondObject", res[0].GetId())
	require.Equal(t, "ThirdObject", res[1].GetId())
}

func TestComplexOr(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "FirstObject", Age: 10}
	obj2 := TestObject{Name: "SecondObject", Age: 20}
	obj3 := TestObject{Name: "ThirdObject", Age: 30}

	// Create objects in the store
	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj2)
	require.NoError(t, err)

	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Or(
			Equals([]string{"name"}, "SecondObject"),
			Equals([]string{"name"}, "ThirdObject"),
		),
		Or(
			Equals([]string{"age"}, 10),
			Equals([]string{"age"}, 30),
		),
	).OrderBy().
		Find()

	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "ThirdObject", res[0].GetId())
}

func TestPointerComplexOr(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "FirstObject", Age: 10}
	obj2 := TestObject{Name: "SecondObject", Age: 20}
	obj3 := TestObject{Name: "ThirdObject", Age: 30}

	// Create objects in the store
	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj2)
	require.NoError(t, err)

	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Or(
			Equals([]string{"name"}, "SecondObject"),
			Equals([]string{"name"}, "ThirdObject"),
		),
		Or(
			Equals([]string{"age"}, 10),
			Equals([]string{"age"}, 30),
		),
	).OrderBy().Find()

	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "ThirdObject", res[0].GetId())
}
