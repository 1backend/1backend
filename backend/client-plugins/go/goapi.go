package goclient

import (
	"bytes"
	"html/template"
	"strings"

	apiTypes "github.com/1backend/1backend/backend/client-plugins/types"
	"github.com/1backend/1backend/backend/domain"
	"github.com/iancoleman/strcase"
)

func New(proj *domain.Project) GoClient {
	return GoClient{
		Project: proj,
	}
}

type GoClient struct {
	Project *domain.Project
}

func (g GoClient) Name() string {
	return "Go"
}

func (g GoClient) ClientFiles(c apiTypes.Context) (*apiTypes.ClientFiles, error) {
	code, err := g.generateGo(c)
	if err != nil {
		return nil, err
	}
	return &apiTypes.ClientFiles{
		FolderName: "go",
		Files: [][]string{
			[]string{
				"client.go",
				code,
			},
		}}, nil
}

var goTemplate = `package {{ .ProjectName }}

import(
	"github.com/1backend/go-client"
{{ range $key, $import := .Imports }}	{{ $import.Author }}_{{ $import.ProjectName }} "github.com/1backend/{{ $import.Author }}/{{ $import.ProjectName }}"
{{ end }})

var Token string
{{ range $typeName, $type := .TypeDefinitions }}
type {{ $typeName | gTypeName }} struct {
{{ range $index, $field := $type.Fields }}	{{ $field.Name | gFieldName }}	{{ $field.Type | gType }}
{{ end }}}
{{ end }}

{{ range $key, $sig := .EndpointSignatures }}func {{ $sig.Method | gMethod }}{{ $sig.Path | gPathAsFunc }}({{ $sig.Input | gInput }}) {{ $sig.Output | gOutput }} {
	var ret {{ $sig.Output | gType }}
	return ret, goclient.New(Token).Call("{{ $.Author }}", "{{ $.ProjectName }}", "{{ $sig.Method }}", "{{ $sig.Path }}", map[string]interface{}{ {{ range $index, $field := $sig.Input }}{{if ne $index 0}}, {{end}}"{{ $field.Name }}": {{ $field.Name }}{{ end }} }, &ret)
}

{{ end }}`

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

func gInputToMap(fields []apiTypes.Field) string {
	inputs := []string{}
	for _, field := range fields {
		inputs = append(inputs, "\""+field.Name+"\":"+field.Name)
	}
	return "{" + strings.Join(inputs, ", ") + "}"
}

func gInput(fields []apiTypes.Field) string {
	inputs := []string{}
	for _, field := range fields {
		inputs = append(inputs, field.Name+" "+gType(field.Type))
	}
	return strings.Join(inputs, ", ")
}

func gOutput(outp apiTypes.Type) string {
	if outp.Name == "" {
		return "error"
	}
	return "(" + gType(outp) + ", error)"
}

func gFieldName(s string) string {
	return strings.Title(s)
}

func correctTypeName(s string) string {
	switch s {
	case "bool":
		return "bool"
	case "bool[]":
		return "bool"
	case "int":
		return "int64"
	case "int[]":
		return "int"
	case "float":
		return "float64"
	case "float[]":
		return "float64"
	case "string":
		return s
	case "string[]":
		return "string"
	}
	return strings.Title(strings.Replace(s, "[]", "", 1))
}

func (g GoClient) generateGo(c apiTypes.Context) (string, error) {
	// Create a new template and parse the letter into it.
	templFuncs := template.FuncMap{
		"gInputToMap": gInputToMap,
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
