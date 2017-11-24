package goapi

import (
	"bytes"
	"html/template"
	"strings"

	apiTypes "github.com/1backend/1backend/backend/api-packs/types"
	"github.com/1backend/1backend/backend/domain"
	"github.com/iancoleman/strcase"
)

func NewGenerator(proj *domain.Project) GoGenerator {
	return GoGenerator{
		Project: proj,
	}
}

type GoGenerator struct {
	Project *domain.Project
}

func (g GoGenerator) GenerateApi(c apiTypes.Context) (string, error) {
	return g.generateGo(c)
}

var goTemplate = `
package {{ .ProjectName }}

import(
	"github.com/1backend/go-client"
{{ range $key, $import := .Imports }}	{{ $import.Author }}_{{ $import.ProjectName }} "github.com/1backend/{{ $import.Author }}/{{ $import.ProjectName }}"
{{ end }}
)

{{ range $typeName, $type := .TypeDefinitions }}
type {{ $typeName | gTypeName }} struct {
	{{ range $index, $field := $type.Fields }}{{ $field.Name | gFieldName }}	{{ $field.Type | gType }}
	{{ end }}
}
{{ end }}

{{ range $key, $sig := .EndpointSignatures }}func {{ $sig.Method | gMethod }}{{ $sig.Path | gPathAsFunc }}({{ $sig.Input | gInput }}) {{ $sig.Output | gType }} {
	var ret {{ $sig.Output | gType }}
	return ret, goclient.Call(&ret)
}{{ end }}`

func gPathAsFunc(path string) string {
	s := strings.Replace(path, "/", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	return strcase.ToCamel(s)
}

func gTypeName(s string) string {
	return strings.Title(s)
}

func gType(f apiTypes.Type) string {
	if f.Import.Author != "" {
		prefix := f.Import.Author + "_" + f.Import.ProjectName
		if f.IsList {
			return "[]" + prefix + "." + correctTypeName(f.Name)
		}
		return prefix + "." + correctTypeName(f.Name)
	}
	if f.IsList {
		return "[]" + correctTypeName(f.Name)
	}
	return correctTypeName(f.Name)
}

func gMethod(method string) string {
	return strings.Title(strings.ToLower(method))
}

func gInput(fields []apiTypes.Field) string {
	inputs := []string{}
	for _, field := range fields {
		inputs = append(inputs, field.Name+" "+gType(field.Type))
	}
	return strings.Join(inputs, ", ")
}

func gOutput(outp string) string {
	return ""
}

func gFieldName(s string) string {
	return strings.Title(s)
}

func correctTypeName(s string) string {
	switch s {
	case "bool":
		return "bool"
	case "bool[]":
		return "[]bool"
	case "int":
		return "int64"
	case "int[]":
		return "[]int"
	case "float":
		return "float64"
	case "float[]":
		return "[]float64"
	case "string":
		return s
	case "string[]":
		return "[]string"
	}
	return strings.Title(strings.Replace(s, "[]", "", 1))
}

func (g GoGenerator) generateGo(c apiTypes.Context) (string, error) {
	// Create a new template and parse the letter into it.
	templFuncs := template.FuncMap{
		"gType":       gType,
		"gTypeName":   gTypeName,
		"gMethod":     gMethod,
		"gInput":      gInput,
		"gFieldName":  gFieldName,
		"gOutput":     gOutput,
		"gPathAsFunc": gPathAsFunc,
		"trim":        strings.TrimSpace,
	}
	t, err := template.New("code").Funcs(templFuncs).Parse(string(goTemplate))
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, c)
	return buf.String(), err
}
