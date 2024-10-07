package db

import (
	"rental_service/models"
)

// Create a new customer in the database
func CreateUser(user models.Customer) (int, error) {
	// Prepare the SQL statement
	stmt, err := DB.Prepare(`
		INSERT INTO customers (name, email, phone_number)
		VALUES ($1, $2, $3, $4) RETURNING customer_id`)
	if err != nil {
		return 0, err // Return the error instead of logging and terminating
	}
	defer stmt.Close()

	// Variable to store the returned customer ID
	var customerID int

	// Execute the query and scan the result into customerID
	err = stmt.QueryRow(user.Name, user.Email, user.PhoneNumber).Scan(&customerID)
	if err != nil {
		return 0, err
	}

	// Return the newly created customer ID
	return customerID, nil
}

// Retrieve all users from the database
func GetAllUsers() ([]models.Customer, error) {
	users := []models.Customer{}

	// Execute the query
	row, err := DB.Query("SELECT customer_id, name, email, phone_number FROM customers")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	// Loop through the rows and populate the users slice
	for row.Next() {
		var user models.Customer

		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Check if there were any errors while scanning
	if err = row.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Retrieve a user by ID from the database
func GetUserById(id string) (models.Customer, error) {
	var user models.Customer

	// Query for a specific user by ID
	err := DB.QueryRow("SELECT customer_id, name, email, phone_number FROM customers WHERE customer_id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber)
	if err != nil {
		return user, err
	}

	return user, nil
}
