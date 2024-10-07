-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    address_id serial PRIMARY KEY,
    street varchar(255) NOT NULL,
    city varchar(100) NOT NULL,
    state varchar(100),
    zipcode varchar(20) NOT NULL
);


CREATE TABLE customers (
    customer_id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    phone_number varchar(20) NOT NULL,
    address_id int REFERENCES addresses (address_id) ON DELETE SET NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS customers;

DROP TABLE IF EXISTS addresses;

-- +goose StatementEnd
