/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Row interface {
	GetId() string
}

var (
	ErrEntryAlreadyExists = errors.New("entry already exists")
)

type DataStore interface {
	/*
	 * Create an object.
	 * Returns ErrEntryAlreadyExists if the object already exists.
	 */
	Create(obj Row) error

	/* Create many objects
	* Returns ErrEntryAlreadyExists if any of the objects are already in set,
	* and no object will be created.
	 */
	CreateMany(objs []Row) error

	/* Create or Update an object */
	Upsert(obj Row) error

	/* Create or Update many objects */
	UpsertMany(objs []Row) error

	Query(filters ...Filter) QueryBuilder

	BeginTransaction() (DataStore, error)
	Commit() error
	Rollback() error
	IsInTransaction() bool

	SetDebug(debug bool)

	Close() error
	Refresh() error
}

type QueryBuilder interface {
	OrderBy(options ...OrderBy) QueryBuilder
	Limit(limit int64) QueryBuilder
	After(value ...any) QueryBuilder

	Select(fields ...string) QueryBuilder
	Find() ([]Row, error)
	FindOne() (Row, bool, error)
	Count() (int64, error)

	// Update by query. Errors if no update happens
	Update(obj Row) error

	// Upsert tries to update by query, and if no update appened, calls create.
	Upsert(obj Row) error

	UpdateFields(fields map[string]interface{}) error
	Delete() error
}

type Op string

const (
	// OpEquals selects objects where the value of a field equals (=) the specified value in the query.
	// Example: `"fieldValue" = "value"`
	// SQL Equivalent: `SELECT * FROM table WHERE field = 'value';`
	// Elasticsearch Equivalent:
	// {
	//   "query": {
	//     "term": {
	//       "field": "value"
	//     }
	//   }
	// }
	OpEquals Op = "equals"

	// OpContainsSubstring selects all objects where the field's value contains a particular substring.
	// Example: `"fieldValue" CONTAINS_SUBSTRING "subString"`
	// SQL Equivalent: `SELECT * FROM table WHERE field LIKE '%subString%';`
	// Elasticsearch Equivalent:
	// {
	//   "query": {
	//     "wildcard": {
	//       "field": "*subString*"
	//     }
	//   }
	// }
	OpContainsSubstring Op = "containsSubstring"

	// OpStartsWith selects all objects where the field's value starts with a particular string.
	// Example: `"fieldValue" STARTS WITH "prefix"`
	// SQL Equivalent: `SELECT * FROM table WHERE field LIKE 'prefix%';`
	// Elasticsearch Equivalent:
	// {
	//   "query": {
	//     "prefix": {
	//       "field": "prefix"
	//     }
	//   }
	// }
	OpStartsWith Op = "startsWith"

	// OpIntersects selects objects where the slice value of a field intersects with the slice value in the query.
	// Example: `["fieldValue2", "fieldValue3"] INTERSECTS ["value1", "value2"]`
	// SQL Equivalent: `SELECT * FROM table WHERE field && ARRAY['value1', 'value2'];` (PostgreSQL syntax)
	// Elasticsearch Equivalent:
	// {
	//   "query": {
	//     "terms_set": {
	//       "field": {
	//         "terms": ["value1", "value2"],
	//         "minimum_should_match_script": {
	//           "source": "1"
	//         }
	//       }
	//     }
	//   }
	// }
	OpIntersects Op = "intersects"

	// OpIsInList selects objects where the value of a field is one of the specified values in a list.
	// Example: `"fieldValue" IS_IN_LIST ["value1", "value2", "value3"]`
	// SQL Equivalent: `SELECT * FROM table WHERE field IN ('value1', 'value2', 'value3');`
	// Elasticsearch Equivalent:
	// {
	//   "query": {
	//     "terms": {
	//       "field": ["value1", "value2", "value3"]
	//     }
	//   }
	// }
	OpIsInList Op = "isInList"
)

type Filter struct {
	Fields []string `json:"fields,omitempty"`

	// JSONValues is a JSON marshalled array of values.
	// It's JSON marhalled due to the limitations of the
	// Swaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.
	JSONValues string `json:"jsonValues,omitempty"`

	Op Op `json:"op"`
}

func (c Filter) FieldIs(fieldName string) bool {
	for _, field := range c.Fields {
		if fieldName == field {
			return true
		}
	}
	return false
}

