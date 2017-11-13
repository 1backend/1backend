package main

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	httpr "github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/crufter/1backend/backend/handlers"
	"github.com/crufter/1backend/backend/proxy"
)

func registerHandlers(r *httpr.Router, h *handlers.Handlers, p *proxy.Proxy) {
	r.GET("/v1/project", h.GetProject)
	r.GET("/v1/projects", h.GetProjects)
	r.POST("/v1/project", h.CreateProject)
	r.PUT("/v1/project", h.UpdateProject)
	r.GET("/v1/builds/:projectId", h.GetBuilds)

	r.PUT("/v1/star", h.PutStar)
	r.DELETE("/v1/star", h.DeleteStar)

	r.GET("/v1/issues", h.GetIssues)
	r.GET("/v1/issue", h.GetIssue)
	r.POST("/v1/issue", h.CreateIssue)
	r.PUT("/v1/issue", h.UpdateIssue)
	r.POST("/v1/comment", h.CreateComment)
	r.PUT("/v1/comment", h.UpdateComment)

	r.POST("/v1/run-sql", h.RunSql)

	r.GET("/v1/user", h.GetUser)
	r.PUT("/v1/user", h.UpdateUser)

	r.POST("/v1/register", h.Register)
	r.POST("/v1/login", h.Login)

	r.GET("/app/:author/:app/*route", p.Proxy)
	r.POST("/app/:author/:app/*route", p.Proxy)
	r.PUT("/app/:author/:app/*route", p.Proxy)
	r.DELETE("/app/:author/:app/*route", p.Proxy)
}

func main() {
	var err error
	var db *gorm.DB
	dsn := fmt.Sprintf("%s@tcp(%s)/backend?parseTime=True", "root:root", "127.0.0.1:3306")
	if db, err = gorm.Open("mysql", dsn); err != nil {
		panic(fmt.Sprintf("[init] unable to initialize gorm: %s", err.Error()))
	}
	db = db.Set("gorm:save_associations", false)
	h := handlers.NewHandlers(db)
	defer db.Close()
	r := httpr.New()
	p := proxy.NewProxy(db)
	go p.Nuker()
	registerHandlers(r, h, p)
	handler := cors.New(cors.Options{AllowedHeaders: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}}).Handler(r)
	log.Info("Starting http server")
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%v", 8883), handler))
}
