package handlers

import (
	"net/http"

	"rental_service/db"
	"rental_service/models"
)

// CreateUser godoc
// @Summary Get all trailers
// @Description Fetches a list of all trailers from the database
// @Tags trailers
// @Produce application/json
// @Success 200 {array} models.Customer
// @Router /api/user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := db.CreateUser()
	// TODO: maybe error handling

	w.WriteHeader(http.StatusCreated)
}
