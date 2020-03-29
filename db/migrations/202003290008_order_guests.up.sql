CREATE TABLE IF NOT EXISTS order_guests(
   id SERIAL PRIMARY KEY,
   order_id SERIAL NOT NULL REFERENCES orders(id),
   name VARCHAR(50) NOT NULL,
   email VARCHAR(50) NOT NULL,
   phone_number VARCHAR(12) NOT NULL
);

CREATE INDEX order_guests_order_id ON order_guests(order_id);