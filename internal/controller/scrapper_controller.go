package controller

import (
	"database/sql"
	"net/http"
	"scrapper-web/config"
	"scrapper-web/internal/usecase"
)


func ScrapperController(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
            return
        }
        scrapperOrder, err := usecase.ParseScrapperOrder(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        // TO DO: MOVE THIS TO THE USECASE METHOD ParseScrapperOrder and rename it to SaveScrapperOrder
        config.InsertData(db, scrapperOrder.Url)
        defer r.Body.Close()
    }
}