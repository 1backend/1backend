package gopack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/tech-pack/utils"
)

func NewPack(project *domain.Project) MysqlPack {
	return MysqlPack{
		project: project,
	}
}

type MysqlPack struct {
	project *domain.Project
}

func (g MysqlPack) Name() string {
	return "mysql"
}

func (g MysqlPack) CreateProjectPlugin() error {
	return generateEndpoints(g.project)
}

func (g MysqlPack) AddTemplateFuncs(t *template.FuncMap) {

}

func (g MysqlPack) CodeSections() (string, error) {
	return "", nil
}

const goExample = `func (w http.ResponseWriter, r *http.Request, p httpr.Params) {
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
