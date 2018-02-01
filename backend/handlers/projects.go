package handlers

import (
	"net/http"
	"os/exec"

	log "github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	httpr "github.com/julienschmidt/httprouter"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/endpoints"
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

func (h *Handlers) GetLogs(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	token := r.URL.Query().Get("token")
	projectId := r.URL.Query().Get("projectId")
	if err := h.ep.OwnsProject(token, projectId); err != nil {
		write400(w, err)
		return
	}
	project := domain.Project{}
	err := h.db.Where("id = ?", projectId).Find(&project).Error
	if err != nil {
		write400(w, err)
		return
	}
	output, _ := exec.Command("/bin/bash", config.C.Path+"/bash/logs.sh", project.Author, project.Name).CombinedOutput()
	write(w, map[string]interface{}{
		"Logs": string(output),
	})
}

func (h *Handlers) GetCallerId(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	token := r.URL.Query().Get("token")
	projectId := r.URL.Query().Get("projectId")
	if err := h.ep.OwnsProject(token, projectId); err != nil {
		write400(w, err)
		return
	}
	project := domain.Project{}
	err := h.db.Where("id = ?", projectId).Find(&project).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"CallerId": project.CallerId,
	})
}

func (h *Handlers) GetProject(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	author := r.URL.Query().Get("author")
	projectName := r.URL.Query().Get("project")
	token := r.URL.Query().Get("token")
	project := domain.Project{}
	err := h.db.Where("author = ? AND name = ?", author, projectName).Preload("Starrers").Preload("Dependencies").Preload("Builds").Preload("Builds.Steps").Preload("Endpoints").Find(&project).Error
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
	db = db.Select("author, name, id, mode, stars, public, open_source, description, recipe, created_at, updated_at").Order("stars desc").Limit(500)
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

func (h *Handlers) DeleteProject(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	token := r.URL.Query().Get("token")
	projectId := r.URL.Query().Get("projectId")
	if err := h.ep.OwnsProject(token, projectId); err != nil {
		write400(w, err)
		return
	}
	err := h.ep.DeleteProject(projectId)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}
