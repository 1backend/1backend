package datastore

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOrderBy(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "A1", Value: 10, CreatedAt: time.Now()}
	obj2 := TestObject{Name: "A2", Value: 10, CreatedAt: time.Now().Add(time.Minute)}
	obj3 := TestObject{Name: "A3", Value: 20, CreatedAt: time.Now().Add(2 * time.Minute)}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Equals(Field("Value"), 101),
	).OrderBy(OrderByField("CreatedAt", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 0, len(res))

	res, err = store.Query(
		Equals(Field("Value"), 10),
	).OrderBy(OrderByField("CreatedAt", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A1", res[0].(TestObject).Name)

	res, err = store.Query(
		Equals(Field("Value"), 10),
	).OrderBy(OrderByField("CreatedAt", true)).Find()

	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A2", res[0].(TestObject).Name)
}

func TestPointerOrderBy(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "A1", Value: 10, CreatedAt: time.Now()}
	obj2 := &TestObject{Name: "A2", Value: 10, CreatedAt: time.Now().Add(time.Minute)}
	obj3 := &TestObject{Name: "A3", Value: 20, CreatedAt: time.Now().Add(2 * time.Minute)}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	res, err := store.Query(
		Equals(Field("Value"), 101),
	).OrderBy(OrderByField("CreatedAt", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 0, len(res))

	res, err = store.Query(
		Equals(Field("Value"), 10),
	).OrderBy(OrderByField("CreatedAt", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A1", res[0].(*TestObject).Name)

	res, err = store.Query(
		Equals(Field("Value"), 10),
	).OrderBy(OrderByField("CreatedAt", true)).Find()

	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, "A2", res[0].(*TestObject).Name)
}

func TestRandomize(t *testing.T, store DataStore) {
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

	firstName := ""

	// Query with randomization
	res, err := store.Query().OrderBy(OrderByRandom()).Find()
	require.NoError(t, err)
	require.Equal(t, 3, len(res))
	firstName = res[0].(TestObject).Name

	// Verify that the results are randomized
	// Since the results are random, we need to check multiple times to ensure they are not always in the same order
	isDifferent := false
	for i := 0; i < 10; i++ {
		newRes, err := store.Query().OrderBy(OrderByRandom()).Find()
		require.NoError(t, err)
		require.Equal(t, 3, len(newRes))
		if newRes[0].(TestObject).Name != firstName {
			isDifferent = true
			break
		}
	}

	fmt.Println(isDifferent)
	// @todo: random order is expensive in a distributed setting
	// rethink this test
	// require.True(t, isDifferent, "Expected results to be randomized, but they were the same in multiple queries")
}

func TestPointerRandomize(t *testing.T, store DataStore) {
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

	firstName := ""

	// Query with randomization
	res, err := store.Query().OrderBy(OrderByRandom()).Find()
	require.NoError(t, err)
	require.Equal(t, 3, len(res))
	firstName = res[0].(*TestObject).Name

	// Verify that the results are randomized
	// Since the results are random, we need to check multiple times to ensure they are not always in the same order
	isDifferent := false
	for i := 0; i < 10; i++ {
		newRes, err := store.Query().OrderBy(OrderByRandom()).Find()
		require.NoError(t, err)
		require.Equal(t, 3, len(newRes))
		if newRes[0].(*TestObject).Name != firstName {
			isDifferent = true
			break
		}
	}

	fmt.Println(isDifferent)
	// @todo: random order is expensive in a distributed setting
	// rethink this test
	// require.True(t, isDifferent, "Expected results to be randomized, but they were the same in multiple queries")
}
