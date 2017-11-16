package proxy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	httpr "github.com/julienschmidt/httprouter"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
)

type Proxy struct {
	db *gorm.DB
}

func NewProxy(db *gorm.DB) *Proxy {
	return &Proxy{
		db: db,
	}
}

// get port from memory, fall back to database
func (p *Proxy) getPort(author, projectName string) (int, error) {
	port := 0
	port = state.Port(author, projectName)
	if port == 0 {
		project := domain.Project{}
		err := p.db.Select("name, port, author").Where("author = ? AND name = ?", author, projectName).First(&project).Error
		if err != nil {
			return 0, err
		}
		port = project.Port
		state.SetPort(project.Author, project.Name, project.Port)
	}
	return port, nil
}

func (p *Proxy) launchAndWait(author, projectName string) error {
	output, err := exec.Command("/bin/bash", "./bash/start-container.sh", author, projectName).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Start container failed: %v | %v ", string(output), err)
	}
	time.Sleep(100 * time.Millisecond)
	output, err = exec.Command("/bin/bash", "./bash/get-port.sh", author, projectName).CombinedOutput()
	if err != nil {
		return err
	}
	var po int64
	po, err = strconv.ParseInt(strings.TrimSpace(string(output)), 10, 64)
	if err != nil {
		return err
	}
	maxIter := 5
	tts := 125 * time.Millisecond
	for i := 0; i <= 5; i++ {
		if err != nil && i == maxIter {
			return err
		}
		state.SetPort(author, projectName, int(po))
		port := int(po)
		time.Sleep(tts)
		_, err = http.Get(fmt.Sprintf("%s://%s:%v/ping", "http", "127.0.0.1", port))
		if err != nil {
			continue
		}
		// log.Infof("Number of requests it took to wake up service: %v", i)
		break
	}
	return err
}

func (p *Proxy) getUserIdOfToken() {

}

func (p *Proxy) Proxy(w http.ResponseWriter, req *http.Request, params httpr.Params) {
	token := http.Header.Get("token")
	if token == "" {
		http.Error(w, "Please send a token in the header", http.StatusBadRequest)
		return
	}
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	path := strings.Split(req.RequestURI, "/")
	author := path[2]
	projectName := path[3]
	if !state.IsUp(author, projectName) {
		state.SetLastCall(author, projectName)
		if state.IsUnderStartup(author, projectName) {
			time.Sleep(450 * time.Millisecond)
		} else {
			state.MarkAsUnderStartup(author, projectName)
			defer state.MarkAsNotUnderStartup(author, projectName)
			err := p.launchAndWait(author, projectName)
			state.MarkAsNotUnderStartup(author, projectName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			state.MarkAsUp(author, projectName)
		}
	}
	state.SetLastCall(author, projectName)
	port, err := p.getPort(author, projectName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// create a new url from the raw RequestURI sent by the client
	newPath := strings.Join(path[4:], "/")
	url := fmt.Sprintf("%s://%s:%v/%s", "http", "127.0.0.1", port, newPath)
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	// status (must come after setting headers and before copying body)
	w.WriteHeader(resp.StatusCode)
	// body
	io.Copy(w, resp.Body)
}

func (p *Proxy) RequestCountPersistor() {
	users := []domain.User{}
	err := p.db.Find(&users).Preload("tokens").Error
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		state.SetQuota(user.Id, user.Quota)
		for _, token := range user.Tokens {
			state.SetTokenToUser(user.Id, token.Id)
		}
	}
	for {
		time.Sleep(1 * time.Minute)
		// @todo this is obv a very bad design, gotta fix before going distributed
		for userId, decrease := range state.NullAndReturn() {
			err := p.db.Exec("UPDATE users SET quota = ? WHERE id = ?", decrease, userId).Error
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (p *Proxy) Nuker() {
	projects := []domain.Project{}
	err := p.db.Select("name, author").Find(&projects).Error
	if err != nil {
		log.Error(err)
	}
	for _, v := range projects {
		exec.Command("/bin/bash", "./bash/stop-container.sh", v.Author, v.Name).CombinedOutput()
		state.MarkAsDown(v.Author, v.Name)
		//if err != nil {
		// don't log this - a lot of containers do not exists hence it's producing lot of logs
		// log.Error("Error stopping container: %v | %v", string(outp), err)
		//}
	}
	i := 0
	for {
		if i%10 == 0 {
			err := p.db.Select("name, author").Find(&projects).Error
			if err != nil {
				log.Error(err)
			}
			for _, v := range projects {
				if !state.IsUp(v.Author, v.Name) {
					exec.Command("/bin/bash", "./bash/stop-container.sh", v.Author, v.Name).CombinedOutput()
					state.MarkAsDown(v.Author, v.Name)
				}
			}
		}
		for _, v := range projects {
			shallLive := state.LastCallIn(v.Author, v.Name, 11*time.Second)
			if shallLive {
				continue
			}
			_, err := exec.Command("/bin/bash", "./bash/stop-container.sh", v.Author, v.Name).CombinedOutput()
			if err == nil {
				state.MarkAsDown(v.Author, v.Name)
			}
		}
		i++
		if i%10 == 0 {
			err := p.db.Select("name, author").Find(&projects).Error
			if err != nil {
				log.Error(err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
