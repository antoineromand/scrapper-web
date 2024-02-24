package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"scrapper-web/internal/common"
	"scrapper-web/internal/usecase"
)

func ScrapperController(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
            return
        }
        err := usecase.InsertScrapperOrder(r, db)
        if !err.Success {
            http.Error(w, err.Error.Error(), http.StatusBadRequest)
            return
        }

        response := common.HttpResponse{
			Error:   nil,
			Success: true,
			Code:    http.StatusCreated,
			Message: "Scrapper order created",
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

		defer r.Body.Close()
    }
}