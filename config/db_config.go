package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `CREATE TABLE IF NOT EXISTS scrapperOrder (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,     
        "url" TEXT
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    return db
}

func InsertData(db *sql.DB, valeur string) {
    insertSQL := `INSERT INTO scrapperOrder(url) VALUES (?)`
    statement, err := db.Prepare(insertSQL)
    if err != nil {
        log.Fatalln(err)
    }
    defer statement.Close()

    _, err = statement.Exec(valeur)
    if err != nil {
        log.Fatalln(err)
    }
}