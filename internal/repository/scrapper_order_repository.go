package repository

import (
	"log"
	"scrapper-web/config"
)
type ScrapperOrderRepository struct {
	BaseRepository
}


func (r *ScrapperOrderRepository) InsertScrapperOrder(url string) error {
	db := config.GetDB()
	insertSQL := `INSERT INTO scrapperOrder(url) VALUES (?)`
    statement, err := db.Prepare(insertSQL)
    if err != nil {
        log.Fatalln(err)
    }
    defer statement.Close()

    _, err = statement.Exec(url)
    if err != nil {
        log.Fatalln(err)
    }
	if err != nil {
		return err
	}
	return nil
}
