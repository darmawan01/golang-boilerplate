CREATE SEQUENCE hotels_id_seq;

CREATE TABLE IF NOT EXISTS hotels(
   id INT UNIQUE DEFAULT NEXTVAL('hotels_id_seq'),
   name VARCHAR (50) UNIQUE NOT NULL,
   address TEXT NOT NULL,
   latitude TEXT NOT NULL,
   longitude TEXT NOT NULL
);

ALTER SEQUENCE hotels_id_seq OWNED BY hotels.id;
CREATE INDEX hotels_id ON hotels(id);

INSERT INTO hotels (id, name, address, latitude, longitude) VALUES (1, 'hotel madun', 'Jl. kembang belok kiri kanan', '085.23712', '0101.02');
INSERT INTO hotels (id, name, address, latitude, longitude) VALUES (2, 'hotel miun', 'Jl. kembang belok kiri kanan', '085.23712', '0101.02');
INSERT INTO hotels (id, name, address, latitude, longitude) VALUES (3, 'hotel mahsun', 'Jl. kembang belok kiri kanan', '085.23713', '0101.03');
INSERT INTO hotels (id, name, address, latitude, longitude) VALUES (4, 'hotel mirun', 'Jl. kembang belok kiri kanan', '085.23714', '0101.04');
INSERT INTO hotels (id, name, address, latitude, longitude) VALUES (5, 'hotel masyun', 'Jl. kembang belok kiri kanan', '085.23715', '0101.05');

-- Make sure the id increment is continue
BEGIN;
   SELECT setval('hotels_id_seq', COALESCE((SELECT MAX(id)+1 FROM hotels), 1), false);
COMMIT;