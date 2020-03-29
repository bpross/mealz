-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE season (
    id SERIAL PRIMARY KEY,
    title varchar (10)
);

CREATE UNIQUE INDEX season_title_idx ON season (title);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE season;
