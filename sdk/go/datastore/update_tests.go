package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateReadUpdateDelete(t *testing.T, store DataStore) {
	obj := TestObject{Name: "test", Value: 10}
	err := store.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]

	require.Equal(t, obj, readObj)

	obj.Value = 20
	err = store.Query(Equals(Field("Name"), "test")).UpdateFields(map[string]interface{}{
		"Value": obj.Value,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj = results[0]

	require.Equal(t, obj, readObj)

	err = store.Query(Equals(Field("Name"), "test")).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func TestPointerCreateReadUpdateDelete(t *testing.T, store DataStore) {
	obj := &TestObject{Name: "test", Value: 10}
	err := store.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]

	require.Equal(t, obj, readObj)

	obj.Value = 20
	err = store.Query(Equals(Field("Name"), "test")).UpdateFields(map[string]interface{}{
		"Value": obj.Value,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj = results[0]

	require.Equal(t, obj, readObj)

	err = store.Query(Equals(Field("Name"), "test")).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func TestCreateManyUpdateDelete(t *testing.T, store DataStore) {
	objs := []Row{
		TestObject{Name: "test1", Age: 100, Value: 10},
		TestObject{Name: "test2", Age: 200, Value: 20},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[0], results[0])

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	err = store.Query(Equals(Field("Name"), "test1")).UpdateFields(map[string]interface{}{
		"Value": 30,
	})
	require.NoError(t, err)

	err = store.Query(Equals(Field("Name"), "test2")).UpdateFields(map[string]interface{}{
		"value": 40,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 30, results[0].(TestObject).Value)
	require.Equal(t, 100, results[0].(TestObject).Age)

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 40, results[0].(TestObject).Value)
	require.Equal(t, 200, results[0].(TestObject).Age)

	err = store.Query(Equals(Field("Name"), "test1")).Delete()
	require.NoError(t, err)

	err = store.Query(Equals(Field("Name"), "test2")).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func TestPointerCreateManyUpdateDelete(t *testing.T, store DataStore) {
	objs := []Row{
		&TestObject{Name: "test1", Age: 100, Value: 10},
		&TestObject{Name: "test2", Age: 200, Value: 20},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[0], results[0])

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	err = store.Query(Equals(Field("Name"), "test1")).UpdateFields(map[string]interface{}{
		"Value": 30,
	})
	require.NoError(t, err)

	err = store.Query(Equals(Field("Name"), "test2")).UpdateFields(map[string]interface{}{
		"value": 40,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 30, results[0].(*TestObject).Value)
	require.Equal(t, 100, results[0].(*TestObject).Age)

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 40, results[0].(*TestObject).Value)
	require.Equal(t, 200, results[0].(*TestObject).Age)

	err = store.Query(Equals(Field("Name"), "test1")).Delete()
	require.NoError(t, err)

	err = store.Query(Equals(Field("Name"), "test2")).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test1")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	results, err = store.Query(Equals(Field("Name"), "test2")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}
