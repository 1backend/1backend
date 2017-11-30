package config

import (
	"encoding/json"
	"os"

	log "github.com/cihub/seelog"
)

var C = Config{}

type Config struct {
	StripeKey string // stripe api key
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
	// CAUTION! Uses the npm user already logged in on the machine
	NpmPublication struct {
		Enabled bool
		// Org to publish the generated api clients under
		NpmOrganisation string
		// The token used to authenticate for npm publish
		NpmToken string
	}
}

func init() {
	file, err := os.Open("/var/1backend-config.json")
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
