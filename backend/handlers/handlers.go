package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	log "github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	httpr "github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/endpoints"
	"github.com/1backend/1backend/backend/state"
)

type Handlers struct {
	db          *gorm.DB
	ep          *endpoints.Endpoints
	redisClient *redis.Client
}

func NewHandlers(db *gorm.DB, rc *redis.Client) *Handlers {
	return &Handlers{
		db:          db,
		ep:          endpoints.NewEndpoints(db, rc),
		redisClient: rc,
	}
}

func write(w http.ResponseWriter, body interface{}) {
	rawBody, err := json.Marshal(body)
	if err != nil {
		write500(w, err)
		return
	}
	writeString(w, 200, string(rawBody))
}

func writeString(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `%v`, body)
}

func readJsonBody(r *http.Request, expectedBody interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, expectedBody); err != nil {
		log.Errorf(
			"[readJsonBody] invalid request, %s, input was %s",
			err.Error(), string(body))
		return errors.New("invalid json body format: " + err.Error())
	}
	return nil
}

func write400(w http.ResponseWriter, err error) {
	rawBody, err := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if err != nil {
		write500(w, err)
		return
	}
	writeString(w, 400, string(rawBody))
}

func write500(w http.ResponseWriter, err error) {
	rawBody, err := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if err != nil {
		write500(w, err)
		return
	}
	writeString(w, 500, string(rawBody))
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

func (h *Handlers) RunSql(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token     string
		ProjectId string
		Sql       string
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	if err := h.ep.OwnsProject(inp.Token, inp.ProjectId); err != nil {
		write400(w, err)
		return
	}
	project := domain.Project{}
	err := h.db.Where("id = ?", inp.ProjectId).Find(&project).Error
	if err != nil {
		write400(w, err)
		return
	}
	output, _ := exec.Command("/bin/bash", config.C.Path+"/bash/run-sql.sh", inp.Sql, project.Author, project.Name).CombinedOutput()
	write(w, map[string]interface{}{
		"Answer": string(output),
	})
}

func (h *Handlers) GetProject(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	author := r.URL.Query().Get("author")
	projectName := r.URL.Query().Get("project")
	token := r.URL.Query().Get("token")
	project := domain.Project{}
	err := h.db.Where("author = ? AND name = ?", author, projectName).Preload("Starrers").Preload("Dependencies").Preload("Builds").Preload("Endpoints").Find(&project).Error
	if err != nil {
		write400(w, err)
		return
	}
	owns := h.ep.OwnsProject(token, project.Id)
	if !project.Public && owns != nil {
		write400(w, owns)
		return
	}
	if !project.OpenSource && owns != nil {
		for _, v := range project.Endpoints {
			v.Code = ""
		}
	}
	write(w, project)
}

func (h *Handlers) GetProjects(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	nick := r.URL.Query().Get("nick")
	token := r.URL.Query().Get("token")
	projects := []domain.Project{}
	db := h.db
	own := false
	if err := h.ep.HasNick(token, nick); nick != "" && err == nil {
		own = true
	}
	getStarrersFor := ""
	if token != "" {
		tk, err := domain.NewAccessTokenDao(h.db).GetByToken(token)
		if err != nil {
			log.Error(err)
		}
		getStarrersFor = tk.UserId
	}
	if len(nick) == 0 {
		db = db.Where("public = true")
	} else if own {
		db = db.Where("author = ?", nick)
	} else {
		db = db.Where("author = ? AND public = true", nick)
	}
	db = db.Select("author, name, id, stars, public, open_source, description, recipe, created_at, updated_at")
	if getStarrersFor != "" {
		db = db.Preload("Starrers", "user_id = ?", getStarrersFor)
	}
	err := db.Find(&projects).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, projects)
}

func (h *Handlers) GetBuilds(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	builds := []domain.Build{}
	err := h.db.Where("project_id = ?", p.ByName("projectId")).Find(&builds).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, builds)
}

func (h *Handlers) CreateProject(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
		Project domain.Project
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	if err := h.ep.HasNick(inp.Token, inp.Project.Author); err != nil {
		write400(w, err)
		return
	}
	err := h.ep.CreateProject(&inp.Project)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) UpdateProject(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
		Project domain.Project
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	if err := h.ep.OwnsProject(inp.Token, inp.Project.Id); err != nil {
		write400(w, err)
		return
	}
	err := h.ep.UpdateProject(&inp.Project)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

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

func (h *Handlers) CreateComment(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
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
	inp.Comment.UserId = user.Id
	err = h.ep.CreateComment(&inp.Comment)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) UpdateComment(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
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
	comm := domain.Comment{}
	err = h.db.Where("id = ?", inp.Comment.Id).First(&comm).Error
	if err != nil {
		write400(w, err)
		return
	}
	if user.Id != comm.UserId {
		write400(w, errors.New("No right"))
		return
	}
	err = h.ep.UpdateComment(&inp.Comment)
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

func (h *Handlers) PutStar(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token     string
		ProjectId string
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
	err = h.ep.PutStar(user.Id, inp.ProjectId)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) DeleteStar(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	user, err := h.ep.GetUser(r.URL.Query().Get("token"))
	if err != nil {
		write400(w, err)
		return
	}
	err = h.ep.DeleteStar(user.Id, r.URL.Query().Get("projectId"))
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

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

func (h *Handlers) Charge(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token        string // access token
		Amount       int64
		PaymentToken string
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	err := h.ep.Charge(inp.Token, inp.PaymentToken, uint64(inp.Amount))
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}
