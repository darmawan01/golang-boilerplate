DO $$ BEGIN
  CREATE TYPE status_type AS ENUM ('PENDING','PAID', 'READY', 'CECKIN', 'CHECKOUT', 'EXPIRED', 'CANCELLED');
EXCEPTION
  WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE orders_id_seq;

CREATE TABLE IF NOT EXISTS orders(
   id INT UNIQUE DEFAULT NEXTVAL('orders_id_seq'),
   hotel_id INT NOT NULL REFERENCES hotels(id),
   guest_id INT NOT NULL REFERENCES guests(id),
   status status_type NOT NULL,
   check_in timestamp NOT NULL,
   check_out timestamp NOT NULL,
   created_at timestamp DEFAULT NOW()
);

ALTER SEQUENCE orders_id_seq OWNED BY orders.id;
CREATE INDEX orders_hotel_id ON orders(hotel_id);
CREATE INDEX orders_guest_id ON orders(guest_id);

INSERT INTO orders (id, hotel_id, guest_id, status, check_in, check_out) VALUES (1, 1, 1, 'PENDING', '2020-04-01 21:03:07.691858', '2020-04-03 21:03:07.691858');
INSERT INTO orders (id, hotel_id, guest_id, status, check_in, check_out) VALUES (2, 2, 2, 'READY', '2020-04-01 21:03:07.691858', '2020-04-03 21:03:07.691858');
INSERT INTO orders (id, hotel_id, guest_id, status, check_in, check_out) VALUES (3, 3, 3, 'PENDING', '2020-04-01 21:03:07.691858', '2020-04-03 21:03:07.691858');
INSERT INTO orders (id, hotel_id, guest_id, status, check_in, check_out) VALUES (4, 4, 4, 'PAID', '2020-04-01 21:03:07.691858', '2020-04-03 21:03:07.691858');
INSERT INTO orders (id, hotel_id, guest_id, status, check_in, check_out) VALUES (5, 5, 5, 'PENDING', '2020-04-01 21:03:07.691858', '2020-04-03 21:03:07.691858');
-- Make sure the id increment is continue
BEGIN;
   SELECT setval('orders_id_seq', COALESCE((SELECT MAX(id)+1 FROM orders), 1), false);
COMMIT;