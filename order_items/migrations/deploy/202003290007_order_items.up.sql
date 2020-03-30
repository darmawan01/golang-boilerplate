CREATE TABLE IF NOT EXISTS order_items(
   id SERIAL PRIMARY KEY,
   order_id INT NOT NULL REFERENCES orders(id),
   room_id INT NOT NULL REFERENCES orders(id),
   quantity INT NOT NULL,
   price DECIMAL(19,2) NOT NULL
);

CREATE INDEX order_items_order_id ON order_items(order_id);
CREATE INDEX order_items_room_id ON order_items(room_id);

INSERT INTO order_items (order_id, room_id, quantity, price) VALUES (1, 1, 1, 250000);
INSERT INTO order_items (order_id, room_id, quantity, price) VALUES (2, 2, 1, 250000);
INSERT INTO order_items (order_id, room_id, quantity, price) VALUES (3, 3, 1, 250000);
INSERT INTO order_items (order_id, room_id, quantity, price) VALUES (4, 4, 1, 250000);