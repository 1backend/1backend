package config

import (
	"encoding/json"
	"os"

	log "github.com/cihub/seelog"
)

var C = Config{
	StripeKey: "sk_test_BQokikJOvBiI2HlWgH4olfQ2",
}

type Config struct {
	StripeKey string // stripe api key
}

func init() {
	file, err := os.Open("config.json")
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
