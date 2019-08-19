package controllers

import (
	"github.com/ciruclation-dev/gotrading/app/models"
	"github.com/ciruclation-dev/gotrading/bitflyer"
	"github.com/ciruclation-dev/gotrading/config"
	"log"
)

func StreamIngestionData() {
	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	go func() {
		for ticker := range tickerChannl {
			log.Println("action=StreamIngestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.Config.TradeDuration {
					// Todo
				}
			}
		}
	}()

}
