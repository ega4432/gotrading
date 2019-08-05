package main

import (
	"fmt"
	"github.com/ciruclation-dev/gotrading/bitflyer"
	"github.com/ciruclation-dev/gotrading/confing"
	"github.com/ciruclation-dev/gotrading/utils"
)

func main() {
	utils.LoggingSettings(confing.Config.LogFile)
	apiClient := bitflyer.New(confing.Config.ApiKey, confing.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
