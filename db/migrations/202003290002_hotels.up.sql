CREATE TABLE IF NOT EXISTS hotels(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   address TEXT NOT NULL,
   latitude TEXT NOT NULL,
   longitude TEXT NOT NULL
);

CREATE INDEX hotels_id ON hotels(id);