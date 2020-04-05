-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE season_to_recipe(
    id SERIAL PRIMARY KEY,
    season_id INTEGER REFERENCES season(id) NOT NULL,
    recipe_id INTEGER REFERENCES recipe(id) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE season_to_recipe;
