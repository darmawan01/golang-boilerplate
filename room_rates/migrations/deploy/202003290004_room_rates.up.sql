CREATE TABLE IF NOT EXISTS room_rates(
   id SERIAL PRIMARY KEY,
   room_id SERIAL NOT NULL REFERENCES rooms(id),
   date DATE NOT NULL,
   price DECIMAL(19, 2) NOT NULL
);

CREATE INDEX room_rates_room_id ON room_rates(room_id);