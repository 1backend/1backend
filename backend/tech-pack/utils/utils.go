package utils

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
)

const readMe = `{{ .ProjectName }}
===

This is a {{ .Mode }} service. If you are new to this, here is some help:

#### What's a service?

A collection of lambda functions that are accessible through HTTP (also called endpoints).
They speak JSON and have access to certain infrastructure elements you selected (for example: databases).

#### Should I write code in the browser?

No, this web interface is only here to assign functions from other packages to paths.
Each lambda should be as short as possible, ideally a single line calling your own package imported in the ` + "`" + "imports" + "`" + ` section.

#### How can I talk to this service?

HTTP is the obvious solution, but there are better ways: for each service type safe clients are generated during builds,
in the following languages (wait for the first build to finish before checkin these):

- Go - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/go/{{ .ProjectName }})
- Angular - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/ng/{{ .ProjectName }}), [NPM](https://www.npmjs.com/package/@{{ .NpmOrg }}}/{{ .Author }}-{{ .ProjectName }}-ng)
`

// GetReadme generates a default readme for different tech-packs.
func GetReadme(project *domain.Project) (string, error) {
	templ := template.New(project.Name + " readme")
	t, err := templ.Parse(readMe)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, map[string]string{
		"ProjectName":        project.Name,
		"Author":             project.Author,
		"Mode":               getMode(project.Mode),
		"GithubOrganisation": config.C.ApiGeneration.GithubOrganisation, // @todo should not reuse this
		"NpmOrg":             config.C.NpmPublication.NpmOrganisation,
	})
	return buf.String(), err
}

func getMode(s string) string {
	switch strings.ToLower(s) {
	case "typescript":
		return "TypeScript"
	case "nodejs":
		return "Node.js"
	case "go", "golang":
		return "Go"
	}
	return s
}
