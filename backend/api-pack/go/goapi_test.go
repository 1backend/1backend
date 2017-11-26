package goapi

import (
	"testing"

	apiTypes "github.com/1backend/1backend/backend/api-pack/types"
	"github.com/1backend/1backend/backend/domain"
)

var proj = &domain.Project{
	Types: `{
		"user": {
			"name":   "string",
			"age":    "int",
			"foods":  "string[]",
			"friend": "curfter/other-service.Person"
		}
	}`,
	Endpoints: []domain.Endpoint{
		domain.Endpoint{
			Method: "GET",
			Url:    "/hi-there",
			Input: `{
				"userId":  "string",
				"friends": "user[]"
			}`,
			Output: `crufter/test.user[]`,
		},
	},
}

func TestBasic(t *testing.T) {
	c, err := apiTypes.GetContext(proj)
	if err != nil {
		t.Fatal(err)
	}
	files, err := NewGenerator(proj).FilesToBuild(*c)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 || files[0][1] == "" {
		t.Fatal("Output is empty")
	}
	t.Log(files[0][1])
}
