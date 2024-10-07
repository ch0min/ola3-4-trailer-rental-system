package db

import (
	"rental_service/models"
	"time"
)

func FetchAllTrailers() ([]models.Trailer, error) {
	trailers := []models.Trailer{}

	rows, err := DB.Query(`
		SELECT
			t.trailer_id,
			t.number,
			t.availability_status,
			l.location_id,
			l.location_name,
			a.address_id,
			a.street,
			a.city,
			a.state,
			a.zipcode
		FROM
			trailers t
		JOIN
			locations l ON t.location_id = l.location_id
		JOIN
			addresses a ON l.address_id = a.address_id;
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trailer models.Trailer
		var location models.Location
		var address models.Address

		err := rows.Scan(
			&trailer.ID,
			&trailer.Number,
			&trailer.AvailabilityStatus,
			&location.ID,
			&location.Name,
			&address.ID,
			&address.Street,
			&address.City,
			&address.State,
			&address.Zipcode,
		)
		if err != nil {
			return nil, err
		}

		location.Address = address
		trailer.Location = location

		trailers = append(trailers, trailer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return trailers, nil

}

func GetTrailerByZip(zipCode string) ([]models.Trailer, error) {
	trailers := []models.Trailer{}

	query := `
		SELECT
			t.trailer_id,
			t.number,
			t.availability_status,
			l.location_id,
			l.location_name,
			a.address_id,
			a.street,
			a.city,
			a.state,
			a.zipcode
		FROM
			trailers t
		JOIN
			locations l ON t.location_id = l.location_id
		JOIN
			addresses a ON l.address_id = a.address_id
		WHERE
			a.zipcode = $1;
	`

	rows, err := DB.Query(query, zipCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trailer models.Trailer
		var location models.Location
		var address models.Address

		err := rows.Scan(
			&trailer.ID,
			&trailer.Number,
			&trailer.AvailabilityStatus,
			&location.ID,
			&location.Name,
			&address.ID,
			&address.Street,
			&address.City,
			&address.State,
			&address.Zipcode,
		)
		if err != nil {
			return nil, err
		}

		location.Address = address
		trailer.Location = location

		trailers = append(trailers, trailer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trailers, nil
}

func CreateRental(trailerId int, userId int) error {
	query := `
		INSERT INTO rentals (customer_id, trailer_id, start_time, rental_fee)
		VALUES ($1, $2, $3, $4)
		RETURNING rental_id;
	`

	startTime := time.Now()

	rentalFee := 100.0 // Should there be a rental fee?
	//maybe also add end_time and excess_fee

	var rentalID int
	err := DB.QueryRow(query, userId, trailerId, startTime, rentalFee).Scan(&rentalID)
	if err != nil {
		return err
	}

	return nil
}
