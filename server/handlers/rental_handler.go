package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"rental_service/db"
	// "rental_service/models"
)

// TODO: getalltrailers with location - addresses (search by zipcode)

// TODO: create new rental - need user login before. ("guest"-user service, / maybe also payment service)

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

// CreateRental godoc
// @Summary Create a new rental
// @Description Create a new rental entry in the database
// @Tags rental
// @Accept  application/json
// @Produce application/json
// @Param trailer_id query string true "trailerId"
// @Param user_id query string true "userId"
// @Success 201 {object} models.Rental
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/rental [post]
func CreateRental(w http.ResponseWriter, r *http.Request) {
	// TODO: make function read trailer_id and user_id from body and use in the db function
	// var trailerId int
	// var userId int
	//
	// err := json.NewDecoder(r.Body).Decode()
	//
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = db.CreateRental(trailerId, userId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//
	w.WriteHeader(http.StatusCreated)
}
