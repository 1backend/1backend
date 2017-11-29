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
	return "ng"
}

var indexTemplate = `
export * from './{{.ProjectName}}.service';
`

var serviceTemplate = `import { Injectable } from '@angular/core';
import { NgClient } from '@1backend/ng-client';
{{ range $key, $import := .Imports }}import * as {{ $import.Author }}_{{ $import.ProjectName }} from  '@1backend/{{ $import.Author }}-ng/{{ $import.ProjectName }}.service';
{{ end }}

{{ range $typeName, $type := .TypeDefinitions }}
export interface {{ $typeName | gTypeName }} {
{{ range $index, $field := $type.Fields }}	{{ $field.Name | gFieldName }}:	{{ $field.Type | gType }};
{{ end }}}
{{ end }}

@Injectable()
export class Service {
  constructor(private ngClient: NgClient) {}

{{ range $key, $sig := .EndpointSignatures }}  {{ $sig.Method | gMethod }}{{ $sig.Path | gPathAsFunc }}({{ $sig.Input | gInput }}): Promise<{{ $sig.Output | gType }}> {
    return this.ngClient.call<{{ $sig.Output | gType }}>("{{ $.Author }}", "{{ $.ProjectName }}", "{{ $sig.Method }}", "{{ $sig.Path }}", { {{ range $index, $field := $sig.Input }}{{if ne $index 0}}, {{end}}"{{ $field.Name }}": {{ $field.Name }}{{ end }} });
  }

{{ end }}}
`

var packageJsonTemplate = `{
	"name": "@1backend/{{ .Author }}-{{ .ProjectName }}-ng",
	"version": "{{ .ProjectVersion }}",
	"description": "Clients for 1Backend services of {{ .Author }}",
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
	  "url": "https://github.com/1backend/{{ .Author }}.git"
	},
	"dependencies": {
	  "@1backend/ng-client": "*"{{ range $key, $import := .Imports }},
	  "@1backend/{{ $import.Author }}-ng": "*";
{{ end }}
	},
	"devDependencies": {
		"@angular/cli": "^1.5.3",
		"@angular/compiler-cli": "^5.0.0",
		"typescript": "~2.4.2",
		"@angular/core": "^5.0.0",
		"@angular/common": "^5.0.0",
		"rxjs": "^5.5.2",
		"zone.js": "^0.8.14"
	}
  }
`

var gitIgnoreTemplate = `node_modules
lib
`

var npmIgnoreTemplate = `node_modules
`

var tsconfigJsonTemplate = `{
    "compilerOptions": {
      "target": "es5",
      "module": "commonjs",
      "declaration": true,
      "lib" : ["es2015", "dom"],
      "experimentalDecorators": true,
      "emitDecoratorMetadata": true,
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
	if f.Name == "" {
		return "void"
	}
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
		inputs = append(inputs, field.Name+": "+gType(field.Type))
	}
	return strings.Join(inputs, ", ")
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
		"gPathAsFunc":  gPathAsFunc,
		"trim":         strings.TrimSpace,
	}
	ret := [][]string{}
	titles := []string{"src/index.ts", "tsconfig.json", "package.json", "src/" + c.ProjectName + ".service.ts", ".gitignore", ".npmignore"}
	toParse := []string{indexTemplate, tsconfigJsonTemplate, packageJsonTemplate, serviceTemplate, gitIgnoreTemplate, npmIgnoreTemplate}
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
