CREATE TABLE IF NOT EXISTS order_guests(
   id SERIAL PRIMARY KEY,
   order_id INT NOT NULL REFERENCES orders(id),
   name VARCHAR(50) NOT NULL,
   email VARCHAR(50) NOT NULL,
   phone_number VARCHAR(12) NOT NULL
);

CREATE INDEX order_guests_order_id ON order_guests(order_id);

INSERT INTO order_guests (order_id, name, email, phone_number) VALUES (1, 'madun', 'madun@gmail.com', '08523710');
INSERT INTO order_guests (order_id, name, email, phone_number) VALUES (2, 'miun', 'miun@gmail.com', '08523712');
INSERT INTO order_guests (order_id, name, email, phone_number) VALUES (3, 'mahsun', 'mahsun@gmail.com', '08523713');
INSERT INTO order_guests (order_id, name, email, phone_number) VALUES (4, 'mirun', 'mirun@gmail.com', '08523714');
INSERT INTO order_guests (order_id, name, email, phone_number) VALUES (5, 'masyun', 'masyun@gmail.com', '08523715');