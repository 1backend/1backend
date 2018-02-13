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

This is a {{ .Mode }} service. If you are new to this, the documentation can be found [here](https://github.com/1backend/1backend/blob/master/docs)

If you know the drill, then happy crufting!
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
