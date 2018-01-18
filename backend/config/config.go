package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/cihub/seelog"
)

func IsTestUser(userName string) bool {
	return len(userName) == 17 && strings.HasPrefix(userName, "user-")
}

var C = Config{}

type Config struct {
	SiteUrl     string // eg. https://1backend.com
	StripeKey   string // stripe api key
	SendGridKey string
	// absolute path to folder containing files (assumes same structure as the repo)
	Path string
	// CAUTION! Uses the git user configured on the machine.
	ApiGeneration struct {
		Enabled            bool // API generation enabled
		GithubOrganisation string
		// user and personal token is used for repo creation when calling GitHub's HTTP API
		GithubUser          string
		GithubPersonalToken string
	}
	// Generated ts, node and ng API packages can be published to npmjs.org
	NpmPublication struct {
		Enabled bool
		// Org to publish the generated api clients under
		NpmOrganisation string
		// The token used to authenticate for npm publish
		NpmToken string
	}
	Sitemap struct {
		Enabled bool
		// defaults to /var/sitemap.xml.gz
		Path string
	}
}

func init() {
	load()
}

const filePath = "/var/1backend-config.json"

func load() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&C)
	if err != nil {
		log.Error(err)
	}
}

func Save(nu Config) error {
	dat, err := json.Marshal(nu)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, dat, 0644)
}
