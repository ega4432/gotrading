package models

import (
	"github.com/ciruclation-dev/gotrading/tradingalgo"
	"github.com/markcheno/go-talib"
	"time"
)

type DataFrameCandle struct {
	ProductCode   string         `json:"product_code"`
	Duration      time.Duration  `json:"duration"`
	Candles       []Candle       `json:"candles"`
	Smas          []Sma          `json:"smas,omitempty"`
	Emas          []Ema          `json:"emas,omitempty"`
	BBands        *BBands        `json:"bbands,omitempty"`
	IchimokuCloud *IchimokuCloud `json:"ichimoku,omitempty"`
	Rsi           *Rsi           `json:"rsi,omitempty"`
	Macd          *Macd          `json:"macd,omitempty"`
	Hvs           []Hv           `json:"hvs,omitempty"`
	Events        *SignalEvents  `json:"events,omitempty"`
}

type Sma struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type Ema struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type BBands struct {
	N    int       `json:"n,omitempty"`
	K    float64   `json:"k,omitempty"`
	Up   []float64 `json:"up,omitempty"`
	Mid  []float64 `json:"mid,omitempty"`
	Down []float64 `json:"down,omitempty"`
}

type IchimokuCloud struct {
	Tenkan  []float64 `json:"tenkan,omitempty"`
	Kijun   []float64 `json:"kijun,omitempty"`
	SenkouA []float64 `json:"senkoua,omitempty"`
	SenkouB []float64 `json:"senkoub,omitempty"`
	Chikou  []float64 `json:"chikou,omitempty"`
}

type Rsi struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type Macd struct {
	FastPeriod   int       `json:"fast_period,omitempty"`
	SlowPeriod   int       `json:"slow_period,omitempty"`
	SignalPeriod int       `json:"signal_period,omitempty"`
	Macd         []float64 `json:"macd,omitempty"`
	MacdSignal   []float64 `json:"macd_signal,omitempty"`
	MacdHist     []float64 `json:"macd_hist,omitempty"`
}

type Hv struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

func (df *DataFrameCandle) Times() []time.Time {
	s := make([]time.Time, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Time
	}
	return s
}

func (df *DataFrameCandle) Opens() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Open
	}
	return s
}

func (df *DataFrameCandle) Closes() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Close
	}
	return s
}

func (df *DataFrameCandle) Highs() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.High
	}
	return s
}

func (df *DataFrameCandle) Low() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Low
	}
	return s
}

func (df *DataFrameCandle) Volume() []float64 {
	s := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		s[i] = candle.Volume
	}
	return s
}

func (df *DataFrameCandle) AddSma(period int) bool {
	if len(df.Candles) > period {
		df.Smas = append(df.Smas, Sma{
			Period: period,
			Values: talib.Sma(df.Closes(), period),
		})
		return true
	}
	return false
}

func (df *DataFrameCandle) AddEma(period int) bool {
	if len(df.Candles) > period {
		df.Emas = append(df.Emas, Ema{
			Period: period,
			Values: talib.Ema(df.Closes(), period),
		})
		return true
	}
	return false
}

func (df *DataFrameCandle) AddBBands(n int, k float64) bool {
	if n <= len(df.Closes()) {
		up, mid, down := talib.BBands(df.Closes(), n, k, k, 0)
		df.BBands = &BBands{
			N:    n,
			K:    k,
			Up:   up,
			Mid:  mid,
			Down: down,
		}
		return true
	}
	return true
}

