package main

import (
  {{.Project.Imports}}
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/cihub/seelog"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	httpr "github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

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

var (
	sql *gorm.DB
)

func main() {
	// init mysql
	var err error
	ip := os.Getenv("MYSQLIP")
  infraPass := os.Getenv("INFRAPASS")
	mysqlAddress := fmt.Sprintf("%s@tcp(%s)/{{ .Project.Author }}_{{ .Project.Name }}?parseTime=True", "{{ .Project.Author }}_{{ .Project.Name }}:" + infraPass, ip + ":3306") // :t8ecNpCf5u0d
	if sql, err = gorm.Open("mysql", mysqlAddress); err != nil {
		panic(fmt.Sprintf("[init] unable to initialize gorm: %s", err.Error()))
	}
	sql = sql.Set("gorm:save_associations", false)
	defer sql.Close()
	r := httpr.New()
	endpoints(r)
	handler := cors.New(cors.Options{AllowedHeaders: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}}).Handler(r)
	log.Info("Starting http server")
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%v", 8883), handler))
}

func endpoints(r *httpr.Router) {
	{{ range $key, $value := .Project.Endpoints }}
		r.{{ $value.Method }}("{{ $value.Url }}", {{ getEndpointName $value.Method $value.Url }})
	{{ end }}
		r.GET("/ping", GetPing)
}

func GetPing(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	write(w, map[string]bool{
	    "pong": true,
	})
}

{{ range $key, $value := .Project.Endpoints }}
	{{ $value.Code }}
{{ end }}
