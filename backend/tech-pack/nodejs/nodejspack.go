package nodejspack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
)

func NewPack(project *domain.Project) NodeJSPack {
	return NodeJSPack{
		project: project,
	}
}

type NodeJSPack struct {
	project *domain.Project
}

func (g NodeJSPack) RecipePath() string {
	return "nodejs"
}

var packageJson = `{
  "name": "1backend-node-example-service",
  "version": "0.1.0",
  "description": "A sample Node.js app for 1backend",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
    "express": "^4.13.3",
    "mysql": "^2.15.0"
  },
  "engines": {
    "node": "4.0.0"
  },
  "license": "MIT"
}`

func (g NodeJSPack) CreateProjectPlugin() error {
	generateEndpoints(g.project)
	return nil
}

func (g NodeJSPack) Outfile() string {
	return "server.js"
}

func (g NodeJSPack) AddTemplateFuncs(t *template.FuncMap) {

}

func (g NodeJSPack) FilesToBuild() [][]string {
	return [][]string{
		[]string{"package.json", packageJson},
	}
}

const hi = `(req, res) => {
  res.send('hi');
}
`

const sqlExample = `(req, res) => {
  db.query('SELECT 1 + 1 AS solution', function (error, results, fields) {
    if (error) throw error;
    res.send('The solution is: ' + results[0].solution);
  });
}
`

func generateEndpoints(proj *domain.Project) {
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        hi,
			Description: "A very simple endpoint in Node.js, saying hi to you",
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
