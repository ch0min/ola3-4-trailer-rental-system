package models


type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state,omitempty"`
	Zipcode string `json:"zipcode"`
}


type Customer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Address     Address `json:"address"`
}

