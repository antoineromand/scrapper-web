package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
    instance *sql.DB
    once     sync.Once
)

func InitDB(driverName string, dataSourceName string) *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open(driverName, dataSourceName)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		createTableSQL := `CREATE TABLE IF NOT EXISTS scrapperOrder (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,     
			"url" TEXT
		);`
		_, err = instance.Exec(createTableSQL)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	})

	return instance
}

func GetDB() *sql.DB {
	if instance == nil {
		log.Fatal("Database not initialized")
	}
	return instance
}
