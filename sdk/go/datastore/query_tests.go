package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "HelloThere"}
	obj2 := TestObject{Name: "HiWhatsUp"}

	err := store.Create(obj1)
	require.NoError(t, err)

	err = store.Create(obj2)
	require.NoError(t, err)

	res, err := store.Query(ContainsSubstring(Field("Name"), "lo")).Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "HelloThere", res[0].(TestObject).Name)

	res, err = store.Query(ContainsSubstring(Field("Name"), "What")).Find()
	require.NoError(t, err)
	require.Equal(t, 1, len(res))
	require.Equal(t, "HiWhatsUp", res[0].(TestObject).Name)
}

func TestQuery(t *testing.T, store DataStore) {
	objs := []Row{
		TestObject{Name: "queryTest1", Value: 10},
		TestObject{Name: "queryTest2", Value: 20},
		TestObject{Name: "queryTest3", Value: 30},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Value"), 20)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	results, err = store.Query().OrderBy(OrderByField("Value", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	results, err = store.Query().OrderBy(OrderByField("Name", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	// order is nondeterministic when no OrderBy is supplied
	// results, err = store.Query().Limit(2).Offset(1).Find()
	// require.NoError(t, err)
	// require.Len(t, results, 2)
	// require.Equal(t, objs[1], results[0])
	// require.Equal(t, objs[2], results[1])

	count, err := store.Query(Equals(Field("Value"), 10)).Count()
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	err = store.Query(Equals(Field("Value"), 10)).UpdateFields(map[string]interface{}{
		"Value": 100,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Value"), 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 100, results[0].(TestObject).Value)

	err = store.Query(Equals(Field("Value"), 100)).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Value"), 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func TestPointerQuery(t *testing.T, store DataStore) {
	objs := []Row{
		&TestObject{Name: "queryTest1", Value: 10},
		&TestObject{Name: "queryTest2", Value: 20},
		&TestObject{Name: "queryTest3", Value: 30},
	}

	err := store.CreateMany(objs)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Value"), 20)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, objs[1], results[0])

	results, err = store.Query().OrderBy(OrderByField("Value", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	results, err = store.Query().OrderBy(OrderByField("Name", true)).Find()
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, objs[2], results[0])
	require.Equal(t, objs[1], results[1])
	require.Equal(t, objs[0], results[2])

	// order is nondeterministic when no OrderBy is supplied
	// results, err = store.Query().Limit(2).Offset(1).Find()
	// require.NoError(t, err)
	// require.Len(t, results, 2)
	// require.Equal(t, objs[1], results[0])
	// require.Equal(t, objs[2], results[1])

	count, err := store.Query(Equals(Field("Value"), 10)).Count()
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	err = store.Query(Equals(Field("Value"), 10)).UpdateFields(map[string]interface{}{
		"Value": 100,
	})
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Value"), 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, 100, results[0].(*TestObject).Value)

	err = store.Query(Equals(Field("Value"), 100)).Delete()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Value"), 100)).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)
}

func TestReverseInClause(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "Alice", NickNames: []string{"A1", "A2"}}
	obj2 := TestObject{Name: "Bob", NickNames: []string{"B1"}}
	obj3 := TestObject{Name: "Charlie"}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	results, err := store.Query(
		Equals(Field("NickNames"), "A1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		Equals(Field("NickNames"), "A2"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		Equals(Field("NickNames"), "B1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)
}

func TestPointerReverseInClause(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "Alice", NickNames: []string{"A1", "A2"}}
	obj2 := &TestObject{Name: "Bob", NickNames: []string{"B1"}}
	obj3 := &TestObject{Name: "Charlie"}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	results, err := store.Query(
		Equals(Field("NickNames"), "A1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		Equals(Field("NickNames"), "A2"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj1)

	results, err = store.Query(
		Equals(Field("NickNames"), "B1"),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)
}

func TestInClause(t *testing.T, store DataStore) {
	obj1 := TestObject{Name: "Alice", Value: 10, Age: 25, NickNames: []string{"A1", "A2"}, NamedType: NamedStringOne}
	obj2 := TestObject{Name: "Bob", Value: 20, Age: 30, NickNames: []string{"B1", "B2"}, NamedType: NamedStringTwo}
	obj3 := TestObject{Name: "Charlie", Value: 30, Age: 35, NickNames: []string{"C1", "C2"}, NamedType: NamedStringThree}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(IsInList(Field("Name"), "Alice", "Bob")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(IsInList(Field("Value"), 10, 30)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(IsInList(Field("Age"))).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test IN clause with one element slice
	results, err = store.Query(IsInList(Field("Age"), 30)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	results, err = store.Query(IsInList(Field("NamedType"), NamedStringTwo)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Test IN clause array intersection
	results, err = store.Query(Intersects(Field("NickNames"), []any{"A1", "B2"})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(Equals(Field("Name"), "Alice")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Name"), "Bob")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Name"), "Charlie")).Delete()
	require.NoError(t, err)
}

func TestPointerInClause(t *testing.T, store DataStore) {
	obj1 := &TestObject{Name: "Alice", Value: 10, Age: 25, NickNames: []string{"A1", "A2"}, NamedType: NamedStringOne}
	obj2 := &TestObject{Name: "Bob", Value: 20, Age: 30, NickNames: []string{"B1", "B2"}, NamedType: NamedStringTwo}
	obj3 := &TestObject{Name: "Charlie", Value: 30, Age: 35, NickNames: []string{"C1", "C2"}, NamedType: NamedStringThree}

	err := store.Create(obj1)
	require.NoError(t, err)
	err = store.Create(obj2)
	require.NoError(t, err)
	err = store.Create(obj3)
	require.NoError(t, err)

	// Test IN clause with string slice
	results, err := store.Query(IsInList(Field("Name"), "Alice", "Bob")).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Test IN clause with int slice
	results, err = store.Query(IsInList(Field("Value"), 10, 30)).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj3)

	// Test IN clause with empty slice (should return no results)
	results, err = store.Query(IsInList(Field("Age"))).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	// Test IN clause with one element slice
	results, err = store.Query(IsInList(Field("Age"), 30)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	results, err = store.Query(IsInList(Field("NamedType"), NamedStringTwo)).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Contains(t, results, obj2)

	// Test IN clause array intersection
	results, err = store.Query(Intersects(Field("NickNames"), []any{"A1", "B2"})).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, obj1)
	require.Contains(t, results, obj2)

	// Clean up
	err = store.Query(Equals(Field("Name"), "Alice")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Name"), "Bob")).Delete()
	require.NoError(t, err)
	err = store.Query(Equals(Field("Name"), "Charlie")).Delete()
	require.NoError(t, err)
}
