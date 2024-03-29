package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"scrapper-web/config"
	"scrapper-web/internal/controller"

	"github.com/rs/zerolog/log"
)



func main() {
    log.Info().Msg("Server started on: http://localhost:3333")
    config.InitDB("sqlite3", "../../config/scrapper_db.db")
	log.Info().Msg("Database initialized")
    controller.GetRoutes()
    log.Info().Msg("Controllers initialized")
	err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
            log.Error().Msg("server closed")
        } else if err != nil {
            log.Error().Msg(fmt.Sprintf("server error: %v", err))
            os.Exit(1)
        }
}
