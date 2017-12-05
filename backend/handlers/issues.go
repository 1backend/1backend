package handlers

import (
	"errors"
	"net/http"

	httpr "github.com/julienschmidt/httprouter"

	"github.com/1backend/1backend/backend/domain"
)

func (h *Handlers) CreateIssue(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
		Issue   domain.Issue
		Comment domain.Comment
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	user, err := h.ep.GetUser(inp.Token)
	if err != nil {
		write400(w, err)
		return
	}
	inp.Issue.UserId = user.Id
	inp.Comment.UserId = user.Id
	err = h.ep.CreateIssue(&inp.Issue)
	if err != nil {
		write400(w, err)
		return
	}
	inp.Comment.IssueId = inp.Issue.Id
	err = h.ep.CreateComment(&inp.Comment)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) UpdateIssue(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token string
		Issue domain.Issue
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	user, err := h.ep.GetUser(inp.Token)
	if err != nil {
		write400(w, err)
		return
	}
	iss := domain.Issue{}
	err = h.db.Where("id = ?", inp.Issue.Id).First(&iss).Error
	if err != nil {
		write400(w, err)
		return
	}
	if user.Id != iss.UserId {
		write400(w, errors.New("No right"))
		return
	}
	err = h.ep.UpdateIssue(&inp.Issue)
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) GetIssues(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	projectId := r.URL.Query().Get("projectId")
	issues := []domain.Issue{}
	err := h.db.Where("project_id = ?", projectId).Preload("User").Find(&issues).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, issues)
}

func (h *Handlers) GetIssue(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	issueId := r.URL.Query().Get("issueId")
	issue := domain.Issue{}
	err := h.db.Where("id = ?", issueId).Preload("User").Preload("Comments").Preload("Comments.User").Find(&issue).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, issue)
}
