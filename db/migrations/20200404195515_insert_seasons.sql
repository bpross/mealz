-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO season (title) VALUES ('winter');
INSERT INTO season (title) VALUES ('spring');
INSERT INTO season (title) VALUES ('summer');
INSERT INTO season (title) VALUES ('fall');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM season WHERE title = 'winter';
DELETE FROM season WHERE title = 'spring';
DELETE FROM season WHERE title = 'summer';
DELETE FROM season WHERE title = 'fall';
