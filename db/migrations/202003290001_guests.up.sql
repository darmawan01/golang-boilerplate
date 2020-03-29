CREATE TABLE IF NOT EXISTS guests(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (50) UNIQUE NOT NULL,
   phone_number VARCHAR (12) UNIQUE NOT NULL,
   identification_id VARCHAR (16) UNIQUE NOT NULL
);

CREATE INDEX guests_id ON guests(id);