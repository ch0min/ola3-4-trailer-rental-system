package main

import (
	"fmt"
	"log"
	"net/http"

	"rental_service/db"
	_ "rental_service/docs"
	"rental_service/handlers"

	"github.com/rs/cors"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func run() (*http.ServeMux, error) {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /api/docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /api/trailer", handlers.GetAllTrailers)
	mux.HandleFunc("GET /api/trailer/{zip}", handlers.GetTrailerByZip)
	mux.HandleFunc("POST /api/rental", handlers.CreateRental)

	return mux, nil
}

func main() {
	mux, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close() // Ensure the database connection is closed at the end

	handler := cors.Default().Handler(mux)

	fmt.Println("Running server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", handler))
}
