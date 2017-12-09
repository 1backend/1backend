package utils

import (
	"bytes"
	"html/template"

	"github.com/1backend/1backend/backend/domain"
)

const readMe = `{{ .ProjectName }}
===

This is a {{ .Language }} service. If you are new to 1Backend, here is some help:

#### What's a service?

A collection of lambda functions that are accessible through HTTP (also called endpoints).
They speak JSON and have access to certain infrastructure elements you selected (for example: databases).

#### Should I write code in the browser?

No, this web interface is only here to assign functions to paths.
Each lambda should be as short as possible, ideally a single line calling your own package imported in the ` + "`" + "imports" + "`" + ` section.

#### How can I talk to this service?

HTTP is the obvious solution, but there are better ways: for each service type safe clients are generated during builds,
in the following languages (wait for the first build to finish before checkin these):

- Go - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/go/{{ .Project }})
- Angular - [GitHub](https://github.com/{{ .GithubOrganisation }}/{{ .Author }}/ng/{{ .Project }}), [NPM](https://www.npmjs.com/package/@{{ .NpmOrg }}}/{{ .Author }}-{{ .ProjectName }}-ng)
`

func GetReadme(project *domain.Project) (string, error) {
	templ := template.New(project.Name + " readme")
	t, err := templ.Parse(readMe)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, map[string]string{
		"ProjectName": project.Name,
		"Author": project.Author,
		"Language": 
	})
	return buf.String(), err
}

func getLanguage(s string) {
	switch strings.ToLower
}
