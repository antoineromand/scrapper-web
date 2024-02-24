package controller

import (
	"net/http"
	"scrapper-web/config"
)

func GetRoutes () {
	db := config.GetDB()
	http.HandleFunc("/scrapper", (ScrapperController(db)))
}