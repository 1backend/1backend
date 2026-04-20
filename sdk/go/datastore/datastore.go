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

	// Create or update an object. By default this replaces the full stored object on update.
	Upsert(obj Row, opts ...UpsertOption) error

	// Create or update many objects. By default this replaces full stored object(s) on update.
	UpsertMany(objs []Row, opts ...UpsertOption) error

	// Create or partially update one object by id from field/value pairs.
	Patch(id string, fields map[string]any) error

	// Create or partially update many objects by id from field/value pairs.
	PatchMany(updates []Patch) error

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

type UpsertOptions struct {
	Fields []string
}

type UpsertOption func(*UpsertOptions)

func WithFields(fields ...string) UpsertOption {
	return func(o *UpsertOptions) {
		o.Fields = append(o.Fields, fields...)
	}
}

func ParseUpsertOptions(opts ...UpsertOption) UpsertOptions {
	options := UpsertOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

type Patch struct {
	ID     string
	Fields map[string]any
}

type Indexable interface {
	Indexes() []Index
}

type AutoIndexStatsProvider interface {
	AutoIndexStats() AutoIndexStats
}

type AutoIndexStatus string

const (
	AutoIndexStatusObserved   AutoIndexStatus = "observed"
	AutoIndexStatusPending    AutoIndexStatus = "pending"
	AutoIndexStatusCreated    AutoIndexStatus = "created"
	AutoIndexStatusFailed     AutoIndexStatus = "failed"
	AutoIndexStatusDiscovered AutoIndexStatus = "discovered"
	AutoIndexStatusManual     AutoIndexStatus = "manual"
)

type AutoIndexKind string

const (
	AutoIndexKindAuto   AutoIndexKind = "auto"
	AutoIndexKindManual AutoIndexKind = "manual"
)

type AutoIndexStats struct {
	Supported bool `json:"supported"`
	Backend   string

	Shapes  []AutoIndexShapeStats `json:"shapes,omitempty"`
	Indexes []AutoIndexEntry      `json:"indexes,omitempty"`
}

type AutoIndexShapeStats struct {
	Fingerprint string    `json:"fingerprint"`
	Hits        int       `json:"hits"`
	LastSeen    time.Time `json:"lastSeen,omitempty"`
	Eligible    bool      `json:"eligible"`
	Reason      string    `json:"reason,omitempty"`
}

type AutoIndexEntry struct {
	Fingerprint string          `json:"fingerprint"`
	Kind        AutoIndexKind   `json:"kind"`
	Method      string          `json:"method,omitempty"`
	Status      AutoIndexStatus `json:"status"`
	Name        string          `json:"name,omitempty"`
	LastSeen    time.Time       `json:"lastSeen,omitempty"`
	Error       string          `json:"error,omitempty"`
}

type Index struct {
	// Support composite indexes
	Fields []string

	// true = ASC, false = DESC
	Ascending bool

	// true = UNIQUE index
	Unique bool
}

type Op string

const (
	// OpOr allows grouping multiple filters with an OR condition.
	// Example: `field = "value1" OR field = "value2"`
	// SQL Equivalent: `SELECT * FROM table WHERE field = 'value1' OR field = 'value2';`
	OpOr Op = "or"

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

	// OpLessThan selects objects where the field value is less than (<) the given value.
	OpLessThan Op = "lessThan"

	// OpLessThanOrEqual selects objects where the field value is less than or equal to (<=) the given value.
	OpLessThanOrEqual Op = "lessThanOrEqual"

	// OpGreaterThan selects objects where the field value is greater than (>) the given value.
	OpGreaterThan Op = "greaterThan"

	// OpGreaterThanOrEqual selects objects where the field value is greater than or equal to (>=) the given value.
	OpGreaterThanOrEqual Op = "greaterThanOrEqual"
)

type Filter struct {
	Fields []string `json:"fields,omitempty"`

	ValuesJson string `json:"valuesJson,omitempty"`

	Op Op `json:"op"`

	// SubFilters is used for operations like OR where multiple filters are combined.
	SubFilters []Filter `json:"subFilters,omitempty"`
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

	// AfterJson is used for cursor-based pagination, which is more
	// effective in scalable and distributed environments compared
	// to offset-based pagination.
	AfterJson string `json:"afterJson,omitempty"`

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
	SortingTypeRandom  SortingType = "random"
)

type OrderBy struct {
	// The field by which to order the results
	Field string `json:"field,omitempty"`

	// Desc indicates whether the sorting should be in descending order.
	Desc bool `json:"desc,omitempty"`

	// Defines the type of sorting to apply (numeric, text, date, etc.)
	SortingType SortingType `json:"sortingType,omitempty"`
}

// random order. not advised for large datasets due to its slow speed
// in a distributed setting
func OrderByRandom() OrderBy {
	return OrderBy{
		SortingType: SortingTypeRandom,
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

// OpOr allows grouping multiple filters with an OR condition.
// Example: `field = "value1" OR field = "value2"`
// SQL Equivalent: `SELECT * FROM table WHERE field = 'value1' OR field = 'value2';`
func Or(filters ...Filter) Filter {
	return Filter{
		Op:         OpOr,
		SubFilters: filters,
	}
}

// OpEquals selects objects where the value of a field equals (=) the specified value in the query.
// Example: `"fieldValue" = "value"`
// SQL Equivalent: `SELECT * FROM table WHERE field = 'value';`
// Elasticsearch Equivalent:
//
//	{
//	  "query": {
//	    "term": {
//	      "field": "value"
//	    }
//	  }
//	}
func Equals(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpEquals,
	}
}

// OpIntersects selects objects where the slice value of a field intersects with the slice value in the query.
// Example: `["fieldValue2", "fieldValue3"] INTERSECTS ["value1", "value2"]`
// SQL Equivalent: `SELECT * FROM table WHERE field && ARRAY['value1', 'value2'];` (PostgreSQL syntax)
// Elasticsearch Equivalent:
//
//	{
//	  "query": {
//	    "terms_set": {
//	      "field": {
//	        "terms": ["value1", "value2"],
//	        "minimum_should_match_script": {
//	          "source": "1"
//	        }
//	      }
//	    }
//	  }
//	}
func Intersects(fields []string, values []any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal(values),
		Op:         OpIntersects,
	}
}

// OpStartsWith selects all objects where the field's value starts with a particular string.
// Example: `"fieldValue" STARTS WITH "prefix"`
// SQL Equivalent: `SELECT * FROM table WHERE field LIKE 'prefix%';`
// Elasticsearch Equivalent:
//
//	{
//	  "query": {
//	    "prefix": {
//	      "field": "prefix"
//	    }
//	  }
//	}
func StartsWith(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpStartsWith,
	}
}

