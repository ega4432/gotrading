package main

import (
	"fmt"
	"github.com/ciruclation-dev/gotrading/bitflyer"
	"github.com/ciruclation-dev/gotrading/confing"
	"github.com/ciruclation-dev/gotrading/utils"
	"time"
)

func main() {
	utils.LoggingSettings(confing.Config.LogFile)
	apiClient := bitflyer.New(confing.Config.ApiKey, confing.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
