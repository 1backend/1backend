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

var InternalIp = os.Getenv("INTERNAL_IP")

var C = Config{
	MySQL: MySQL{
		Ip:       "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "root",
	},
	MySQLPlugin: MySQLPlugin{
		Ip:   "127.0.0.1",
		Port: "3306",
	},
	Redis: Redis{
		Ip:       "127.0.0.1",
		Port:     "6379",
		Password: "",
		Db:       0,
	},
}

type NpmPublication struct {
	Enabled bool
	// Org to publish the generated api clients under
	NpmOrganisation string
	// The token used to authenticate for npm publish
	NpmToken string
}

type Sitemap struct {
	Enabled bool
	// defaults to /var/sitemap.xml.gz
	Path string
}

type ApiGeneration struct {
	Enabled            bool // API generation enabled
	GithubOrganisation string
	// user and personal token is used for repo creation when calling GitHub's HTTP API
	GithubUser          string
	GithubPersonalToken string
}

type MySQL struct {
	Ip       string
	Port     string
	Username string
	Password string
}

type Redis struct {
	Ip       string
	Port     string
	Password string
	Db       int
}

type MySQLPlugin struct {
	Ip   string
	Port string
}

type Config struct {
	SiteUrl     string // eg. https://1backend.com
	StripeKey   string // stripe api key
	SendGridKey string
	// absolute path to folder containing files (assumes same structure as the repo)
	Path string
	// CAUTION! Uses the git user configured on the machine.
	ApiGeneration ApiGeneration
	// Generated ts, node and ng API packages can be published to npmjs.org
	NpmPublication NpmPublication
	Sitemap        Sitemap
	// The mysql instance the 1backend server connects to
	// Can be different from the mysql instance the containers running on 1backend connect to
	MySQL       MySQL
	Redis       Redis
	MySQLPlugin MySQLPlugin
}

func init() {
	err := load()
	if err != nil {
		log.Error(err)
	}
}

const filePath = "/var/1backend-config.json"

func load() error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&C)
}

func Save(nu Config) error {
	dat, err := json.MarshalIndent(nu, "", "   ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, dat, 0644)
	if err != nil {
		return err
	}
	return load()
}