// OpContainsSubstring selects all objects where the field's value contains a particular substring.
// Example: `"fieldValue" CONTAINS_SUBSTRING "subString"`
// SQL Equivalent: `SELECT * FROM table WHERE field LIKE '%subString%';`
// Elasticsearch Equivalent:
//
//	{
//	  "query": {
//	    "wildcard": {
//	      "field": "*subString*"
//	    }
//	  }
//	}
func ContainsSubstring(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpContainsSubstring,
	}
}

// OpIsInList selects objects where the value of a field is one of the specified values in a list.
// Example: `"fieldValue" IS_IN_LIST ["value1", "value2", "value3"]`
// SQL Equivalent: `SELECT * FROM table WHERE field IN ('value1', 'value2', 'value3');`
// Elasticsearch Equivalent:
//
//	{
//	  "query": {
//	    "terms": {
//	      "field": ["value1", "value2", "value3"]
//	    }
//	  }
//	}
func IsInList(fields []string, values ...any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal(values),
		Op:         OpIsInList,
	}
}

func LessThan(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpLessThan,
	}
}

func LessThanOrEqual(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpLessThanOrEqual,
	}
}

func GreaterThan(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpGreaterThan,
	}
}

func GreaterThanOrEqual(fields []string, value any) Filter {
	return Filter{
		Fields:     fields,
		ValuesJson: marshal([]any{value}),
		Op:         OpGreaterThanOrEqual,
	}
}

func Id(id any) Filter {
	return Filter{
		Fields:     []string{"id"},
		ValuesJson: marshal([]any{id}),
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

func marshal(x any) string {
	bs, _ := json.Marshal(x)
	return string(bs)
}
