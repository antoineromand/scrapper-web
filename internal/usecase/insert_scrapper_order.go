package usecase

import (
	"database/sql"
	"scrapper-web/internal/common/errors"
	Models "scrapper-web/internal/model"
	"scrapper-web/internal/repository"
)



func InsertScrapperOrder(order Models.ScrapperOrder, db *sql.DB) errors.InsertError {
    if order.Url == "" {
        return errors.InsertError{Error: nil, Success: false}
    }
    scrapperOrderRepository := repository.ScrapperOrderRepository{DB: db}    
    err := scrapperOrderRepository.InsertScrapperOrder(order.Url)
    if err != nil {
        return errors.InsertError{Error: err, Success: false}
    }
    return errors.InsertError{Error: nil, Success: true}
}
