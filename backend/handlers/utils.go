package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"
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
