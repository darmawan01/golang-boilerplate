DO $$ BEGIN
  CREATE TYPE status_type AS ENUM ('PENDING','PAID', 'READY', 'CECKIN', 'CHECKOUT', 'EXPIRED', 'CANCELLED');
EXCEPTION
  WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS orders(
   id serial PRIMARY KEY,
   hotel_id SERIAL NOT NULL REFERENCES orders(id),
   guest_id SERIAL NOT NULL REFERENCES orders(id),
   status status_type NOT NULL,
   check_in timestamp NOT NULL,
   check_out timestamp NOT NULL,
   created_at timestamp NOT NULL
);

CREATE INDEX orders_hotels_id ON orders(hotel_id);
CREATE INDEX orders_guest_id ON orders(guest_id);