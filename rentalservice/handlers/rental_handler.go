package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"rental_service/db"
	// "rental_service/models"
)

// GetAllTrailers godoc
// @Summary Get all trailers
// @Description Fetches a list of all trailers from the database
// @Tags trailers
// @Produce application/json
// @Success 200 {array} models.Trailer
// @Router /api/trailer [get]
func GetAllTrailers(w http.ResponseWriter, r *http.Request) {
	trailers, err := db.FetchAllTrailers()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(trailers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetTrailerByZip godoc
// @Summary Get a trailer
// @Description Fetches a trailer based on the zipcode from the database
// @Tags trailer
// @Produce application/json
// @Param zip path string true "Zip"
// @Success 200 {object} models.Trailer
// @Router /api/trailer/{zip} [get]
func GetTrailerByZip(w http.ResponseWriter, r *http.Request) {
	zip := r.PathValue("zip")
	trailer, err := db.GetTrailerByZip(zip)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := json.Marshal(trailer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type CreateRentalRequest struct {
	TrailerID int `json:"trailer_id"`
	UserID    int `json:"user_id"`
}

// CreateRental godoc
// @Summary Create a new rental
// @Description Create a new rental entry in the database
// @Tags rental
// @Accept  application/json
// @Produce application/json
// @Param request body CreateRentalRequest true "CreateRentalRequest"
// @Success 201 {object} models.Rental
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/rental [post]
func CreateRental(w http.ResponseWriter, r *http.Request) {
	var req CreateRentalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = db.CreateRental(req.TrailerID, req.UserID)
	if err != nil {
		http.Error(w, "Failed to create rental", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Rental created successfully"})
}
