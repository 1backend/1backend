package handlers

import (
	"errors"

	"net/http"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
	httpr "github.com/julienschmidt/httprouter"
)

func (h *Handlers) getUser(tokenId string) (*domain.User, error) {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return nil, err
	}
	u, err := domain.NewUserDao(h.db).GetById(token.UserId)
	return &u, err
}

func (h *Handlers) ownsProject(tokenId, projectId string) error {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(h.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	proj := domain.Project{}
	err = h.db.Select("author").Where("id = ?", projectId).First(&proj).Error
	if err != nil {
		return err
	}
	if proj.Author != user.Nick {
		return errors.New("No right to project")
	}
	return nil
}

func (h *Handlers) hasNick(tokenId, author string) error {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(h.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	if user.Nick != author {
		return errors.New("No right to access")
	}
	return nil
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Email    string
		Password string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	user, token, err := h.ep.Login(eb.Email, eb.Password)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// We either get sent a nick, or a nick + token
// We get a token when we want to read the user by token
// We get a nick + token when either viewing our own profile or other peoples' profile
func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	token := r.URL.Query().Get("token")
	nick := r.URL.Query().Get("nick")
	ownErr := h.ep.HasNick(token, nick)
	if nick == "" || ownErr == nil {
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
		st := state.NewState(h.redisClient)
		for i, v := range user.Tokens {
			val, _ := st.GetQuota(v.Token)
			user.Tokens[i].Quota = val
		}
		write(w, user)
		return
	}
	user, err := domain.NewUserDao(h.db).GetByNick(nick)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, user)
	return
}

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token string
		User  domain.User
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
	user.Email = inp.User.Email
	user.AvatarLink = inp.User.AvatarLink
	user.Name = inp.User.Name
	err = domain.NewUserDao(h.db).Update(user)
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) ChangePassword(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token       string
		OldPassword string
		NewPassword string
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
	err = h.ep.ChangePassword(&user, inp.OldPassword, inp.NewPassword)
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Email    string
		Password string
		Nick     string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	user, token, err := h.ep.Register(eb.Email, eb.Nick, eb.Password)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *Handlers) SendResetEmail(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Email string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	err := h.ep.SendResetEmail(eb.Email)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{})
}

func (h *Handlers) ResetPassword(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Secret      string
		NewPassword string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	token, err := h.ep.ResetPassword(eb.Secret, eb.NewPassword)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"token": token,
	})
}
