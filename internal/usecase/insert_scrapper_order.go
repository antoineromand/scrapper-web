package usecase

import (
	"encoding/json"
	"net/http"
	Models "scrapper-web/internal/model"
)

func ParseScrapperOrder(r *http.Request) (Models.ScrapperOrder, error) {
    var scrapperOrder Models.ScrapperOrder
    err := json.NewDecoder(r.Body).Decode(&scrapperOrder)
    if err != nil {
        return Models.ScrapperOrder{}, err
    }
    return scrapperOrder, nil
}
