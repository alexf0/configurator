CREATE TABLE IF NOT EXISTS configs (
  id serial PRIMARY KEY,
  type VARCHAR(80) NOT NULL,
  data VARCHAR(80) NOT NULL,
  params json
);