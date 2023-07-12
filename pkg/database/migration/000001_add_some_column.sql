-- +goose Up
CREATE TABLE IF NOT EXISTS publications(
  id BIGSERIAL PRIMARY KEY,
  guid TEXT UNIQUE,
  title TEXT,
  description TEXT,
  pubtime TIMESTAMP WITH TIME ZONE,
  link TEXT
);

-- +goose Down
DROP TABLE publications;
