package nodejsplugin

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	techt "github.com/1backend/1backend/backend/tech-plugins/types"
	"github.com/1backend/1backend/backend/tech-plugins/utils"
)

func NewPlugin(project *domain.Project) NodeJsPlugin {
	return NodeJsPlugin{
		project: project,
	}
}

type NodeJsPlugin struct {
	project *domain.Project
}

func (g NodeJsPlugin) PreCreate() error {
	return generateEndpoints(g.project)
}

func (g NodeJsPlugin) Build(t *template.FuncMap) (*techt.Build, error) {
	return &techt.Build{
		Outfile:    "server.js",
		RecipePath: "nodejs",
		FilesBuilt: [][]string{{"package.json", g.project.Packages}},
	}, nil
}

var packageJson = `{
  "name": "1backend-nodejs-service",
  "version": "0.1.0",
  "description": "A sample Node.js app for 1backend",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
		"express": "^4.13.3",
		"mysql": "^2.15.0",
		"@1backend/nodejs-example-service": "^0.0"
  },
  "engines": {
    "node": "4.0.0"
  },
  "license": "MIT"
}`

const hi = `(req, res) => {
  res.send(JSON.stringify('hi'));
}
`
const hiInput = `[]`
const hiOutput = "string"

const importedHi = `(req, res) => {
  service.hi(req, res)
}
`
const importedHiInput = `[]`
const importedHiOutput = "string"

const sqlExample = `(req, res) => {
  db.query('SELECT 1 + 1 AS solution', function (error, results, fields) {
    if (error) throw error;
    const outpout = 'The solution is: ' + rows[0]['solution'];
    rsp.send(JSON.stringify(output));
  });
}
`
const sqlExampleInput = `[]`
const sqlExampleOutput = "string"

const inputExample = `service.calculateRectangleArea`
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

func generateEndpoints(proj *domain.Project) error {
	proj.Description = "An empty Node.js project"
	proj.Imports = `var service = require("@1backend/nodejs-example-service")`
	proj.Packages = packageJson
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
