package main

import (
	"github.com/ciruclation-dev/gotrading/confing"
	"github.com/ciruclation-dev/gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(confing.Config.LogFile)
	log.Println("test")
}
