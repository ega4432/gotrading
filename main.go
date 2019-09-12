package main

import (
	"github.com/ciruclation-dev/gotrading/app/controllers"
	"github.com/ciruclation-dev/gotrading/config"
	"github.com/ciruclation-dev/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
