package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFindOne(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "A1", Value: 10, CreatedAt: time.Now()}
	obj2 := TestObject{Name: "A2", Value: 10, CreatedAt: time.Now().Add(time.Minute)}
	obj3 := TestObject{Name: "A3", Value: 20, CreatedAt: time.Now().Add(2 * time.Minute)}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	res, found, err := store.Query(
		Equals(Field("Value"), 20),
	).FindOne()
	require.Equal(t, true, found)
	require.NoError(t, err)
	require.Equal(t, obj3.Name, res.(TestObject).Name)
}

func TestPointerFindOne(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "A1", Value: 10, CreatedAt: time.Now()}
	obj2 := &TestObject{Name: "A2", Value: 10, CreatedAt: time.Now().Add(time.Minute)}
	obj3 := &TestObject{Name: "A3", Value: 20, CreatedAt: time.Now().Add(2 * time.Minute)}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	res, found, err := store.Query(
		Equals(Field("Value"), 20),
	).FindOne()
	require.Equal(t, true, found)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, obj3.Name, res.(*TestObject).Name)
}

func TestCreate(t *testing.T, store DataStore) {
	obj1 := TestObject{
		Name:           "AliceCreate",
		NickNames:      []string{"A1", "A2"},
		Friends:        []Friend{{Name: "joe1"}, {Name: "joe2"}},
		FriendPointers: []*Friend{&Friend{Name: "jane1"}, &Friend{Name: "jane1"}},
		Value:          10,
		Age:            25,
	}

	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj1)
	require.Error(t, err)

	objs, err := store.Query().Find()
	require.NoError(t, err)

	require.Contains(t, objs, obj1)
}

func TestPointerCreate(t *testing.T, store DataStore) {
	obj1 := &TestObject{
		Name:           "AliceCreate",
		NickNames:      []string{"A1", "A2"},
		Friends:        []Friend{{Name: "joe1"}, {Name: "joe2"}},
		FriendPointers: []*Friend{&Friend{Name: "jane1"}, &Friend{Name: "jane1"}},
		Value:          10,
		Age:            25,
	}

	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj1)
	require.Error(t, err)
}

func TestUpsert(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}

	err := store.Upsert(obj1)
	require.NoError(t, err)

	res, err := store.Query().Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))

	err = store.Upsert(obj1)
	require.NoError(t, err)
}

func TestPointerUpsert(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}

	err := store.Upsert(obj1)
	require.NoError(t, err)

	res, err := store.Query().Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))

	err = store.Upsert(obj1)
	require.NoError(t, err)
}

func TestUpsertPartial(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}
	err := store.Upsert(obj1)
	require.NoError(t, err)

	obj2 := TestObject{Name: "AliceCreate", Value: 50, Age: 99, CreatedAt: time.Now()}
	err = store.Upsert(obj2, WithFields("Value"))
	require.NoError(t, err)

	res, found, err := store.Query(Equals(Field("Name"), "AliceCreate")).FindOne()
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 50, res.(TestObject).Value)
	require.Equal(t, 25, res.(TestObject).Age)
}

func TestPointerUpsertPartial(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}
	err := store.Upsert(obj1)
	require.NoError(t, err)

	obj2 := &TestObject{Name: "AliceCreate", Value: 50, Age: 99, CreatedAt: time.Now()}
	err = store.Upsert(obj2, WithFields("Value"))
	require.NoError(t, err)

	res, found, err := store.Query(Equals(Field("Name"), "AliceCreate")).FindOne()
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, 50, res.(*TestObject).Value)
	require.Equal(t, 25, res.(*TestObject).Age)
}

