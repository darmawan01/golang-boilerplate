CREATE SEQUENCE guests_id_seq;

CREATE TABLE IF NOT EXISTS guests(
   id INT UNIQUE DEFAULT NEXTVAL('guests_id_seq'),
   name VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (50) UNIQUE NOT NULL,
   phone_number VARCHAR (12) UNIQUE NOT NULL,
   identification_id VARCHAR (16) UNIQUE NOT NULL
);

ALTER SEQUENCE guests_id_seq OWNED BY guests.id;
CREATE INDEX guests_id ON guests(id);

INSERT INTO guests (id, name, email, phone_number, identification_id) VALUES (1, 'madun', 'madun@gmail.com', '08523710', '00110010101');
INSERT INTO guests (id, name, email, phone_number, identification_id) VALUES (2, 'miun', 'miun@gmail.com', '08523712', '00110010102');
INSERT INTO guests (id, name, email, phone_number, identification_id) VALUES (3, 'mahsun', 'mahsun@gmail.com', '08523713', '00110010103');
INSERT INTO guests (id, name, email, phone_number, identification_id) VALUES (4, 'mirun', 'mirun@gmail.com', '08523714', '00110010104');
INSERT INTO guests (id, name, email, phone_number, identification_id) VALUES (5, 'masyun', 'masyun@gmail.com', '08523715', '00110010105');

-- Make sure the id increment is continue
BEGIN;
   SELECT setval('guests_id_seq', COALESCE((SELECT MAX(id)+1 FROM guests), 1), false);
COMMIT;