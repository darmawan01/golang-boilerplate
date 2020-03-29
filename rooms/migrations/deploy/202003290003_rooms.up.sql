CREATE TABLE IF NOT EXISTS rooms(
   id serial PRIMARY KEY,
   hotel_id SERIAL NOT NULL REFERENCES hotels(id),
   name VARCHAR (50) UNIQUE NOT NULL,
   quantity INT NOT NULL,
   price DECIMAL(19, 2) NOT NULL
);

CREATE INDEX rooms_hotel_id ON rooms(hotel_id);