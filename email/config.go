package email

import (
	"log"

	"github.com/tkanos/gonfig"
)

const (
	filePath string = "/Users/Dipesh/Documents/pass/config.json"
)

// Configuration stores the config
type Configuration struct {
	Username string
	Password string
}

// GetConfig exports the configuration
func GetConfig(param ...string) Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(filePath, &configuration)
	if err != nil {
		log.Println(err)
	}
	return configuration
}
