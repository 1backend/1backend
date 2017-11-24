package goapi

import (
	"fmt"
	"testing"

	apiTypes "github.com/1backend/1backend/backend/api-packs/types"
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
	str, err := NewGenerator(proj).GenerateApi(*c)
	if err != nil {
		t.Fatal(err)
	}
	if str == "" {
		t.Fatal("Output is empty")
	}
	fmt.Println(str)
	t.Fatal(str)
	t.Log(str)
}
