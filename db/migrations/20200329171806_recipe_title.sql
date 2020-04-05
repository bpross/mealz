-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE recipe ADD COLUMN title varchar (1024);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE recipe DROP COLUMN title varchar (1024);
