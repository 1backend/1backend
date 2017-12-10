package typescriptpack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/tech-pack/utils"
)

func NewPack(project *domain.Project) TypeScriptPack {
	return TypeScriptPack{
		project: project,
	}
}

type TypeScriptPack struct {
	project *domain.Project
}

func (g TypeScriptPack) RecipePath() string {
	return "typescript"
}

// The service itself is not published to NPM, unlike its clients,
// we still generate a package json though.
var packageJson = `{
  "name": "1backend-typescript-service",
  "version": "0.1.0",
  "description": "A sample TypeScript app for 1backend",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
    "express": "^4.13.3",
		"mysql": "^2.15.0",
		"body-parser": "*",
    "@types/mysql": "^2.15.0",
		"@types/express": "^4.0",
		"@types/body-parser": "*",
    "@1backend/typescript-example-service": "^0.0"
  },
  "engines": {
    "node": "4.0.0"
  },
  "license": "MIT"
}`

func (g TypeScriptPack) CreateProjectPlugin() error {
	return generateEndpoints(g.project)
}

func (g TypeScriptPack) Outfile() string {
	return "server.ts"
}

func (g TypeScriptPack) AddTemplateFuncs(t *template.FuncMap) {

}

func (g TypeScriptPack) FilesToBuild() [][]string {
	return [][]string{
		[]string{"package.json", packageJson},
	}
}

const hi = `(req: express.Request, rsp: express.Response) => {
  rsp.send(JSON.stringify('hi'));
}
`
const hiInput = `[]`
const hiOutput = "string"

const importedHi = `(req: express.Request, rsp: express.Response) => {
  service.Hi(req, rsp)
}
`
const importedHiInput = `[]`
const importedHiOutput = "string"

const sqlExample = `(req: express.Request, rsp: express.Response) => {
  sql.query('SELECT 1 + 1 AS solution', (err: mysql.MysqlError, rows) => {
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
	proj.Description = "An empty TypeScript project"
	proj.Imports = `import * as service from '@1backend/typescript-example-service';`
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
