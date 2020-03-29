-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TYPE unitOfMeasure AS ENUM ('TEASPOON', 'TABLESPOON', 'POUND', 'OUNCE', 'CUP', 'PIECE');

CREATE TABLE ingredient_to_recipe(
    id SERIAL PRIMARY KEY,
    ingredient_id INTEGER REFERENCES ingredient(id),
    recipe_id INTEGER REFERENCES recipe(id),
    unit_of_measure unitOfMeasure,
    amount float
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE ingredient_to_recipe;
DROP TYPE unitOfMeasure;
