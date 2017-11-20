package typescriptpack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
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
    "@types/mysql": "^2.15.0",
    "@types/express": "^4.0",
    "1backend-typescript-example-service": "^0.0.1"
  },
  "engines": {
    "node": "4.0.0"
  },
  "license": "MIT"
}`

func (g TypeScriptPack) CreateProjectPlugin() error {
	generateEndpoints(g.project)
	return nil
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
  rsp.send('hi');
}
`

const importedHi = `(req: express.Request, rsp: express.Response) => {
  service.Hi(req, rsp)
}
`

const sqlExample = `(req: express.Request, rsp: express.Response) => {
  db.query('SELECT 1 + 1 AS solution', (err: mysql.MysqlError, rows) => {
    rsp.send('The solution is: ' + rows[0]['solution']);
  });
}
`

func generateEndpoints(proj *domain.Project) {
	proj.Imports = `import * as service from '1backend-typescript-example-service';`
	proj.Packages = packageJson
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        hi,
			Description: "A very simple endpoint in TypeScript, saying hi to you",
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
