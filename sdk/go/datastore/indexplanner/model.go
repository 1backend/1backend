package indexplanner

import (
	"reflect"
	"strings"
	"time"
)

type modelInfo struct {
	order []string
	roots map[string]reflect.Type
}

type resolvedField struct {
	Field  string
	Path   []string
	Type   reflect.Type
	Scalar bool
	Slice  bool
	Text   bool
}

func describeModel(instance any) modelInfo {
	typ := reflect.TypeOf(instance)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	info := modelInfo{
		roots: map[string]reflect.Type{},
	}

	var walk func(reflect.Type)
	walk = func(t reflect.Type) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				walk(field.Type)
				continue
			}

			name := fieldName(field)
			info.order = append(info.order, name)
			info.roots[name] = field.Type
		}
	}

	walk(typ)
	return info
}

func resolveField(info modelInfo, name string) (resolvedField, bool) {
	if name == "" {
		return resolvedField{}, false
	}

	parts := strings.Split(name, ".")
	rootType, ok := info.roots[parts[0]]
	if !ok {
		return resolvedField{}, false
	}

	current := rootType
	for _, pathPart := range parts[1:] {
		current = derefType(current)
		if current.Kind() != reflect.Struct || current == reflect.TypeOf(time.Time{}) {
			return resolvedField{}, false
		}

		found := false
		for i := 0; i < current.NumField(); i++ {
			field := current.Field(i)
			if !field.IsExported() {
				continue
			}
			if fieldName(field) == pathPart {
				current = field.Type
				found = true
				break
			}
		}
		if !found {
			return resolvedField{}, false
		}
	}

	base := derefType(current)
	return resolvedField{
		Field:  parts[0],
		Path:   append([]string(nil), parts[1:]...),
		Type:   current,
		Scalar: isScalarType(current),
		Slice:  base.Kind() == reflect.Slice,
		Text:   derefType(current).Kind() == reflect.String,
	}, true
}

func derefType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func isScalarType(t reflect.Type) bool {
	t = derefType(t)
	if t == reflect.TypeOf(time.Time{}) {
		return true
	}

	switch t.Kind() {
	case reflect.String, reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func fieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag != "" {
		if idx := strings.Index(tag, ","); idx >= 0 {
			tag = tag[:idx]
		}
		if tag != "" {
			return tag
		}
	}
	name := field.Name
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}

func schemaRank(info modelInfo, field string) int {
	for i, name := range info.order {
		if name == field {
			return i
		}
	}
	return len(info.order) + 1
}
