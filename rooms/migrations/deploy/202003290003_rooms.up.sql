CREATE SEQUENCE rooms_id_seq;

CREATE TABLE IF NOT EXISTS rooms(
   id INT UNIQUE DEFAULT NEXTVAL('rooms_id_seq'),
   hotel_id INT NOT NULL REFERENCES rooms(id),
   name VARCHAR (50) UNIQUE NOT NULL,
   quantity INT NOT NULL,
   price DECIMAL(19, 2) NOT NULL
);

ALTER SEQUENCE rooms_id_seq OWNED BY rooms.id;
CREATE INDEX rooms_hotel_id ON rooms(hotel_id);

INSERT INTO rooms (id, hotel_id, name, quantity, price) VALUES (1, 1, 'Melati', 10, 200000);
INSERT INTO rooms (id, hotel_id, name, quantity, price) VALUES (2, 2, 'Melon', 10, 200000);
INSERT INTO rooms (id, hotel_id, name, quantity, price) VALUES (3, 3, 'Jambu', 10, 200000);
INSERT INTO rooms (id, hotel_id, name, quantity, price) VALUES (4, 4, 'Durian', 10, 200000);
INSERT INTO rooms (id, hotel_id, name, quantity, price) VALUES (5, 5, 'Seruti', 10, 200000);

-- Make sure the id increment is continue
BEGIN;
   SELECT setval('rooms_id_seq', COALESCE((SELECT MAX(id)+1 FROM rooms), 1), false);
COMMIT;