package utils

import (
	"bytes"
	"strings"
	"text/template"

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
Each lambda should be as short as possible, ideally a single line calling your own package imported in the "imports" section.

#### How can I talk to this service?

HTTP is the obvious solution, but there are better ways: for each service type safe clients are generated during builds,
in the following languages (wait for the first build to finish before checking these):

- Go - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/go/{{ .ProjectName }})
- Angular - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/ng/{{ .ProjectName }}), [NPM](https://www.npmjs.com/package/@{{ .NpmOrganisation }}}/{{ .Author }}-{{ .ProjectName }}-ng)

#### What are these "types", "input", "output" things?

With the help of a simple language agnostic DSL you can define the following things:
- In the "types" section you can define the types (examples: User, Customer, Product etc.) of your service 
- In the "input" and "output" section of endpoints you can define parameters for your endpoint.
You can refer to types defined in the mentioned "types" section of your service, or in other services.

You can read more about types and the DSL [here](https://github.com/1backend/1backend/blob/master/docs/types.md)

The complete documentation can be found [here](https://github.com/1backend/1backend/blob/master/docs)
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
		"NpmOrganisation":    config.C.NpmPublication.NpmOrganisation,
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
