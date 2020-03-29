CREATE TABLE IF NOT EXISTS order_items(
   id SERIAL PRIMARY KEY,
   order_id SERIAL NOT NULL REFERENCES orders(id),
   room_id SERIAL NOT NULL REFERENCES orders(id),
   quantity INT NOT NULL,
   price DECIMAL(19,2) NOT NULL
);

CREATE INDEX order_items_order_id ON order_items(order_id);
CREATE INDEX order_items_room_id ON order_items(room_id);