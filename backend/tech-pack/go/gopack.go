package gopack

import (
	"text/template"

	"github.com/crufter/1backend/backend/domain"
)

func NewPack(project *domain.Project) GoPack {
	return GoPack{
		project: project,
	}
}

type GoPack struct {
	project *domain.Project
}

func (g GoPack) RecipePath() string {
	return "go"
}

func (g GoPack) CreateProjectPlugin() error {
	generateEndpoints(g.project)
	return nil
}

func (g GoPack) Outfile() string {
	return "main.go"
}

func (g GoPack) AddTemplateFuncs(t *template.FuncMap) {

}

func (g GoPack) FilesToBuild() [][]string {
	return nil
}

const goCreateUserSql = `type User struct {
  Id string
  Name string
}

func PostUser(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	input := struct {
		User User
	}{}
	if err := readJsonBody(r, &input); err != nil {
		write400(w, err)
		return
	}
	err := sql.Save(&input.User).Error
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}
`

const goHi = `
var hiCount = 0

func GetHi(w http.ResponseWriter, r *http.Request, p httpr.Params) {
  hiCount++
  write(w, map[string]interface{}{
	    "hi": hiCount,
	})
}
`

const goGetUserSql = `
func GetUser(w http.ResponseWriter, r *http.Request, p httpr.Params) {
  userId := r.URL.Query().Get("id")
  user := &User{}
	err := sql.Where("id = ?", userId).Find(user).Error
	if err != nil {
		write500(w, err)
		return
  }
	write(w, user)
}
`

func generateEndpoints(proj *domain.Project) {
	proj.Endpoints = []domain.Endpoint{
		domain.Endpoint{
			Url:         "/hi",
			Method:      "GET",
			Code:        goHi,
			Description: "A very simple endpoint in go, saying hi to you",
		},
	}
	for _, v := range proj.Dependencies {
		if v.Type == "mysql" {
			proj.Endpoints = append(proj.Endpoints, []domain.Endpoint{
				domain.Endpoint{
					Url:         "/user",
					Method:      "POST",
					Code:        goCreateUserSql,
					Description: "Saves a user to the mysql database",
				},
				domain.Endpoint{
					Url:         "/user",
					Method:      "GET",
					Code:        goGetUserSql,
					Description: "Loads a user from the mysql database",
				},
			}...)
		}
	}
}
