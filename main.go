package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"scrapper-web/config"
	"scrapper-web/controller"
)



func main() {
    fmt.Println("Server started on: http://localhost:3333")
    db := config.InitDB("scrapper.db")
    fmt.Println("Database initialized")
    http.HandleFunc("/scrapper", (controller.ScrapperController(db)))
    fmt.Println("Scrapper controller initialized")
	err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
            fmt.Printf("server closed\n")
        } else if err != nil {
            fmt.Printf("error starting server: %s\n", err)
            os.Exit(1)
        }
}
