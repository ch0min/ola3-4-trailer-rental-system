-- +goose Up
-- +goose StatementBegin
INSERT INTO addresses (street, city, state, zipcode)
    VALUES ('Hovedbanegården 1', 'København', 'Hovedstaden', '1570'),
    ('Dokk1, Hack Kampmanns Plads 2', 'Aarhus', 'Midtjylland', '8000'),
    ('Jernbanegade 4', 'Odense', 'Syddanmark', '5000');

INSERT INTO locations (location_name, address_id)
    VALUES ('København H', 1),
    ('Aarhus Ø', 2),
    ('Odense C', 3);

INSERT INTO customers (name, email, phone_number, address_id)
    VALUES ('Lars Hansen', 'lars.hansen@example.com', '12345678', 1),
    ('Mette Sørensen', 'mette.sorensen@example.com', '23456789', 2),
    ('Pia Jensen', 'pia.jensen@example.com', '34567890', 3);

INSERT INTO trailers (number, availability_status, location_id)
    VALUES ('T-001', 'ledig', 1),
    ('T-002', 'udlejet', 2),
    ('T-003', 'ledig', 3);

INSERT INTO rentals (customer_id, trailer_id, start_time, end_time, rental_fee, excess_fee)
    VALUES (1, 1, '2024-10-01 10:00:00', '2024-10-01 14:00:00', 350.00, 0),
    (2, 2, '2024-10-03 09:00:00', '2024-10-03 12:00:00', 400.00, 0),
    (3, 3, '2024-10-05 11:00:00', NULL, 500.00, NULL);

INSERT INTO partnerships (company_name, partnership_start_date, partnership_end_date)
    VALUES ('Jem og Fix', '2023-01-01', NULL),
    ('Fog Trailere', '2022-05-15', '2024-12-31');

INSERT INTO payments (rental_id, amount, payment_type, status)
    VALUES (1, 350.00, 'Kreditkort', 'Behandlet'),
    (2, 400.00, 'MobilePay', 'Behandlet');

INSERT INTO insurances (rental_id, insurance_fee, provider, insurance_details)
    VALUES (1, 100.00, 'TrailerForsikring', 'Basis dækning af trailer'),
    (2, 150.00, 'TrailerSikring', 'Fuld dækning af trailer');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
