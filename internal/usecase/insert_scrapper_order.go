package usecase

import (
	"database/sql"
	"encoding/json"
	"net/http"
	Models "scrapper-web/internal/model"
	"scrapper-web/internal/repository"
)

type InsertScrapperError struct {
    Error error
    Success bool
}

func InsertScrapperOrder(r *http.Request, db *sql.DB) InsertScrapperError {
    var scrapperOrder Models.ScrapperOrder
    err := json.NewDecoder(r.Body).Decode(&scrapperOrder)
    if err != nil {
        return InsertScrapperError{Error: err, Success: false}
    }
    if scrapperOrder.Url == "" {
        return InsertScrapperError{Error: nil, Success: false}
    }
    scrapperOrderRepository := repository.ScrapperOrderRepository{
            BaseRepository: repository.BaseRepository{
                DB: db,
            },
        }    
    err = scrapperOrderRepository.InsertScrapperOrder(scrapperOrder.Url)
    if err != nil {
        return InsertScrapperError{Error: err, Success: false}
    }
    return InsertScrapperError{Error: nil, Success: true}
}
