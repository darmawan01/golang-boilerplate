CREATE TABLE IF NOT EXISTS room_availabilities(
   id SERIAL PRIMARY KEY,
   room_id SERIAL NOT NULL REFERENCES rooms(id),
   date DATE NOT NULL,
   quantity INT NOT NULL
);

CREATE INDEX room_availabilities_room_id ON room_availabilities(room_id);