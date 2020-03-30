CREATE TABLE IF NOT EXISTS room_availabilities(
   id SERIAL PRIMARY KEY,
   room_id INT NOT NULL REFERENCES rooms(id),
   date DATE NOT NULL,
   quantity INT NOT NULL
);

CREATE INDEX room_availabilities_room_id ON room_availabilities(room_id);

INSERT INTO room_availabilities (room_id, date, quantity) VALUES(1, '2020-03-02', 5);
INSERT INTO room_availabilities (room_id, date, quantity) VALUES(2, '2020-03-02', 9);
INSERT INTO room_availabilities (room_id, date, quantity) VALUES(3, '2020-03-02', 2);
INSERT INTO room_availabilities (room_id, date, quantity) VALUES(4, '2020-03-02', 4);
INSERT INTO room_availabilities (room_id, date, quantity) VALUES(5, '2020-03-02', 7);