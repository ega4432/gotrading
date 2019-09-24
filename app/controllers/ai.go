package controllers

import (
	"github.com/ciruclation-dev/gotrading/app/models"
	"github.com/ciruclation-dev/gotrading/bitflyer"
	"github.com/ciruclation-dev/gotrading/config"
	"github.com/ciruclation-dev/gotrading/tradingalgo"
	"github.com/markcheno/go-talib"
	"golang.org/x/sync/semaphore"
	"log"
	"strings"
	"time"
)

type AI struct {
	API                  *bitflyer.APIClient
	ProductCode          string
	CurrencyCode         string
	CoinCode             string
	UsePercent           float64
	MinuteToExpires      int
	Duration             time.Duration
	PastPeriod           int
	SignalEvents         *models.SignalEvents
	OptimizedTradeParams *models.TradeParams
	TradeSemaphore       *semaphore.Weighted
	StopLimit            float64
	StopLimitPercent     float64
	BackTest             bool
	StartTrade           time.Time
}

func NewAI(productCode string, duration time.Duration, pastPeriod int, UsePercent, stopLimitPercent float64, backTest bool) *AI {
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	var signalEvents *models.SignalEvents
	if backTest {
		signalEvents = models.NewSignalEvents()
	} else {
		signalEvents = models.GetSignalEventsByCount(1)
	}
	codes := strings.Split(productCode, "_")
	ai := &AI{
		API:              apiClient,
		ProductCode:      productCode,
		CoinCode:         codes[0],
		CurrencyCode:     codes[1],
		UsePercent:       UsePercent,
		MinuteToExpires:  1,
		PastPeriod:       pastPeriod,
		Duration:         duration,
		SignalEvents:     signalEvents,
		TradeSemaphore:   semaphore.NewWeighted(1),
		BackTest:         backTest,
		StartTrade:       time.Now().UTC(),
		StopLimitPercent: stopLimitPercent,
	}
	ai.UpdateOptimizeParams()
	return ai
}

func (ai *AI) UpdateOptimizeParams() {
	df, _ := models.GetAllCandle(ai.ProductCode, ai.Duration, ai.PastPeriod)
	ai.OptimizedTradeParams = df.OptimizeParams()
}

func (ai *AI) Buy(candle models.Candle) (childOrderAcceptanceID string, isOrderComleted bool) {
	if ai.BackTest {
		couldBuy := ai.SignalEvents.Buy(ai.ProductCode, candle.Time, candle.Close, 1.0, false)
		return "", couldBuy
	}
	//TODO
	return childOrderAcceptanceID, isOrderComleted
}

func (ai *AI) Sell(candle models.Candle) (childOrderAcceptanceID string, isOrderComleted bool) {
	if ai.BackTest {
		couldBuy := ai.SignalEvents.Sell(ai.ProductCode, candle.Time, candle.Close, 1.0, false)
		return "", couldBuy
	}
	//TODO
	return childOrderAcceptanceID, isOrderComleted
}

func (ai *AI) Trade() {
	isAcuire := ai.TradeSemaphore.TryAcquire(1)
	if !isAcuire {
		log.Println("Could not get trade lock")
		return
	}
	defer ai.TradeSemaphore.Release(1)
	params := ai.OptimizedTradeParams
	df, _ := models.GetAllCandle(ai.ProductCode, ai.Duration, ai.PastPeriod)
	lenCandles := len(df.Candles)

	var emaValues1 []float64
	var emaValues2 []float64
	if params.EmaEnable {
		emaValues1 = talib.Ema(df.Closes(), params.EmaPeriod1)
		emaValues2 = talib.Ema(df.Closes(), params.EmaPeriod2)
	}

	var bbUp []float64
	var bbDown []float64
	if params.BbEnable {
		bbUp, _, bbDown = talib.BBands(df.Closes(), params.BbN, params.BbK, params.BbK, 0)
	}

	var tenkan, kijun, senkouA, senkouB, chikou []float64
	if params.IchimokuEnable {
		tenkan, kijun, senkouA, senkouB, chikou = tradingalgo.IchimokuCloud(df.Closes())
	}

	var outMACD, outMACDSignal []float64
	if params.MacdEnable {
		outMACD, outMACDSignal, _ = talib.Macd(df.Closes(), params.MacdFastPeriod, params.MacdSlowPeriod, params.MacdSignalPeriod)
	}

	var rsiValues []float64
	if params.RsiEnable {
		rsiValues = talib.Rsi(df.Closes(), params.RsiPeriod)
	}
	for i := 1; i < lenCandles; i++ {

	}

}
