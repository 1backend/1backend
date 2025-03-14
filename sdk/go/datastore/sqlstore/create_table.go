/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sqlstore

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (s *SQLStore) createTable(instance any, db DB, tableName string) (map[string]reflect.Type, error) {
	typeMap := map[string]reflect.Type{}

	typ := reflect.TypeOf(instance)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	fieldNames := []string{}
	fieldTypes := []string{}

	// Recursive function to process struct fields (including embedded fields)
	var processFields func(reflect.Type)
	processFields = func(typ reflect.Type) {
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if !field.IsExported() {
				continue
			}

			// Handle embedded structs by recursion
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				processFields(field.Type)
				continue
			}

			// Use the json tag for the field name, if available, and strip out ",omitempty"
			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = s.fieldName(field.Name)
			} else if idx := strings.Index(fieldName, ","); idx != -1 {
				fieldName = fieldName[:idx]
			}
			fieldName = escape(fieldName)

			typeMap[fieldName] = field.Type

			// Map field type to SQL type
			fieldType := s.sqlType(field.Type)

			fieldNames = append(fieldNames, fieldName)
			fieldTypes = append(fieldTypes, fieldType)
		}
	}

	// Process all fields of the instance's struct
	processFields(typ)

	if tableName == "" {
		tableName = strings.ToLower(typ.Name())
	}
	createQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ();", tableName)
	for index, fieldName := range fieldNames {
		createQuery += fmt.Sprintf(" ALTER TABLE %s ADD COLUMN IF NOT EXISTS %s %s;", tableName, fieldName, fieldTypes[index])
	}

	_, err := db.Exec(createQuery)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to create table with query '%v'", createQuery))
	}

	return typeMap, nil
}

func (s *SQLStore) sqlType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "TEXT"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return "INTEGER"
	case reflect.Int64:
		return "BIGINT"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "INTEGER"
	case reflect.Uint64:
		return "BIGINT"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Map:
		switch s.driverName {
		case DriverPostGRES:
			return "JSONB"
		case DriverMySQL:
			return "JSON"
		}
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP"
		}
		switch s.driverName {
		case DriverPostGRES:
			return "JSONB"
		case DriverMySQL:
			return "JSON"
		}
	case reflect.Ptr:
		return s.sqlType(t.Elem())
	case reflect.Slice:
		switch s.driverName {
		case DriverMySQL:
			return "JSON"
		case DriverPostGRES:
			elemType := s.sqlType(t.Elem())
			if elemType != "JSONB" {
				// You can store as JSONB, or as native array if preferred
				return fmt.Sprintf("%s[]", elemType) // Use PostgreSQL array type
			}

			// Default to JSONB for complex slices or fallback
			return "JSONB"
		}

	default:
		return "TEXT" // Default to TEXT for unknown types
	}
	return "TEXT"
}
