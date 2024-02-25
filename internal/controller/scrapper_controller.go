package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"scrapper-web/internal/common"
	Models "scrapper-web/internal/model"
	"scrapper-web/internal/usecase"
)

func ScrapperController(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        switch r.Method {
		case "GET":
			GetAllScrapperOrder(db, r, w)
		case "POST":
			AddScrapperOrder(db, r, w)
		}
    }
}

func AddScrapperOrder(db *sql.DB, r *http.Request, w http.ResponseWriter) {
		var scrapperOrder Models.ScrapperOrder
		err := json.NewDecoder(r.Body).Decode(&scrapperOrder)
		if err != nil {
			return
		}
		res := usecase.InsertScrapperOrder(scrapperOrder, db)
		if !res.Success {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
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

func GetAllScrapperOrder(db *sql.DB, r *http.Request, w http.ResponseWriter) {
    orders, err := usecase.GetAllScrapperOrder(r, db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := common.HttpResponse{
        Error:   nil,
        Success: true,
        Code:    http.StatusOK,
        Message: "Scrapper orders retrieved",
        Data:    orders,
    }

    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Println("Failed to encode response:", err)
    }
}