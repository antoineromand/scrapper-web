package controller

import (
	"net/http"
	"scrapper-web/config"

	"github.com/rs/zerolog/log"
)

func GetRoutes () {
	var db = config.InitDB("sqlite3", "scrapper.db")
	log.Info().Msg("Database initialized")
	http.HandleFunc("/scrapper", (ScrapperController(db)))
}