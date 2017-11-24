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
package {{ .PackageName }}

import(
{{ range $key, $value := .Imports }}\t{{ $value }}\n{{  }}
)

{{ range $typeName, $keyValues := .Types }}
type {{ $typeName | gTypeName }} struct {
	{{ range $fieldName, $fieldType := $keyValues }}{{ $fieldName }}\t\t{{ $fieldType }}\n{{ end }}
}
{{ end }}

{{ range $key, $value := .Endpoints }}func {{ .Path | gName}}({{ .Input | generateInput }}) {{ .Input | gOutput }} {
	
}{{ end }}`

func gInput(inp string) string {
	return ""
}

func gOutput(outp string) string {
	return ""
}

func (g GoGenerator) generateGo(c apiTypes.Context) (string, error) {
	// Create a new template and parse the letter into it.
	templFuncs := template.FuncMap{
		"gName": func(method, path string) string {
			s := strings.Replace(path, "/", "_", -1)
			s = strings.Replace(s, "-", "_", -1)
			return strings.Title(strings.ToLower(method)) + strcase.ToCamel(s)
		},
		"gInputs":  gInput,
		"gOutputs": gOutput,
		"trim":     strings.TrimSpace,
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
