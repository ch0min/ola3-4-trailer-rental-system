-- +goose Up
-- +goose StatementBegin
INSERT INTO addresses (street, city, state, zipcode)
    VALUES ('Hovedbanegården 1', 'København', 'Hovedstaden', '1570'),
    ('Dokk1, Hack Kampmanns Plads 2', 'Aarhus', 'Midtjylland', '8000'),
    ('Jernbanegade 4', 'Odense', 'Syddanmark', '5000');

INSERT INTO customers (name, email, phone_number, address_id)
    VALUES ('Lars Hansen', 'lars.hansen@example.com', '12345678', 1),
    ('Mette Sørensen', 'mette.sorensen@example.com', '23456789', 2),
    ('Pia Jensen', 'pia.jensen@example.com', '34567890', 3);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
