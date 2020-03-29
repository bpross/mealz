-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE ingredient (
    id SERIAL PRIMARY KEY,
    object_id bytea NOT NULL,
    vegetarian boolean,
    title varchar (1024)
);

CREATE UNIQUE INDEX ingredient_object_id_idx ON ingredient (object_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE ingredient;
