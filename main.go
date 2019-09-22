package main

import (
	"github.com/ciruclation-dev/gotrading/app/controllers"
	"github.com/ciruclation-dev/gotrading/config"
	"github.com/ciruclation-dev/gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