// Query as a type is not used in the DataStore interface but mostly to accept
// a DataStore query through a HTTP API
type Query struct {
	// Filters are filtering options of a query. It is advised to use
	// It's advised to use helper functions in your respective client library such as filter constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
	Filters []Filter `json:"filters,omitempty"`

	// JSONAfter is used for cursor-based pagination, which is more
	// effective in scalable and distributed environments compared
	// to offset-based pagination.
	//
	// JSONAfter is a JSON-encoded string due to limitations in Swaggo (e.g., []interface{} gets converted to []map[string]interface{}).
	// There is no way to specify a type that results in an any/interface{} type in the `go -> openapi -> go` generation process.
	// As a result, JSONAfter is a JSON-marshalled string representing an array, e.g., `[42]`.
	JSONAfter string `json:"jsonAfter,omitempty"`

	// Limit the number of records in the result set.
	Limit int64 `json:"limit,omitempty"`

	// OrderBys order the result set.
	OrderBys []OrderBy `json:"orderBys,omitempty"`

	// Count true means return the count of the dataset filtered by Filters
	// without after or limit.
	Count bool `json:"count,omitempty"`
}

func (q *Query) HasFieldFilter(fieldName string) bool {
	for _, v := range q.Filters {
		if v.FieldIs(fieldName) {
			return true
		}
	}

	return false
}

type SortingType string

const (
	SortingTypeDefault SortingType = ""
	SortingTypeNumeric SortingType = "numeric"
	SortingTypeText    SortingType = "text"
	SortingTypeDate    SortingType = "date"
)

type OrderBy struct {
	// The field by which to order the results
	Field string `json:"field,omitempty"`

	// Desc indicates whether the sorting should be in descending order.
	Desc bool `json:"desc,omitempty"`

	// Randomize indicates that the results should be randomized instead of ordered by the `field` and `desc` criteria
	Randomize bool `json:"randomize,omitempty"`

	// Defines the type of sorting to apply (numeric, text, date, etc.)
	SortingType SortingType `json:"sortingType,omitempty"`
}

// random order. not advised for large datasets due to its slow speed
// in a distributed setting
func OrderByRandom() OrderBy {
	return OrderBy{
		Randomize: true,
	}
}

func OrderByField(field string, desc bool) OrderBy {
	return OrderBy{
		Field: field,
		Desc:  desc,
	}
}

type AllMatch struct {
}

func marshal(value any) string {
	jsonBytes, _ := json.Marshal(value)

	return string(jsonBytes)
}

func Equals(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		JSONValues: marshal([]any{value}),
		Op:         OpEquals,
	}
}

func Intersects(fields []string, values []any) Filter {
	return Filter{
		Fields:     fields,
		JSONValues: marshal(values),
		Op:         OpIntersects,
	}
}

func StartsWith(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		JSONValues: marshal([]any{value}),
		Op:         OpStartsWith,
	}
}

func ContainsSubstring(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		JSONValues: marshal([]any{value}),
		Op:         OpContainsSubstring,
	}
}

// See `OpIsInList` for more information.
func IsInList(fields []string, values ...any) Filter {
	return Filter{
		Fields:     fields,
		JSONValues: marshal(values),
		Op:         OpIsInList,
	}
}

func Id(id any) Filter {
	return Filter{
		Fields:     []string{"id"},
		JSONValues: marshal([]any{id}),
		Op:         OpEquals,
	}
}

func Field(fieldName string) []string {
	return []string{fieldName}
}

func Fields(fieldNames ...string) []string {
	return fieldNames
}

func AnyField() []string {
	return nil
}

var dateFormats = []string{
	time.RFC3339,
	time.RFC1123,
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02",
	"02-Jan-2006 15:04:05",
	"02-Jan-2006 15:04",
	"02-Jan-2006",
	"02/01/2006 15:04:05",
	"02/01/2006 15:04",
	"02/01/2006",
	"01/02/2006 15:04:05",
	"01/02/2006 15:04",
	"01/02/2006",
	"2006-1-2 15:04:05",
	"2006-1-2 15:04",
	"2006-1-2",
	"1/2/2006 15:04:05",
	"1/2/2006 15:04",
	"1/2/2006",
}

func ParseAnyDate(input string) (time.Time, error) {
	var t time.Time
	var err error
	for _, format := range dateFormats {
		t, err = time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse date: %v", input)
}
