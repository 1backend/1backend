package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T, store DataStore) {
	obj := TestObject{Name: "A", Amap: map[string]interface{}{
		"name": "A",
	}, AmapPointer: &map[string]interface{}{
		"namePointer": "AP",
	}}
	err := store.Create(obj)
	require.NoError(t, err)

	res, err := store.Query().Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "A", res[0].(TestObject).Amap["name"])
}

func TestMapPointer(t *testing.T, store DataStore) {
	obj := &TestObject{Name: "A", Amap: map[string]interface{}{
		"name": "A",
	}, AmapPointer: &map[string]interface{}{
		"namePointer": "AP",
	}}
	err := store.Create(obj)
	require.NoError(t, err)

	res, err := store.Query().Find()

	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "A", res[0].(*TestObject).Amap["name"])
}

func TestNamedTypeArray(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "Alice", NamedTypes: []NamedString{NamedStringOne, NamedStringTwo}}
	obj2 := TestObject{Name: "Bob", NamedTypes: []NamedString{NamedStringTwo, NamedStringThree}}
	obj3 := TestObject{Name: "Charlie", NamedTypes: []NamedString{NamedStringOne, NamedStringThree}}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(Equals(Field("NamedTypes"), NamedStringOne)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)
}

func TestPointerNamedTypeArray(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "Alice", NamedTypes: []NamedString{NamedStringOne, NamedStringTwo}}
	obj2 := &TestObject{Name: "Bob", NamedTypes: []NamedString{NamedStringTwo, NamedStringThree}}
	obj3 := &TestObject{Name: "Charlie", NamedTypes: []NamedString{NamedStringOne, NamedStringThree}}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(Equals(Field("NamedTypes"), NamedStringOne)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)
}
