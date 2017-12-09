package gopack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/tech-pack/utils"
)

func NewPack(project *domain.Project) GoPack {
	return GoPack{
		project: project,
	}
}

type GoPack struct {
	project *domain.Project
}

func (g GoPack) RecipePath() string {
	return "go"
}

func (g GoPack) CreateProjectPlugin() error {
	return generateEndpoints(g.project)
}

func (g GoPack) Outfile() string {
	return "main.go"
}

func (g GoPack) AddTemplateFuncs(t *template.FuncMap) {

}

func (g GoPack) FilesToBuild() [][]string {
	return nil
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
	rows.Scan(&solution)
	utils.Write(w, solution)
}
`
const sqlExampleInput = `[]`
const sqlExampleOutput = "string"

const inputExample = `service.CalculateRectangleArea`
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
			Description: "An endpoint that demonstrates the usage of an imported package",
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
