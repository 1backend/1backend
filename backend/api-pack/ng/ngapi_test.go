package ngapi

import (
	"testing"

	apiTypes "github.com/1backend/1backend/backend/api-pack/types"
	"github.com/1backend/1backend/backend/domain"
)

var proj = &domain.Project{
	Name: "test-service",
	Types: `{
		"user": [
			{"name":   "string"},
			{"age":    "int"},
			{"likesBananas": "bool"}, 
			{"foods":  "string[]"},
			{"friend": "curfter/other-service.Person"}
		],
		"thing": [
			{"id": "string"},
			{"fieldA": "janos/service.ThatThing"}
		]
	}`,
	Endpoints: []domain.Endpoint{
		domain.Endpoint{
			Method: "GET",
			Url:    "/hi-there",
			Input: `[
				{"userId":  "string"},
				{"friends": "user[]"}
			]`,
			Output: `crufter/test.user[]`,
		},
		domain.Endpoint{
			Method: "POST",
			Url:    "/save-something",
			Input: `[
				{"thing":  "thing"},
				{"tag": "string[]"}
			]`,
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
	for _, file := range files {
		t.Log(file[0])
		t.Log("====")
		t.Log(file[1])
	}
}
