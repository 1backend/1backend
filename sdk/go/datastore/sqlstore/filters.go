/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sqlstore

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/openorch/openorch/sdk/go/datastore"
)

func (s *SQLStore) placeholder(counter int) string {
	switch s.placeholderStyle {
	case QuestionMarkPlaceholder:
		return "?"
	case DollarSignPlaceholder:
		return fmt.Sprintf("$%d", counter)
	default:
		panic(fmt.Sprintf("unrecognized placeholder style '%v'", s.placeholderStyle))
	}
}

func (q *SQLQueryBuilder) buildFilters(start ...int) ([]string, []interface{}, error) {
	var params []interface{}
	paramCounter := 1
	if len(start) > 0 {
		paramCounter += start[0]
	}
	var filters []string

	for _, cond := range q.filters {
		var param any
		var err error

		fieldNames := cond.Fields

		if cond.Op == datastore.OpEquals {
			orFilters := []string{}

			var values []any
			err = json.Unmarshal([]byte(cond.JSONValues), &values)
			if err != nil {
				return nil, nil, err
			}

			for _, field := range fieldNames {
				fieldName := q.store.fieldName(field)
				placeHolder := q.store.placeholder(paramCounter)

				parts := strings.Split(field, ".")

				if len(values) > 1 {
					orFilters = append(orFilters, fmt.Sprintf("%s = ANY(%s)", placeHolder, fieldName))
					param, err = q.store.convertParam(values)
				} else if v := q.store.fieldTypes[lowercaseFirstChar(parts[0])]; v != nil && v.Kind() == reflect.Slice {
					orFilters = append(orFilters, fmt.Sprintf("%s = ANY(%s)", placeHolder, fieldName))
					param, err = q.store.convertParam(values[0])
				} else {
					orFilters = append(orFilters, fmt.Sprintf("%s = %s", placeHolder, fieldName))
					param, err = q.store.convertParam(values[0])
				}

				params = append(params, param)
				paramCounter++
			}

			if len(orFilters) == 1 {
				filters = append(filters, orFilters...)
			} else {
				filters = append(filters, fmt.Sprintf("(%s)", strings.Join(orFilters, " OR ")))
			}
		} else if cond.Op == datastore.OpContainsSubstring {
			orFilters := []string{}

			var values []any
			err = json.Unmarshal([]byte(cond.JSONValues), &values)
			if err != nil {
				return nil, nil, err
			}
			if len(values) > 1 {
				panic("OpContainsSubstring does not support multiple values")
			}

			for _, field := range fieldNames {
				fieldName := q.store.fieldName(field)
				placeHolder := q.store.placeholder(paramCounter)

				orFilters = append(orFilters, fmt.Sprintf("%s ILIKE %s", fieldName, placeHolder))
				param, err = q.store.convertParam(fmt.Sprintf("%%%s%%", values[0]))

				params = append(params, param)
				paramCounter++
			}

			if len(orFilters) == 1 {
				filters = append(filters, orFilters...)
			} else {
				filters = append(filters, fmt.Sprintf("(%s)", strings.Join(orFilters, " OR ")))
			}
		} else if cond.Op == datastore.OpStartsWith {
			orFilters := []string{}

			var values []any
			err = json.Unmarshal([]byte(cond.JSONValues), &values)
			if err != nil {
				return nil, nil, err
			}
			if len(values) > 1 {
				panic("OpStartsWith does not support multiple values")
			}

			for _, field := range fieldNames {
				fieldName := q.store.fieldName(field)
				placeHolder := q.store.placeholder(paramCounter)

				orFilters = append(orFilters, fmt.Sprintf("%s ILIKE %s", fieldName, placeHolder))
				param, err = q.store.convertParam(fmt.Sprintf("%s%%", values[0]))

				params = append(params, param)
				paramCounter++
			}

			if len(orFilters) == 1 {
				filters = append(filters, orFilters...)
			} else {
				filters = append(filters, fmt.Sprintf("(%s)", strings.Join(orFilters, " OR ")))
			}
		} else if cond.Op == datastore.OpIsInList {
			orFilters := []string{}

			var values []any
			err = json.Unmarshal([]byte(cond.JSONValues), &values)
			if err != nil {
				return nil, nil, err
			}

			for _, field := range fieldNames {
				fieldName := q.store.fieldName(field)
				placeHolder := q.store.placeholder(paramCounter)

				if reflect.TypeOf(values).Kind() == reflect.Slice {
					orFilters = append(orFilters, fmt.Sprintf("%s = ANY(%s)", fieldName, placeHolder))
					param, err = q.store.convertParam(values)
				} else if typ, hasTyp := q.store.fieldTypes[fieldName]; hasTyp && typ.Kind() == reflect.Slice {
					// "reverse" IN clause
					orFilters = append(orFilters, fmt.Sprintf("%s = ANY(%s)", placeHolder, fieldName))
					param, err = q.store.convertParam(values)
				} else {
					orFilters = append(orFilters, fmt.Sprintf("%s = %s", fieldName, placeHolder))
					param, err = q.store.convertParam(values)
				}

				params = append(params, param)
				paramCounter++
			}

			if len(orFilters) == 1 {
				filters = append(filters, orFilters...)
			} else {
				filters = append(filters, fmt.Sprintf("(%s)", strings.Join(orFilters, " OR ")))
			}
		} else if cond.Op == datastore.OpIntersects {
			orFilters := []string{}

			var values []any
			err = json.Unmarshal([]byte(cond.JSONValues), &values)
			if err != nil {
				return nil, nil, err
			}

			for _, field := range fieldNames {
				fieldName := q.store.fieldName(field)
				placeHolder := q.store.placeholder(paramCounter)

				// Ensure that the values being passed are an array (e.g., converting to PostgreSQL array format)
				if reflect.TypeOf(values).Kind() != reflect.Slice {
					return nil, nil, fmt.Errorf("OpIntersects requires array values")
				}

				// Use PostgreSQL's array intersection operator (&&)
				orFilters = append(orFilters, fmt.Sprintf("%s && %s", fieldName, placeHolder))

				// Convert the values to a PostgreSQL array or the relevant database format
				param, err = q.store.convertParam(values)
				if err != nil {
					return nil, nil, err
				}

				params = append(params, param)
				paramCounter++
			}

			if len(orFilters) == 1 {
				filters = append(filters, orFilters...)
			} else {
				filters = append(filters, fmt.Sprintf("(%s)", strings.Join(orFilters, " OR ")))
			}
		} else {
			panic(fmt.Sprintf("unknown filter %v", cond))
		}

		if err != nil {
			return nil, nil, err
		}
	}

	if len(q.after) > 0 {
		fieldName := q.store.fieldName(q.orderField, castType(q.after[0]))

		placeHolder := q.store.placeholder(paramCounter)
		if q.orderDesc {
			filters = append(filters, fmt.Sprintf("%s < %s", fieldName, placeHolder))
		} else {
			filters = append(filters, fmt.Sprintf("%s > %s", fieldName, placeHolder))
		}
		params = append(params, q.after[0])
		paramCounter++

	}

	return filters, params, nil
}

func lowercaseFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func castType(v any) string {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return "numeric"
	}

	return ""
}

func castTypeSlice(v []any) string {
	switch v[0].(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return "numeric"
	}

	return ""
}
