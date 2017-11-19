package gopack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
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
	generateEndpoints(g.project)
	return nil
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

const importedHi = `func (w http.ResponseWriter, r *http.Request, p httpr.Params) {
  example.Hi(w, r, p)
}

// alternatively, replace the above 3 lines with the following single line:
// example.Hi
`

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

const imports = `utils "github.com/1backend/go-utils"
example "github.com/1backend/go-example-service"
`

func generateEndpoints(proj *domain.Project) {
	proj.Imports = imports
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        hi,
			Description: "A very simple endpoint in go, saying hi to you",
		},
		domain.Endpoint{
			Url:         "/imported-hi",
			Method:      "GET",
			Code:        importedHi,
			Description: "An endpoint that demonstrates the usage of an imported package",
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
}
