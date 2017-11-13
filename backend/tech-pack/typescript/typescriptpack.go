package typescriptpack

import (
	"text/template"

	"github.com/crufter/1backend/backend/domain"
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
  "name": "1backend-typescript-example-service",
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
    "@types/express": "^4.0"
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

const nodeHi = `(req: express.Request, rsp: express.Response) => {
  rsp.send('hi');
}
`

const nodeSqlExample = `(req: express.Request, rsp: express.Response) => {
  db.query('SELECT 1 + 1 AS solution', (err: mysql.MysqlError, rows) => {
    rsp.send('The solution is: ' + rows[0]['solution']);
  });
}
`

func generateEndpoints(proj *domain.Project) {
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        nodeHi,
			Description: "A very simple endpoint in go, saying hi to you",
		},
	}
	for _, v := range proj.Dependencies {
		if v.Type == "mysql" {
			proj.Endpoints = append(proj.Endpoints, []domain.Endpoint{
				domain.Endpoint{
					Url:         "/sql-example",
					Method:      "GET",
					Code:        nodeSqlExample,
					Description: "A basic MySQL example",
				},
			}...)
		}
	}
}
