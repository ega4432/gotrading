package models

import (
	"database/sql"
	"fmt"
	"github.com/ciruclation-dev/gotrading/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const tableNameSignalEvents = "single_events"

var DbConnection *sql.DB

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	DbConnection, err := sql.Open(config.Config.SQLDriver, config.Config.Dbname)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	cmd := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.Productcode, duration)
		log.Println(tableName)
		c := fmt.Sprintf(
			`CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			open FLOAT,
			close FLOAT,
			high FLOAT,
			low open FLOAT,
			volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	}
}
