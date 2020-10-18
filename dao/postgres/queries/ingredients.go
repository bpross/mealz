package queries

// IngredientGet retrieves a specific ingredient
const IngredientGet = `
SELECT id, vegetarian, title
FROM ingredient
WHERE object_id = $1;
`

// IngredientInsert inserts an ingredient
const IngredientInsert = `
INSERT INTO ingredient (object_id, vegetarian, title)
VALUES ($1, $2, $3)
RETURNING id;
`

// IngredientAssociate associates an ingredient to a recipe and sets amount information
const IngredientAssociate = `
INSERT INTO ingredient_to_recipe (ingredient_id, recipe_id, unit_of_measure, amount)
VALUES ((SELECT id FROM ingredient WHERE object_id = $1), (SELECT id FROM recipe WHERE object_id = $2), $3, $4)
RETURNING id;
`

// IngredientDelete removes the specific ingredient
const IngredientDelete = `
DELETE FROM ingredient
WHERE object_id = $1
RETURNING id, vegetarian, title;
`
