package ngapi

import (
	"bytes"
	"strings"
	"text/template"

	apiTypes "github.com/1backend/1backend/backend/api-pack/types"
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

func (g GoGenerator) FilesToBuild(c apiTypes.Context) ([][]string, error) {
	return g.generateFiles(c)
}

func (g GoGenerator) FolderName() string {
	return "go"
}

var indexTemplate = `
{{ range $key, $export := .Projects }} export * from './{{ $.ProjectName }}';

let Token: string
`
var serviceTemplate = `import { Injectable } from '@angular/core';
import * as ngClient from '@1backend/ng-client';
{{ range $key, $import := .Imports }}import * as {{ $import.Author }} from  '@1backend/{{ $import.Author }}-ng
{{ end }}

{{ range $typeName, $type := .TypeDefinitions }}
export interface {{ $typeName | gTypeName }} {
{{ range $index, $field := $type.Fields }}	{{ $field.Name | gFieldName }}:	{{ $field.Type | gType }};
{{ end }}}
{{ end }}

@Injectable()
export class {{ $.ProjectName | gServiceName }} {
  constructor() {}

  {{ range $key, $sig := .EndpointSignatures }}{{ $sig.Method | gMethod }}{{ $sig.Path | gPathAsFunc }}({{ $sig.Input | gInput }}) {{ $sig.Output | gOutput }} {
	return ret, new ngClient<{{ $sig.Output | gType }}>.(Token).Call("{{ $.Author }}", "{{ $.ProjectName }}", "{{ $sig.Method }}", "{{ $sig.Path }}", { {{ range $index, $field := $sig.Input }}{{if ne $index 0}}, {{end}}"{{ $field.Name }}": {{ $field.Name }}{{ end }} });
	}
{{ end }}
}


`

var packagejsonTemplate = `{
	"name": "david-test-xd-xd",
	"version": "0.0.2",
	"description": "David test xd",
	"main": "./lib/index.js",
	"typings": "./lib/index.d.ts",
	"scripts": {
	  "prepare": "npm run build",
	  "build": "tsc"
	},
	"keywords": [
	  "typescript",
	  "node"
	],
	"repository": {
	  "type": "git",
	  "url": "https://github.com/dobika/typescript-example-service.git"
	},
	"dependencies": {
	  "@angular/core": "5.0.0",
	  "@angular/common": "5.0.0",
	  "rxjs": "5.5.2",
	  "zone.js": "0.8.14"
	}
  }
`
var tsconfigjsonTemplate = `{
	"compilerOptions": {
	  "target": "ES2015",
	  "module": "commonjs",
	  "declaration": true,
	  "experimentalDecorators": true,
	  "outDir": "./lib",
	  "strict": true
	},
	"include": ["src/**/*"],
	"exclude": ["src/**/*.spec.*"]
  }
`

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
			return prefix + "." + correctTypeName(f.Name) + "[]"
		}
		return prefix + "." + correctTypeName(f.Name)
	}
	if f.IsList {
		return correctTypeName(f.Name) + "[]"
	}
	return correctTypeName(f.Name)
}

func gServiceName(s string) string {
	s = strings.Replace(s, "-", "_", -1)
	return strcase.ToCamel(s)
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

	return "Promise<" + gType(outp) + ">"
}

func gFieldName(s string) string {
	return s
}

func correctTypeName(s string) string {
	switch s {
	case "bool":
		return "boolean"
	case "bool[]":
		return "boolean"
	case "int":
		return "number"
	case "int[]":
		return "number"
	case "float":
		return "number"
	case "float[]":
		return "number"
	case "string":
		return s
	case "string[]":
		return "string"
	}
	return strings.Title(strings.Replace(s, "[]", "", 1))
}

func (g GoGenerator) generateFiles(c apiTypes.Context) ([][]string, error) {
	templFuncs := template.FuncMap{
		"gServiceName": gServiceName,
		"gInputToMap":  gInputToMap,
		"gType":        gType,
		"gTypeName":    gTypeName,
		"gMethod":      gMethod,
		"gInput":       gInput,
		"gFieldName":   gFieldName,
		"gOutput":      gOutput,
		"gPathAsFunc":  gPathAsFunc,
		"trim":         strings.TrimSpace,
	}
	ret := [][]string{}
	titles := []string{"index.ts", "package.json", c.ProjectName + ".service.ts"}
	toParse := []string{indexTemplate, packagejsonTemplate, serviceTemplate}
	for i, f := range toParse {
		templ := template.New(titles[i]).Funcs(templFuncs)
		t, err := templ.Parse(string(f))

		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer([]byte{})
		err = t.Execute(buf, c)
		ret = append(ret, []string{titles[i], buf.String()})
	}

	return ret, nil
}
