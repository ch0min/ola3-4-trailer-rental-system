package models

import "time"

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state,omitempty"`
	Zipcode string `json:"zipcode"`
}

type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// AddressID int     `json:"address_id"`
	Address Address `json:"address"`
}

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	// AddressID   int     `json:"address_id"`
	Address Address `json:"address"`
}

type Trailer struct {
	ID                 int    `json:"id"`
	Number             string `json:"number"`
	AvailabilityStatus string `json:"availability_status"`
	// LocationID         int      `json:"location_id"`
	Location Location `json:"location"`
}

type Rental struct {
	ID         int        `json:"id"`
	CustomerID int        `json:"customer_id"`
	TrailerID  int        `json:"trailer_id"`
	StartTime  time.Time  `json:"start_time"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	RentalFee  float64    `json:"rental_fee"`
	ExcessFee  *float64   `json:"excess_fee,omitempty"`
	Customer   Customer   `json:"customer"`
	Trailer    Trailer    `json:"trailer"`
}

type Partnership struct {
	ID                   int        `json:"id"`
	CompanyName          string     `json:"company_name"`
	PartnershipStartDate time.Time  `json:"partnership_start_date"`
	PartnershipEndDate   *time.Time `json:"partnership_end_date,omitempty"`
}

type Payment struct {
	ID          int       `json:"id"`
	RentalID    int       `json:"rental_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	PaymentType string    `json:"payment_type"`
	Status      string    `json:"status"`
	Rental      Rental    `json:"rental"`
}

type Insurance struct {
	ID               int     `json:"id"`
	RentalID         int     `json:"rental_id"`
	InsuranceFee     float64 `json:"insurance_fee"`
	Provider         string  `json:"provider"`
	InsuranceDetails string  `json:"insurance_details"`
	Rental           Rental  `json:"rental"`
}
