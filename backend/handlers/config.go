package handlers

import (
	"errors"
	"net/http"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	httpr "github.com/julienschmidt/httprouter"
)

// We either get sent a nick, or a nick + token
// We get a token when we want to read the user by token
// We get a nick + token when either viewing our own profile or other peoples' profile
func (h *Handlers) GetConfig(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	token := r.URL.Query().Get("token")
	tk, err := domain.NewAccessTokenDao(h.db).GetByToken(token)
	if err != nil {
		write400(w, err)
		return
	}
	user, err := domain.NewUserDao(h.db).GetById(tk.UserId)
	if err != nil {
		write400(w, err)
		return
	}
	if user.Level < 100 {
		write500(w, errors.New("No rights to view config"))
		return
	}
	write(w, config.C)
	return
}

func (h *Handlers) UpdateConfig(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token  string
		Config domain.User
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	tk, err := domain.NewAccessTokenDao(h.db).GetByToken(inp.Token)
	if err != nil {
		write400(w, err)
		return
	}
	user, err := domain.NewUserDao(h.db).GetById(tk.UserId)
	if err != nil {
		write400(w, err)
		return
	}
	if user.Level < 100 {
		write500(w, errors.New("No rights to view config"))
		return
	}

	write(w, map[string]string{})
}
