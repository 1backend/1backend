package indexplanner

import (
	"reflect"
	"strings"
	"time"
)

type modelInfo struct {
	order []string
	roots map[string]modelField
}

type modelField struct {
	Name string
	Type reflect.Type
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
		roots: map[string]modelField{},
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
			addFieldAliases(info.roots, field, modelField{
				Name: name,
				Type: field.Type,
			})
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
	root, ok := lookupModelField(info.roots, parts[0])
	if !ok {
		return resolvedField{}, false
	}

	current := root.Type
	var path []string
	if len(parts) > 1 {
		path = make([]string, 0, len(parts)-1)
	}
	for _, pathPart := range parts[1:] {
		current = derefType(current)
		if current.Kind() != reflect.Struct || current == reflect.TypeOf(time.Time{}) {
			return resolvedField{}, false
		}

		field, ok := resolveStructField(current, pathPart)
		if !ok {
			return resolvedField{}, false
		}
		current = field.Type
		path = append(path, field.Name)
	}

	base := derefType(current)
	return resolvedField{
		Field:  root.Name,
		Path:   path,
		Type:   current,
		Scalar: isScalarType(current),
		Slice:  base.Kind() == reflect.Slice,
		Text:   derefType(current).Kind() == reflect.String,
	}, true
}

func addFieldAliases(fields map[string]modelField, field reflect.StructField, target modelField) {
	for _, alias := range fieldAliases(field) {
		if alias == "" || alias == "-" {
			continue
		}
		if _, ok := fields[alias]; !ok {
			fields[alias] = target
		}
		lower := strings.ToLower(alias)
		if _, ok := fields[lower]; !ok {
			fields[lower] = target
		}
	}
}

func lookupModelField(fields map[string]modelField, name string) (modelField, bool) {
	if field, ok := fields[name]; ok {
		return field, true
	}
	field, ok := fields[strings.ToLower(name)]
	return field, ok
}

func resolveStructField(t reflect.Type, name string) (modelField, bool) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}
		target := modelField{
			Name: fieldName(field),
			Type: field.Type,
		}
		for _, alias := range fieldAliases(field) {
			if alias == "" || alias == "-" {
				continue
			}
			if alias == name || strings.EqualFold(alias, name) {
				return target, true
			}
		}
	}
	return modelField{}, false
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

func fieldAliases(field reflect.StructField) []string {
	canonical := fieldName(field)
	aliases := []string{canonical, field.Name}
	if field.Name != "" {
		aliases = append(aliases, strings.ToLower(field.Name[:1])+field.Name[1:])
	}
	return aliases
}

func schemaRank(info modelInfo, field string) int {
	for i, name := range info.order {
		if name == field {
			return i
		}
	}
	return len(info.order) + 1
}
