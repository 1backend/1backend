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
	Path          string
	ApiGeneration struct {
		Enabled             bool // API generation enabled
		GithubOrganisation  string
		GithubUser          string
		GithubPersonalToken string
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
