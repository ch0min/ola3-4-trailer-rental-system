package db

import (
	"database/sql"
	"log"

	"rental_service/models"
)

func FetchAllTrailers() ([]models.Trailer, error) {
	trailers := []models.Trailer{}
	//TODO: change sql query
	rows, err := DB.Query("SELECT id, title, text, isCompleted, category, deadline FROM todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var trailer models.Trailer

		//TODO: change fields to match trailer
		var category sql.NullString
		var deadline sql.NullTime

		err := rows.Scan(&trailer.ID, &trailer.Title, &trailer.Body, &trailer.Done, &category, &deadline)
		if err != nil {
			log.Fatal(err)
		}

		trailers = append(trailers, trailer)
	}
	if err != nil {
		log.Fatal(err)
	}
	return trailers, err
}

func GetTrailerByZip(zip string) (models.Trailer, error) {
	trailer := models.Trailer{}

	// TODO: change sql
	row := DB.QueryRow("SELECT id, title, text, isCompleted, category, deadline FROM todo WHERE id = $1", zip)

	// TODO: update fields to match trailer
	var category sql.NullString
	var deadline sql.NullTime

	err := row.Scan(&trailer.ID, &trailer.Title, &trailer.Body, &trailer.Done, &category, &deadline)
	if err != nil {
		log.Fatal(err)
	}

	return trailer, err
}

func CreateRental(trailerId int, userId int) error {
	// TODO change implementation
	var lastInsertId int
	query := `INSERT INTO todo (title, text, iscompleted, category, deadline)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := DB.QueryRow(query, todo.Title, todo.Body, todo.Done, todo.Category, todo.Deadline).Scan(&lastInsertId)
	return err
}
