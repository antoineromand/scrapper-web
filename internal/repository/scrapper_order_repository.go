package repository

import (
	"database/sql"
	"log"
	Models "scrapper-web/internal/model"
)
type ScrapperOrderRepository struct {
	DB *sql.DB
}


func (r *ScrapperOrderRepository) InsertScrapperOrder(url string) error {
	insertSQL := `INSERT INTO scrapperOrder(url) VALUES (?)`
    statement, err := r.DB.Prepare(insertSQL)
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

func (r *ScrapperOrderRepository) GetScrapperOrders() ([]Models.ScrapperOrder, error) {
	rows, err := r.DB.Query("SELECT id, url FROM scrapperOrder")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()
	var orders []Models.ScrapperOrder
	for rows.Next() {
		var order Models.ScrapperOrder
        if err := rows.Scan(&order.Id, &order.Url); err != nil {
            log.Printf("Error while reading a row : %v", err)
            return nil, err
        }
        orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
        log.Fatalln(err)
        return nil, err
    }

	return orders, nil
}
