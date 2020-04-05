package queries

// RecipeInsert inserts a recipe
const RecipeInsert = `
INSERT INTO recipe (object_id, title, vegetarian, ethnicity, source)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
`

// RecipeGet retrieves a specific recipe
const RecipeGet = `
SELECT id, title, vegetarian, ethnicity, source
FROM recipe
WHERE object_id = $1;
`

// RecipeUpdate updates all of the values for the given recipe
const RecipeUpdate = `
UPDATE recipe
SET title = $2, vegetarian = $3, ethnicity = $4, source = $5
WHERE object_id = $1
RETURNING id;
`

// RecipeDelete removes the specific recipe
const RecipeDelete = `
DELETE FROM recipe
WHERE object_id = $1
RETURNING id, title, vegetarian, ethnicity, source;
`
