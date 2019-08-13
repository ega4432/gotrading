package models

import (
	"database/sql"
	"fmt"
	"github.com/ciruclation-dev/gotrading/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const tableNameSingleEvents = "single_events"

var Dbconnection *sql.DB

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	Dbconnection, err := sql.Open(config.Config.SQLDriver, config.Config.Dbname)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
		 CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT)`, tableNameSingleEvents)
	Dbconnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.Productcode, duration)
		c := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY NOT NULL,
			open FLOAT,
			close FLOAT,
			high FLOAT,
			low open FLOAT,
			volume FLOAT)`, tableName)
		Dbconnection.Exec(c)
	}
}
