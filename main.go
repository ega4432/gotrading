package main

import (
	"fmt"
	"github.com/ciruclation-dev/gotrading/app/models"
	//"github.com/ciruclation-dev/gotrading/bitflyer"
	"github.com/ciruclation-dev/gotrading/config"
	"github.com/ciruclation-dev/gotrading/utils"
	//"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	//apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	//ticker, _ := apiClient.GetTicker("BTC_USD")
	//fmt.Println(ticker)
	//fmt.Println(ticker.GetMidPrice())
	//fmt.Println(ticker.DateTime())
	//fmt.Println(ticker.TruncateDateTime(time.Hour))
	fmt.Println(models.Dbconnection)
}
