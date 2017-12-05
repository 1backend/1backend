package handlers

import (
	"errors"
	"net/http"

	httpr "github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
)

func (h *Handlers) CreateToken(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token                   string
		ServiceTokenName        string
		ServiceTokenDescription string
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
	token := &domain.Token{
		Token:       uuid.NewV4().String(),
		Name:        inp.ServiceTokenName,
		Description: inp.ServiceTokenDescription,
		UserId:      user.Id,
	}
	err = h.ep.CreateToken(token)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) QuotaTransfer(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token  string // access token
		From   string // from token
		To     string // to token
		Amount int64
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
	tokens := []domain.Token{}
	err = h.db.Where("token in (?)", []string{inp.From, inp.To}).Find(&tokens).Error
	if err != nil {
		write400(w, err)
		return
	}
	for _, v := range tokens {
		if v.UserId != user.Id {
			write400(w, errors.New("No right"))
			return
		}
	}
	st := state.NewState(h.redisClient)
	err = st.DecrementBy(inp.From, inp.Amount)
	if err != nil {
		write500(w, err)
		return
	}
	err = st.IncrementBy(inp.To, inp.Amount)
	// @todo rollback or at least save into db to recover manually
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}
