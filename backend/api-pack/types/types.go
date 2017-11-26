package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/1backend/1backend/backend/domain"
)

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
	TypeDefinitions    map[string]TypeDefinition
	Imports            []Import
	EndpointSignatures []EndpointSignature
}

type Import struct {
	Author      string
	ProjectName string
}

func GetContext(proj *domain.Project) (*Context, error) {
	typs := map[string]map[string]string{}
	err := json.Unmarshal([]byte(proj.Types), &typs)
	if err != nil {
		return nil, err
	}
	types := parseTypeDefinitions(typs)
	sigs, err := parseSignatures(proj.Endpoints)
	if err != nil {
		return nil, err
	}
	imports := collectImports(types, sigs)
	return &Context{
		Imports:            imports,
		TypeDefinitions:    types,
		EndpointSignatures: sigs,
		ProjectName:        proj.Name,
		Author:             proj.Author,
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
		input := map[string]string{}
		err := json.Unmarshal([]byte(endpoint.Input), &input)
		if err != nil {
			return nil, err
		}
		typ := parseTypeDefinition("", input)
		ret = append(ret, EndpointSignature{
			Method: endpoint.Method,
			Path:   endpoint.Url,
			Input:  typ.Fields,
			Output: parseType(endpoint.Output),
		})
	}
	return ret, nil
}

// parseTypes parses a whole type definition
func parseTypeDefinitions(types map[string]map[string]string) map[string]TypeDefinition {
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

func parseTypeDefinition(name string, typ map[string]string) TypeDefinition {
	t := TypeDefinition{
		Name:   name,
		Fields: []Field{},
	}
	for fieldName, fieldType := range typ {
		t.Fields = append(t.Fields, Field{
			Name: fieldName,
			Type: parseType(fieldType),
		})
	}
	return t
}
