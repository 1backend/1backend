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
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	httpr "github.com/julienschmidt/httprouter"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
)

type Proxy struct {
	db          *gorm.DB
	redisClient *redis.Client
	state       *state.State
}

func NewProxy(db *gorm.DB, redisClient *redis.Client) *Proxy {
	return &Proxy{
		db:    db,
		state: state.NewState(redisClient),
	}
}

func (p *Proxy) launchAndWait(author, projectName string) error {
	output, err := exec.Command("/bin/bash", config.C.Path+"/bash/start-container.sh", author, projectName).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Start container failed: %v | %v ", string(output), err)
	}
	time.Sleep(50 * time.Millisecond)
	var po int64
	for i := 0; i <= 5; i++ {
		time.Sleep(50 * time.Millisecond)
		if err != nil && i == 5 {
			return err
		}
		output, err = exec.Command("/bin/bash", config.C.Path+"/bash/get-port.sh", author, projectName).CombinedOutput()
		if err != nil {
			continue
		}
		po, err = strconv.ParseInt(strings.TrimSpace(string(output)), 10, 64)
		if err != nil {
			continue
		}
		break
	}
	err = p.state.SetPort(author, projectName, int(po))
	if err != nil {
		return err
	}
	for i := 0; i <= 5; i++ {
		if err != nil && i == 5 {
			return err
		}
		time.Sleep(125 * time.Millisecond)
		port := int(po)
		_, err = http.Get(fmt.Sprintf("%s://%s:%v/ping", "http", "127.0.0.1", port))
		if err != nil {
			continue
		}
		// log.Infof("Number of requests it took to wake up service: %v", i)
		break
	}
	return err
}

// This is a horribly overcomplicated abomination
func (p *Proxy) Proxy(w http.ResponseWriter, req *http.Request, params httpr.Params) {
	path := strings.Split(req.RequestURI, "/")
	author := path[2]
	projectName := path[3]
	newPath := strings.Join(path[4:], "/")
	token := req.Header.Get("token")
	if newPath != "ping" && token == "" {
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
	isUp, err := p.state.IsUp(author, projectName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isUp {
		err := p.state.SetLastCall(author, projectName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		underStartup, err := p.state.IsUnderStartup(author, projectName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if underStartup {
			time.Sleep(450 * time.Millisecond)
		} else {
			err := p.state.MarkAsUnderStartup(author, projectName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer p.state.MarkAsNotUnderStartup(author, projectName)
			err = p.launchAndWait(author, projectName)
			p.state.MarkAsNotUnderStartup(author, projectName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			p.state.MarkAsUp(author, projectName)
		}
	}
	p.state.SetLastCall(author, projectName)
	port, err := p.state.Port(author, projectName)
	if err != nil {
		http.Error(w, fmt.Sprintf("port not found: %v", err), http.StatusInternalServerError)
		return
	}
	// create a new url from the raw RequestURI sent by the client
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
	callerNamespace, err := p.state.CallerIdToNameSpace(req.Header.Get("caller-id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	proxyReq.Header.Set("caller-namespace", callerNamespace)
	proxyReq.Header.Set("caller-id", req.Header.Get("caller-id"))
	p.state.Decrement(token)
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

func (p *Proxy) Nuker() {
	projects := []domain.Project{}
	err := p.db.Select("name, author").Find(&projects).Error
	if err != nil {
		log.Error(err)
	}
	for _, v := range projects {
		exec.Command("/bin/bash", config.C.Path+"/bash/stop-container.sh", v.Author, v.Name).CombinedOutput()
		p.state.MarkAsDown(v.Author, v.Name)
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
		}
		for _, v := range projects {
			shallLive, _ := p.state.LastCallIn(v.Author, v.Name, 11*time.Second)
			if shallLive {
				continue
			}
			_, err := exec.Command("/bin/bash", config.C.Path+"/bash/stop-container.sh", v.Author, v.Name).CombinedOutput()
			if err == nil {
				p.state.MarkAsDown(v.Author, v.Name)
			}
		}
		i++
		time.Sleep(10 * time.Second)
	}
}
