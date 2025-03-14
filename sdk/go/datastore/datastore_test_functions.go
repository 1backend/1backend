/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package datastore

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NamedString string

const (
	NamedStringOne   = "one"
	NamedStringTwo   = "two"
	NamedStringThree = "three"
)

type TestObject struct {
	Name              string                  `json:"name"`
	Value             int                     `json:"value"`
	Age               int                     `json:"age"`
	NickNames         []string                `json:"nickNames"`
	Friend            Friend                  `json:"friend"`
	Friends           []Friend                `json:"friends"`
	FriendPointers    []*Friend               `json:"friendPointers"`
	Amap              map[string]interface{}  `json:"amap"`
	AmapPointer       *map[string]interface{} `json:"amapPointer"`
	AmapString        map[string]string       `json:"amapString"`
	AmapStringPointer *map[string]string      `json:"amapStringPointer"`
	FriendPointer     *Friend                 `json:"friendPointer"`
	CreatedAt         time.Time               `json:"createdAt"`
	NamedType         NamedString             `json:"namedType"`
}

func (t TestObject) GetId() string {
	return t.Name
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

func TestPagination(t *testing.T, store DataStore) {
	for i := 1; i <= 10; i++ {
		obj := TestObject{Name: fmt.Sprintf("PaginationTest%d", i), Value: i}
		err := store.Create(obj)
		require.NoError(t, err)
	}

	results, err := store.Query().OrderBy(OrderByField("Value", true)).Limit(5).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest10", results[0].(TestObject).Name)
	require.Equal(t, 10, results[0].(TestObject).Value)

	lastValue := results[len(results)-1].(TestObject).Value
	results, err = store.Query().OrderBy(OrderByField("Value", true)).Limit(5).After(lastValue).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest5", results[0].(TestObject).Name)
	require.Equal(t, 5, results[0].(TestObject).Value)
}

func TestPointerPagination(t *testing.T, store DataStore) {
	for i := 1; i <= 10; i++ {
		obj := &TestObject{Name: fmt.Sprintf("PaginationTest%d", i), Value: i}
		err := store.Create(obj)
		require.NoError(t, err)
	}

	results, err := store.Query().OrderBy(OrderByField("Value", true)).Limit(5).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest10", results[0].(*TestObject).Name)
	require.Equal(t, 10, results[0].(*TestObject).Value)

	lastValue := results[len(results)-1].(*TestObject).Value
	results, err = store.Query().OrderBy(OrderByField("Value", true)).Limit(5).After(lastValue).Find()
	require.NoError(t, err)
	require.Len(t, results, 5)
	require.Equal(t, "PaginationTest5", results[0].(*TestObject).Name)
	require.Equal(t, 5, results[0].(*TestObject).Value)
}

func TestCreatedAt(t *testing.T, store DataStore) {
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

func TestPointerCreatedAt(t *testing.T, store DataStore) {
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

func TestTransactions(t *testing.T, store DataStore) {
	tx, err := store.BeginTransaction()
	require.NoError(t, err)

	obj := TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	err = tx.Commit()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]
	require.Equal(t, obj, readObj)
}

func TestPointerTransactions(t *testing.T, store DataStore) {
	tx, err := store.BeginTransaction()
	require.NoError(t, err)

	obj := &TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	err = tx.Commit()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]
	require.Equal(t, obj, readObj)
}
