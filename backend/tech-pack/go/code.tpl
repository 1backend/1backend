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
	gorm *gorm.DB
	db *sql.DB
)

func main() {
	var err error

  	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlIp := os.Getenv("MYSQL_IP")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlDatabase := os.Getenv("MYSQL_DB")

	mysqlAddress := fmt.Sprintf("%s@tcp(%s)/%s?parseTime=True", mysqlUser + ":" + mysqlPassword, mysqlIp + ":3306", mysqlDatabase)
	if gorm, err = gorm.Open("mysql", mysqlAddress); err != nil {
		panic(fmt.Sprintf("[init] unable to initialize gorm: %s", err.Error()))
	}
	defer gormClient.Close()
	db = gormClient.DB()
	
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
	bod, _ := json.Marshal(map[string]bool{
		"pong": true ,
	})
	body := string(bod)
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, `%v`, body)
}