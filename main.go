package main

import (
	"fmt"
	"github.com/ciruclation-dev/gotrading/app/models"
	"time"
)

//import (
//	"github.com/ciruclation-dev/gotrading/app/controllers"
//	"github.com/ciruclation-dev/gotrading/config"
//	"github.com/ciruclation-dev/gotrading/utils"
//)
//
//func main() {
//	utils.LoggingSettings(config.Config.LogFile)
//	controllers.StreamIngestionData()
//	controllers.StartWebServer()
//}

func main() {
	s := models.NewSignalEvents()
	df, _ := models.GetAllCandle("BTC_USD", time.Minute, 10)
	c1 := df.Candles[0]
	c2 := df.Candles[5]
	s.Buy("BTC_USD", c1.Time.UTC(), c1.Close, 1.0, true)
	s.Sell("BTC_USD", c2.Time.UTC(), c2.Close, 1.0, true)
	fmt.Println(models.GetSignalEventsByCount(1))
	fmt.Println(models.GetSignalEventsAfterTime(time.Now().UTC()))
	fmt.Println(models.GetSignalEventsAfterTime(c1.Time))
}
