-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    address_id serial PRIMARY KEY,
    street varchar(255) NOT NULL,
    city varchar(100) NOT NULL,
    state varchar(100),
    zipcode varchar(20) NOT NULL
);

CREATE TABLE locations (
    location_id serial PRIMARY KEY,
    location_name varchar(255) NOT NULL,
    address_id int REFERENCES addresses (address_id) ON DELETE SET NULL
);

CREATE TABLE customers (
    customer_id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    phone_number varchar(20) NOT NULL,
    address_id int REFERENCES addresses (address_id) ON DELETE SET NULL
);

CREATE TABLE trailers (
    trailer_id serial PRIMARY KEY,
    number varchar(50) NOT NULL UNIQUE,
    availability_status varchar(50) NOT NULL,
    location_id int REFERENCES locations (location_id) ON DELETE SET NULL
);

CREATE TABLE rentals (
    rental_id serial PRIMARY KEY,
    customer_id int REFERENCES customers (customer_id) ON DELETE CASCADE,
    trailer_id int REFERENCES trailers (trailer_id) ON DELETE CASCADE,
    start_time timestamp NOT NULL,
    end_time timestamp,
    rental_fee DECIMAL(10, 2) NOT NULL,
    excess_fee DECIMAL(10, 2)
);

CREATE TABLE partnerships (
    partnership_id serial PRIMARY KEY,
    company_name varchar(255) NOT NULL,
    partnership_start_date date NOT NULL,
    partnership_end_date date
);

CREATE TABLE payments (
    payment_id serial PRIMARY KEY,
    rental_id int REFERENCES rentals (rental_id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date timestamp DEFAULT CURRENT_TIMESTAMP,
    payment_type varchar(50) NOT NULL,
    status varchar(50) NOT NULL
);

CREATE TABLE insurances (
    insurance_id serial PRIMARY KEY,
    rental_id int REFERENCES rentals (rental_id) ON DELETE CASCADE,
    insurance_fee DECIMAL(10, 2) NOT NULL,
    provider varchar(255) NOT NULL,
    insurance_details text NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS insurances;

DROP TABLE IF EXISTS payments;

DROP TABLE IF EXISTS partnerships;

DROP TABLE IF EXISTS rentals;

DROP TABLE IF EXISTS trailers;

DROP TABLE IF EXISTS customers;

DROP TABLE IF EXISTS locations;

DROP TABLE IF EXISTS addresses;

-- +goose StatementEnd
