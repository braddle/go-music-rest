CREATE TABLE IF NOT EXISTS artist (
    id serial PRIMARY KEY,
    name VARCHAR ( 255 ) NOT NULL,
    image VARCHAR ( 255 ) NOT NULL,
    genres VARCHAR ( 255 ) NOT NULL,
    year_started int NOT NULL,
    year_ened int
);