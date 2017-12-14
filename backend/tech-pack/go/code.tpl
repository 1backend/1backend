package main

import (
  {{.Project.Imports}}
	"fmt"
	"net/http"
	"os"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"database/sql"
	httpr "github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/cihub/seelog"
)

var (
	gormClient *gorm.DB
	sqlClient *sql.DB
)

func main() {
	var err error
  	infraPass := os.Getenv("INFRAPASS")
	
	ip := os.Getenv("MYSQLIP")
	sqlUser := os.Getenv("MYSQLUSER")
	sqlDatabase := os.Getenv("MYSQLDATABASE")
	sqlAddress := fmt.Sprintf("%s@tcp(%s)/%s?parseTime=True", sqlUser + ":" + infraPass, ip + ":3306", sqlDatabase)
	if gormClient, err = gorm.Open("mysql", sqlAddress); err != nil {
		panic(fmt.Sprintf("[init] unable to initialize gorm: %s", err.Error()))
	}
	defer gormClient.Close()
	sqlClient = gormClient.DB()
	
	r := httpr.New()
	endpoints(r)
	handler := cors.New(cors.Options{AllowedHeaders: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}}).Handler(r)
	log.Info("Starting http server")
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%v", 8883), handler))
}

func endpoints(r *httpr.Router) {
	{{ range $key, $value := .Project.Endpoints }}
	r.{{ $value.Method }}("{{ $value.Url }}", {{ $value.Code | trim }})
	{{ end }}
	r.GET("/ping", GetPing)
}

func GetPing(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	body, _ := json.Marshal(map[string]bool{
		"pong": true ,
	})
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, `%v`, body)
}