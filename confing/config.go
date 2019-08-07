package confing

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Fail to read file : %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("secret_key").String(),
		LogFile:   cfg.Section("gotrading").Key("log_file").String(),
	}

}