func TestUpsertManyPartial(t *testing.T, store DataStore) {
	err := store.UpsertMany([]Row{
		TestObject{Name: "test1", Value: 10, Age: 30, CreatedAt: time.Now()},
		TestObject{Name: "test2", Value: 20, Age: 40, CreatedAt: time.Now()},
	})
	require.NoError(t, err)

	err = store.UpsertMany([]Row{
		TestObject{Name: "test1", Value: 100, Age: 300, CreatedAt: time.Now()},
		TestObject{Name: "test2", Value: 200, Age: 400, CreatedAt: time.Now()},
	}, WithFields("Value"))
	require.NoError(t, err)

	res, err := store.Query().OrderBy(OrderByField("Name", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, 100, res[0].(TestObject).Value)
	require.Equal(t, 30, res[0].(TestObject).Age)
	require.Equal(t, 200, res[1].(TestObject).Value)
	require.Equal(t, 40, res[1].(TestObject).Age)
}

func TestPointerUpsertManyPartial(t *testing.T, store DataStore) {
	err := store.UpsertMany([]Row{
		&TestObject{Name: "test1", Value: 10, Age: 30, CreatedAt: time.Now()},
		&TestObject{Name: "test2", Value: 20, Age: 40, CreatedAt: time.Now()},
	})
	require.NoError(t, err)

	err = store.UpsertMany([]Row{
		&TestObject{Name: "test1", Value: 100, Age: 300, CreatedAt: time.Now()},
		&TestObject{Name: "test2", Value: 200, Age: 400, CreatedAt: time.Now()},
	}, WithFields("Value"))
	require.NoError(t, err)

	res, err := store.Query().OrderBy(OrderByField("Name", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 2, len(res))
	require.Equal(t, 100, res[0].(*TestObject).Value)
	require.Equal(t, 30, res[0].(*TestObject).Age)
	require.Equal(t, 200, res[1].(*TestObject).Value)
	require.Equal(t, 40, res[1].(*TestObject).Age)
}

func TestPatchMany(t *testing.T, store DataStore) {
	err := store.UpsertMany([]Row{
		TestObject{Name: "test1", Value: 10, Age: 30, CreatedAt: time.Now()},
		TestObject{Name: "test2", Value: 20, Age: 40, CreatedAt: time.Now()},
	})
	require.NoError(t, err)

	err = store.PatchMany([]Patch{
		{ID: "test1", Fields: map[string]any{"Value": 100}},
		{ID: "test2", Fields: map[string]any{"Value": 200}},
		{ID: "test3", Fields: map[string]any{"Value": 300}},
	})
	require.NoError(t, err)

	res, err := store.Query().OrderBy(OrderByField("Name", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 3, len(res))
	require.Equal(t, 100, res[0].(TestObject).Value)
	require.Equal(t, 30, res[0].(TestObject).Age)
	require.Equal(t, 200, res[1].(TestObject).Value)
	require.Equal(t, 40, res[1].(TestObject).Age)
	require.Equal(t, 300, res[2].(TestObject).Value)
}

func TestPointerPatchMany(t *testing.T, store DataStore) {
	err := store.UpsertMany([]Row{
		&TestObject{Name: "test1", Value: 10, Age: 30, CreatedAt: time.Now()},
		&TestObject{Name: "test2", Value: 20, Age: 40, CreatedAt: time.Now()},
	})
	require.NoError(t, err)

	err = store.PatchMany([]Patch{
		{ID: "test1", Fields: map[string]any{"Value": 100}},
		{ID: "test2", Fields: map[string]any{"Value": 200}},
		{ID: "test3", Fields: map[string]any{"Value": 300}},
	})
	require.NoError(t, err)

	res, err := store.Query().OrderBy(OrderByField("Name", false)).Find()
	require.NoError(t, err)
	require.Equal(t, 3, len(res))
	require.Equal(t, 100, res[0].(*TestObject).Value)
	require.Equal(t, 30, res[0].(*TestObject).Age)
	require.Equal(t, 200, res[1].(*TestObject).Value)
	require.Equal(t, 40, res[1].(*TestObject).Age)
	require.Equal(t, 300, res[2].(*TestObject).Value)
}

func TestUpdate(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}

	err := store.Create(obj1)
	require.NoError(t, err)

	obj1.Value = 50
	err = store.Query(Equals(Field("Name"), "AliceCreate")).Update(obj1)
	require.NoError(t, err)

	res, err := store.Query(
		Equals(Field("Value"), 50),
	).Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, res[0].(TestObject).Value, 50)
}

func TestPointerUpdate(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "AliceCreate", Value: 10, Age: 25, CreatedAt: time.Now()}

	err := store.Create(obj1)
	require.NoError(t, err)

	obj1.Value = 50

	err = store.Query(Equals(Field("Name"), "AliceCreate")).Update(obj1)
	require.NoError(t, err)

	res, err := store.Query(
		Equals(Field("Value"), 50),
	).Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, res[0].(*TestObject).Value, 50)
}