func (df *DataFrameCandle) AddIchimoku() bool {
	tenkanN := 9
	if len(df.Closes()) >= tenkanN {
		tenkan, kijun, senkouA, senkouB, chikou := tradingalgo.IchimokuCloud(df.Closes())
		df.IchimokuCloud = &IchimokuCloud{
			Tenkan:  tenkan,
			Kijun:   kijun,
			SenkouA: senkouA,
			SenkouB: senkouB,
			Chikou:  chikou,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddRsi(period int) bool {
	if len(df.Candles) > period {
		values := talib.Rsi(df.Closes(), period)
		df.Rsi = &Rsi{
			Period: period,
			Values: values,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddMacd(inFastPeriod, inSlowPeriod, inSignalPeriod int) bool {
	if len(df.Candles) > 1 {
		outMACD, outMACDSignal, outMACDHist := talib.Macd(df.Closes(), inFastPeriod, inSlowPeriod, inSignalPeriod)
		df.Macd = &Macd{
			FastPeriod:   inFastPeriod,
			SlowPeriod:   inSlowPeriod,
			SignalPeriod: inSignalPeriod,
			Macd:         outMACD,
			MacdSignal:   outMACDSignal,
			MacdHist:     outMACDHist,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddHv(period int) bool {
	if len(df.Candles) >= period {
		df.Hvs = append(df.Hvs, Hv{
			Period: period,
			Values: tradingalgo.Hv(df.Closes(), period),
		})
		return true
	}
	return false
}

func (df *DataFrameCandle) AddEvents(timeTime time.Time) bool {
	signalEvents := GetSignalEventsAfterTime(timeTime)
	if len(signalEvents.Signals) > 0 {
		df.Events = signalEvents
		return true
	}
	return false
}

func (df *DataFrameCandle) BackTestEma(period1, period2 int) *SignalEvents {
	lenCandles := len(df.Candles)
	if lenCandles <= period1 || lenCandles <= period2 {
		return nil
	}
	signalEvents := NewSignalEvents()
	emaValue1 := talib.Ema(df.Closes(), period1)
	emaValue2 := talib.Ema(df.Closes(), period2)

	for i := 1; i < lenCandles; i++ {
		if i < period1 || i < period2 {
			continue
		}

		if emaValue1[i-1] < emaValue2[i-1] && emaValue1[i] >= emaValue2[i] {
			signalEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}

		if emaValue1[i-1] > emaValue2[i-1] && emaValue1[i] <= emaValue2[i] {
			signalEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
	}
	return signalEvents
}

func (df *DataFrameCandle) OptimizeEma() (performance float64, bestPeriod1 int, bestPeriod2 int) {
	bestPeriod1 = 7
	bestPeriod2 = 14

	for period1 := 5; period1 < 50; period1++ {
		for period2 := 12; period2 < 50; period2++ {
			signalEvents := df.BackTestEma(period1, period2)
			if signalEvents == nil {
				continue
			}
			profit := signalEvents.Profit()
			if performance < profit {
				performance = profit
				bestPeriod1 = period1
				bestPeriod2 = period2
			}
		}
	}
	return performance, bestPeriod1, bestPeriod2
}

func (df *DataFrameCandle) BackTestBb(n int, k float64) *SignalEvents {
	lenCandles := len(df.Candles)

	if lenCandles <= n {
		return nil
	}

	signalEvents := &SignalEvents{}
	bbUp, _, bbDown := talib.BBands(df.Closes(), n, k, k, 0)
	for i := 1; i < lenCandles; i++ {
		if i < n {
			continue
		}
		if bbDown[i-1] > df.Candles[i-1].Close && bbDown[i] <= df.Candles[i].Close {
			signalEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
		if bbUp[i-1] > df.Candles[i-1].Close && bbUp[i] <= df.Candles[i].Close {
			signalEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
	}
	return signalEvents
}

func (df *DataFrameCandle) OptimizeBb() (performance float64, bestN int, bestK float64) {
	bestN = 20
	bestK = 2.0

	for n := 10; n < 60; n++ {
		for k := 1.9; k < 2.1; k += 0.1 {
			signalEvents := df.BackTestBb(n, k)
			if signalEvents == nil {
				continue
			}
			profit := signalEvents.Profit()
			if performance < profit {
				performance = profit
				bestN = n
				bestK = k
			}
		}
	}
	return performance, bestN, bestK
}

func (df *DataFrameCandle) BackTestIchimoku() *SignalEvents {
	lenCandles := len(df.Candles)

	if lenCandles <= 52 {
		return nil
	}

	var signaleEvents SignalEvents
	tenkan, kijun, senkouA, senkouB, chikou := tradingalgo.IchimokuCloud(df.Closes())

	for i := 1; i < lenCandles; i++ {
		if chikou[i-1] < df.Candles[i-1].High && chikou[i] >= df.Candles[i].High &&
			senkouA[i] < df.Candles[i].Low && senkouB[i] < df.Candles[i].Low &&
			tenkan[i] > kijun[i] {
			signaleEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}

		if chikou[i-1] > df.Candles[i-1].Low && chikou[i] <= df.Candles[i].Low &&
			senkouA[i] > df.Candles[i].High && senkouB[i] > df.Candles[i].High &&
			tenkan[i] < kijun[i] {
			signaleEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
	}
	return &signaleEvents
}

func (df *DataFrameCandle) OptimizeIchimoku() (performance float64) {
	signalEvents := df.BackTestIchimoku()
	if signalEvents == nil {
		return 0.0
	}
	performance = signalEvents.Profit()
	return performance
}

func (df *DataFrameCandle) BackTestMacd(macdFastPerio, macdSlowPeriod, macdSignalPeriod int) *SignalEvents {
	lenCandles := len(df.Candles)

	if lenCandles <= macdFastPerio || lenCandles <= macdSlowPeriod || lenCandles <= macdSignalPeriod {
		return nil
	}

	signaleEvents := &SignalEvents{}
	outMACD, outMACDSignal, _ := talib.Macd(df.Closes(), macdFastPerio, macdSlowPeriod, macdSignalPeriod)

	for i := 1; i < lenCandles; i++ {
		if outMACD[i] < 0 &&
			outMACDSignal[i] < 0 &&
			outMACD[i-1] < outMACDSignal[i-1] &&
			outMACD[i] >= outMACDSignal[i] {
			signaleEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}

		if outMACD[i] > 0 &&
			outMACDSignal[i] > 0 &&
			outMACD[i-1] > outMACDSignal[i-1] &&
			outMACD[i] <= outMACDSignal[i] {
			signaleEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
	}
	return signaleEvents
}

func (df *DataFrameCandle) OptimizeMacd() (performance float64, bestMacdFastPeriod, bestMacdSlowPeriod, bestMacdSignalePeriod int) {
	bestMacdFastPeriod = 12
	bestMacdSlowPeriod = 26
	bestMacdSignalePeriod = 9

	for fastPeriod := 10; fastPeriod < 19; fastPeriod++ {
		for slowPeriod := 20; slowPeriod < 30; slowPeriod++ {
			for signalPeriod := 5; signalPeriod < 15; signalPeriod++ {
				signaleEvents := df.BackTestMacd(bestMacdFastPeriod, bestMacdSlowPeriod, bestMacdSignalePeriod)
				if signaleEvents == nil {
					continue
				}
				profit := signaleEvents.Profit()
				if performance < profit {
					performance = profit
					bestMacdFastPeriod = fastPeriod
					bestMacdSlowPeriod = slowPeriod
					bestMacdSignalePeriod = signalPeriod
				}
			}
		}
	}
	return performance, bestMacdFastPeriod, bestMacdSlowPeriod, bestMacdSignalePeriod
}

func (df *DataFrameCandle) BackTestRsi(period int, buyThread, sellThread float64) *SignalEvents {
	lenCandles := len(df.Candles)
	if lenCandles <= period {
		return nil
	}

	signalEvents := NewSignalEvents()
	values := talib.Rsi(df.Closes(), period)
	for i := 1; i < lenCandles; i++ {
		if values[i-1] == 0 || values[i-1] == 100 {
			continue
		}
		if values[i-1] < buyThread && values[i] >= buyThread {
			signalEvents.Buy(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
		if values[i-1] > sellThread && values[i] <= sellThread {
			signalEvents.Sell(df.ProductCode, df.Candles[i].Time, df.Candles[i].Close, 1.0, false)
		}
	}
	return signalEvents
}

func (df *DataFrameCandle) OptimizeRsi() (performance float64, bestPeriod int, bestBuyThread, bestSellThread float64) {
	bestPeriod = 14
	bestBuyThread, bestSellThread = 30.0, 70.0

	for period := 5; period < 25; period++ {
		signaleEvents := df.BackTestRsi(period, bestBuyThread, bestSellThread)
		if signaleEvents == nil {
			continue
		}
		profit := signaleEvents.Profit()
		if performance < profit {
			performance = profit
			bestPeriod = period
			bestBuyThread = bestBuyThread
			bestSellThread = bestSellThread
		}
	}
	return performance, bestPeriod, bestBuyThread, bestSellThread
}
