package models

import (
	"fmt"
	"github.com/ciruclation-dev/gotrading/config"
	"log"
	"strings"
	"time"
)

type SignalEvent struct {
	Time        time.Time `json:"time"`
	ProductCode string    `json:"product_code"`
	Side        string    `json:"side"`
	Price       float64   `json:"price"`
	Size        float64   `json:"size"`
}

func (s *SignalEvent) Save() bool {
	cmd := fmt.Sprintf("INSERT INTO %s (time, product_code, side, price, size) VALUES(?, ?, ?, ?, ?)", tableNameSignalEvents)
	_, err := DbConnection.Exec(cmd, s.Time.Format(time.RFC3339), s.ProductCode, s.Price, s.Size)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err)
			return true

		}
		return false
	}
	return true
}

type SignalEvents struct {
	Signals []SignalEvent `json:"signals,omitempty"`
}

func NewSignalEvents() *SignalEvents {
	return &SignalEvents{}
}

func GetSignalEventsByCount(loadEvents int) *SignalEvents {
	cmd := fmt.Sprintf(`SELECT * FROM (
		SELECT time, product_code, side, price, size FROM %s WHERE product_code = ? ORDER BY time DESC LIMIT ? )
		ORDER BY time ASC;`, tableNameSignalEvents)
	rows, err := DbConnection.Query(cmd, config.Config.ProductCode, loadEvents)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents SignalEvents
	for rows.Next() {
		var signalEvent SignalEvent
		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return &signalEvents
}

func GetSignalEventsAfterTime(timeTime time.Time) *SignalEvents {
	cmd := fmt.Sprintf(`SELECT FROM (
		SELECT time, product_codem, side, price, size FROM %s
		WHERE DATETIME(time) >= DATETIME(?)
		ORDER BY time DESC
		) ORDER BY time ASC:`, tableNameSignalEvents)
	rows, err := DbConnection.Query(cmd, timeTime.Format(time.RFC3339))
	if err != nil {
		return nil
	}
	defer rows.Close()

	var signalEvents SignalEvents
	for rows.Next() {
		var signalEvent SignalEvent
		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
		signalEvents.Signals= append(signalEvents.Signals, signalEvent)
	}
	return &signalEvents
}