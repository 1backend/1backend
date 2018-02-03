package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/1backend/1backend/backend/domain"
)

type ClientPlugin interface {
	// Pretty name of language/tech
	Name() string
	// Called when generating client files for each language
	ClientFiles(c Context) (*ClientFiles, error)
}

type EndpointSignature struct {
	Method string
	Path   string
	Input  []Field
	Output Type
}

type Type struct {
	IsList bool
	Import Import // empty for local packages
	Name   string
}

// Reprensents both struct definition fields
// and input arguments (both are a list of name - type pairs)
type Field struct {
	Name string
	Type Type // not recursive
}

type TypeDefinition struct {
	Name   string
	Fields []Field
}

type Context struct {
	Author             string
	ProjectName        string
	ProjectVersion     string
	ProjectNames       []string
	TypeDefinitions    map[string]TypeDefinition
	Imports            []Import
	EndpointSignatures []EndpointSignature
	Meta               map[string]interface{}
}

type Import struct {
	Author      string
	ProjectName string
}

type ClientFiles struct {
	FolderName string
	Files      [][]string // list of (fileName, fileContent) tuple
}

func GetContext(project *domain.Project, projectNames []string) (*Context, error) {
	types := map[string]TypeDefinition{}
	if project.Types != "" {
		typs := map[string][]map[string]string{}
		err := json.Unmarshal([]byte(project.Types), &typs)
		if err != nil {
			return nil, err
		}
		types = parseTypeDefinitions(typs)
	}
	sigs, err := parseSignatures(project.Endpoints)
	if err != nil {
		return nil, err
	}
	imports := collectImports(types, sigs)
	return &Context{
		Imports:            imports,
		TypeDefinitions:    types,
		EndpointSignatures: sigs,
		ProjectName:        project.Name,
		ProjectNames:       projectNames,
		Author:             project.Author,
		ProjectVersion:     project.Version,
	}, nil
}

// goes over types and signatures and collects all imports
func collectImports(types map[string]TypeDefinition, sigs []EndpointSignature) []Import {
	index := map[string]bool{}
	imports := []Import{}
	for _, v := range types {
		for _, v1 := range v.Fields {
			if v1.Type.Import.Author != "" && !index[v1.Type.Import.Author+"_"+v1.Type.Import.ProjectName] {
				imports = append(imports, v1.Type.Import)
				index[v1.Type.Import.Author+"_"+v1.Type.Import.ProjectName] = true
			}
		}
	}
	return imports
}

func parseSignatures(eps []domain.Endpoint) ([]EndpointSignature, error) {
	ret := []EndpointSignature{}
	for _, endpoint := range eps {
		inputTypeDef := TypeDefinition{}
		if endpoint.Input != "" {
			input := []map[string]string{}
			err := json.Unmarshal([]byte(endpoint.Input), &input)
			if err != nil {
				return nil, fmt.Errorf("Endoint \"%v\" input json is invalid: %v", endpoint.Url, err)
			}
			inputTypeDef = parseTypeDefinition("", input)
		}
		outputType := Type{}
		if endpoint.Output != "" {
			outputType = parseType(endpoint.Output)
		}
		ret = append(ret, EndpointSignature{
			Method: endpoint.Method,
			Path:   endpoint.Url,
			Input:  inputTypeDef.Fields,
			Output: outputType,
		})
	}
	return ret, nil
}

// parseTypes parses a whole type definition
func parseTypeDefinitions(types map[string][]map[string]string) map[string]TypeDefinition {
	ts := map[string]TypeDefinition{}
	for typeName, fields := range types {
		ts[typeName] = parseTypeDefinition(typeName, fields)
	}
	return ts
}

func parseType(fieldType string) Type {
	isList := strings.HasSuffix(fieldType, "[]")
	if strings.Contains(fieldType, "/") {
		parts := strings.Split(fieldType, ".")
		ap := strings.Split(parts[0], "/")
		fmt.Println("ap", ap)
		return Type{
			Name: parts[1],
			Import: Import{
				Author:      ap[0],
				ProjectName: ap[1],
			},
			IsList: isList,
		}
	}
	return Type{
		Name:   fieldType,
		IsList: isList,
	}
}

func parseTypeDefinition(name string, typ []map[string]string) TypeDefinition {
	t := TypeDefinition{
		Name:   name,
		Fields: []Field{},
	}
	for _, fieldPairMap := range typ {
		for fieldName, fieldType := range fieldPairMap {
			t.Fields = append(t.Fields, Field{
				Name: fieldName,
				Type: parseType(fieldType),
			})
		}
	}
	return t
}
