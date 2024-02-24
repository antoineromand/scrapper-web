package usecase

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"scrapper-web/internal/common/errors"
	Models "scrapper-web/internal/model"
	"scrapper-web/internal/repository"
)



func InsertScrapperOrder(r *http.Request, db *sql.DB) errors.InsertError {
    var scrapperOrder Models.ScrapperOrder
    err := json.NewDecoder(r.Body).Decode(&scrapperOrder)
    if err != nil {
        return errors.InsertError{Error: err, Success: false}
    }
    if scrapperOrder.Url == "" {
        return errors.InsertError{Error: nil, Success: false}
    }
    scrapperOrderRepository := repository.ScrapperOrderRepository{
            BaseRepository: repository.BaseRepository{
                DB: db,
            },
        }    
    err = scrapperOrderRepository.InsertScrapperOrder(scrapperOrder.Url)
    if err != nil {
        return errors.InsertError{Error: err, Success: false}
    }
    return errors.InsertError{Error: nil, Success: true}
}
