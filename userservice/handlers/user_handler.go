package handlers

import (
	"encoding/json"
	"net/http"
	"user_service/db"
	"user_service/models"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieves all users from the database
// @Tags users
// @Produce application/json
// @Success 200 {array} models.Customer
// @Failure 500 {object} map[string]string "error"
// @Router /api/user [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Call the GetAllUsers function from the db package to get all users
	users, err := db.GetAllUsers()
	if err != nil {
		// Return a 500 Internal Server Error if retrieving users fails
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Return the list of users in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}


// CreateUser godoc
// @Summary Creates a new user
// @Description Creates a new user entry in the database
// @Tags users
// @Accept json
// @Produce application/json
// @Param user body models.Customer true "User Data"
// @Success 201 {object} map[string]int "user_id"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /api/user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the user details
	var user models.Customer
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Return a 400 Bad Request if JSON parsing fails
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the CreateUser function from the db package to create a new user
	userID, err := db.CreateUser(user)
	if err != nil {
		// Return a 500 Internal Server Error if database insertion fails
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return the newly created user ID in the response
	response := map[string]int{"user_id": userID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



// GetUserById godoc
// @Summary Get a user by ID
// @Description Retrieves a user from the database based on the user ID
// @Tags users
// @Produce application/json
// @Param id path string true "User ID"
// @Success 200 {object} models.Customer
// @Failure 500 {object} map[string]string "error"
// @Router /api/user/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	// Call the GetAllUsers function from the db package to get all users
	id := r.PathValue("id")
	users, err := db.GetUserById(id)
	if err != nil {
		// Return a 500 Internal Server Error if retrieving users fails
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Return the list of users in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}


