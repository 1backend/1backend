package goplugin

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	techt "github.com/1backend/1backend/backend/tech-plugins/types"
	"github.com/1backend/1backend/backend/tech-plugins/utils"
)

func NewPlugin(project *domain.Project) GoPlugin {
	return GoPlugin{
		project: project,
	}
}

type GoPlugin struct {
	project *domain.Project
}

func (g GoPlugin) RecipePath() string {
	return "go"
}

func (g GoPlugin) PreCreate() error {
	return generateEndpoints(g.project)
}

func (g GoPlugin) Build(t *template.FuncMap) (*techt.Build, error) {
	return &techt.Build{
		Outfile:    "main.go",
		RecipePath: "go",
	}, nil
}

const hi = `func (w http.ResponseWriter, r *http.Request, p httpr.Params) {
  utils.Write(w, "hi")
}
`
const hiInput = `[]`
const hiOutput = "string"

const importedHi = `func (w http.ResponseWriter, r *http.Request, p httpr.Params) {
  example.Hi(w, r, p)
}
`
const importedHiInput = `[]`
const importedHiOutput = "string"

const sqlExample = `func (w http.ResponseWriter, r *http.Request, p httpr.Params) {
	rows, err := sqlClient.Query("SELECT 1 + 1 AS solution")
	if err != nil {
		utils.Write500(w, err)
		return
	}
	defer rows.Close()
	solution := ""
	rows.Next()
	err = rows.Scan(&solution)
	if err != nil {
	    utils.Write500(w, err)
		return
	}
	utils.Write(w, solution)
}
`
const sqlExampleInput = `[]`
const sqlExampleOutput = "string"

const inputExample = `example.CalculateRectangleArea`
const inputExampleInput = `[
	{"rect": "rectangle"},
	{"unit": "string"}
]`
const inputExampleOutput = `output`

const types = `{
	"rectangle": [
		{"sideA": "number"},
		{"sideB": "number"}
	]
}`

const imports = `utils "github.com/1backend/go-utils"
example "github.com/1backend/go-example-service"
`

func generateEndpoints(proj *domain.Project) error {
	proj.Description = "An empty Go project"
	proj.Imports = imports
	readme, err := utils.GetReadme(proj)
	if err != nil {
		return err
	}
	proj.ReadMe = readme
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        hi,
			Input:       importedHiInput,
			Output:      importedHiOutput,
			Description: "A very simple endpoint in TypeScript, saying hi to you",
		},
		domain.Endpoint{
			Url:         "/imported-hi",
			Method:      "GET",
			Code:        importedHi,
			Input:       importedHiInput,
			Output:      importedHiOutput,
			Description: "An endpoint that demonstrates the usage of an imported Pluginage",
		},
		domain.Endpoint{
			Url:         "/input-example",
			Method:      "POST",
			Code:        inputExample,
			Input:       inputExampleInput,
			Output:      inputExampleOutput,
			Description: "An endpoint that demonstrates endoint input type usage",
		},
	}
	for _, v := range proj.Dependencies {
		if v.Type == "mysql" {
			proj.Endpoints = append(proj.Endpoints, []domain.Endpoint{
				domain.Endpoint{
					Url:         "/sql-example",
					Method:      "GET",
					Code:        sqlExample,
					Description: "A basic SQL example",
				},
			}...)
		}
	}
	return nil
}
