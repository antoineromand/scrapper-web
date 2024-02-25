package usecase

import (
	"database/sql"
	"net/http"
	Models "scrapper-web/internal/model"
	"scrapper-web/internal/repository"
)

func GetAllScrapperOrder(r *http.Request, db *sql.DB) ([]Models.ScrapperOrder, error) {
	scrapperOrderRepository := repository.ScrapperOrderRepository{DB: db}
	orders, err := scrapperOrderRepository.GetScrapperOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}