CREATE TABLE IF NOT EXISTS room_rates(
   id SERIAL PRIMARY KEY,
   room_id INT NOT NULL REFERENCES rooms(id),
   date DATE NOT NULL,
   price DECIMAL(19, 2) NOT NULL
);

CREATE INDEX room_rates_room_id ON room_rates(room_id);

INSERT INTO room_rates (room_id, date, price) VALUES (1, '2020-03-01', 250000);
INSERT INTO room_rates (room_id, date, price) VALUES (2, '2020-03-01', 250000);
INSERT INTO room_rates (room_id, date, price) VALUES (3, '2020-03-01', 250000);
INSERT INTO room_rates (room_id, date, price) VALUES (4, '2020-03-01', 250000);
INSERT INTO room_rates (room_id, date, price) VALUES (5, '2020-03-01', 250000);