package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     int
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)

	return conf
}
