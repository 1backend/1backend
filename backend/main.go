package main

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	httpr "github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/1backend/1backend/backend/handlers"
	"github.com/1backend/1backend/backend/proxy"
)

func registerHandlers(r *httpr.Router, h *handlers.Handlers, p *proxy.Proxy) {
	r.GET("/v1/project", h.GetProject)
	r.GET("/v1/projects", h.GetProjects)
	r.POST("/v1/project", h.CreateProject)
	r.PUT("/v1/project", h.UpdateProject)

	r.PUT("/v1/star", h.PutStar)
	r.DELETE("/v1/star", h.DeleteStar)

	r.GET("/v1/issues", h.GetIssues)
	r.GET("/v1/issue", h.GetIssue)
	r.POST("/v1/issue", h.CreateIssue)
	r.PUT("/v1/issue", h.UpdateIssue)

	r.POST("/v1/comment", h.CreateComment)
	r.PUT("/v1/comment", h.UpdateComment)

	r.GET("/v1/posts", h.GetPosts)
	r.GET("/v1/post", h.GetPost)
	r.POST("/v1/post", h.CreatePost)
	r.PUT("/v1/post", h.UpdatePost)

	r.POST("/v1/charge", h.Charge)

	r.POST("/v1/token", h.CreateToken)
	r.POST("/v1/token/transfer", h.QuotaTransfer)

	r.POST("/v1/run-sql", h.RunSql)

	r.GET("/v1/user", h.GetUser)
	r.PUT("/v1/user", h.UpdateUser)
	r.POST("/v1/register", h.Register)
	r.POST("/v1/login", h.Login)
	r.POST("/v1/change-password", h.ChangePassword)

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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	h := handlers.NewHandlers(db, redisClient)
	defer db.Close()
	r := httpr.New()
	p := proxy.NewProxy(db, redisClient)
	go p.Nuker()
	registerHandlers(r, h, p)

	handler := cors.New(cors.Options{AllowedHeaders: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}}).Handler(r)
	log.Info("Starting http server")
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%v", 8883), handler))
}

func init() {
	logger, err := log.LoggerFromConfigAsString(seelogConf)
	if err == nil {
		log.ReplaceLogger(logger)
	} else {
		log.Error(err)
	}
}

const seelogConf = `
<seelog>
  <outputs>
    <console formatid="colored"/>
  </outputs>
  <formats>
    <format id="colored"  format="%Date(2006 Jan 02 3:04:05.00 PM MST) (%File:%Line) [%EscM(36)%LEVEL%EscM(39)] %Msg%n%EscM(0)"/>
  </formats>
</seelog>
`
