package techpack

import (
	"encoding/json"
	"fmt"

	"github.com/1backend/1backend/backend/api-packs/go"
	apiTypes "github.com/1backend/1backend/backend/api-packs/types"
	"github.com/1backend/1backend/backend/domain"

	"strings"
)

// GetProvider returns the plugin the project uses based on its mode and recipe
func GetGenerator(proj *domain.Project, language string) Generator {
	switch {
	case language == "go":
		return goapi.NewGenerator(proj)
	}
	panic(fmt.Sprintf("Can't find generator %v for language", language))
}

type Generator interface {
	GenerateApi(apiTypes.Context) (string, error)
}

// example project
//
//proj := &domain.Project{
//	Types: `
//	"user": map[string]string{
//		"name":   "string",
//		"age":    "int",
//		"foods":  "string[]",
//		"friend": "user",
//	}`,
//	Endpoints: []domain.Endpoint{
//		domain.Endpoint{
//			Input: `{
//				"userId":  "string",
//				"friends": "user[]"
//			}`,
//			Output: `crufter/test.user[]`,
//		},
//	},
//}

// goes over types and signatures and collects all imports
func collectImports(types map[string]apiTypes.Type, sigs []apiTypes.EndpointSignature) []apiTypes.Import {
	index := map[string]bool{}
	imports := []apiTypes.Import{}
	for _, v := range types {
		for _, v1 := range v.Fields {
			if v1.Import.Author != "" && !index[v1.Import.Author+"_"+v1.Import.ProjectName] {
				imports = append(imports, v1.Import)
				index[v1.Import.Author+"_"+v1.Import.ProjectName] = true
			}
		}
	}
	return imports
}

func parseSignatures(eps []domain.Endpoint) []apiTypes.EndpointSignature {
	ret := []apiTypes.EndpointSignature{}
	for _, endpoint := range eps {
		input := map[string]string{}
		err := json.Unmarshal([]byte(endpoint.Input), &input)
		if err != nil {
			panic(err)
		}
		typ := parseType("", input)
		ret = append(ret, apiTypes.EndpointSignature{
			Method: endpoint.Method,
			Path:   endpoint.Url,
			Input:  typ.Fields,
		})
	}
	return ret
}

// parseTypes parses a whole type definition
func parseTypes(types map[string]map[string]string) map[string]apiTypes.Type {
	ts := map[string]apiTypes.Type{}
	for typeName, fields := range types {
		ts[typeName] = parseType(typeName, fields)
	}
	return ts
}

func parseField(fieldName, fieldType string) apiTypes.Field {
	isList := strings.HasSuffix(fieldType, "[]")
	if strings.Contains(fieldType, "/") {
		parts := strings.Split("typ", ".")
		ap := strings.Split(parts[0], "/")
		return apiTypes.Field{
			Name:   fieldName,
			Type:   parts[1],
			IsList: isList,
			Import: apiTypes.Import{
				Author:      ap[0],
				ProjectName: ap[1],
			},
		}
	}
	return apiTypes.Field{
		Name:   fieldName,
		Type:   fieldType,
		IsList: isList,
	}
}

func parseType(name string, typ map[string]string) apiTypes.Type {
	t := apiTypes.Type{
		Name:   name,
		Fields: []apiTypes.Field{},
	}
	for fieldName, fieldType := range typ {
		t.Fields = append(t.Fields, parseField(fieldName, fieldType))
	}
	return t
}

func GetContext(proj *domain.Project) (*apiTypes.Context, error) {
	typs := map[string]map[string]string{}
	err := json.Unmarshal([]byte(proj.Types), &typs)
	if err != nil {
		return nil, err
	}
	types := parseTypes(typs)
	sigs := parseSignatures(proj.Endpoints)
	imports := collectImports(types, sigs)
	return &apiTypes.Context{
		Imports: imports,
		Types:   types,
	}, nil
}
