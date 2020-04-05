-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE season_to_recipe(
    id SERIAL PRIMARY KEY,
    season_id INTEGER NOT NULL,
    recipe_id INTEGER NOT NULL,
    FOREIGN KEY (season_id) REFERENCES season(id) ON DELETE CASCADE,
    FOREIGN KEY (recipe_id) REFERENCES recipe(id) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE season_to_recipe;
