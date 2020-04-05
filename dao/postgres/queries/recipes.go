package queries

const RecipeInsert = `
INSERT INTO recipe (object_id, title, vegetarian, ethnicity, source)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
`

const RecipeGet = `
SELECT id, title, vegetarian, ethnicity, source
FROM recipe
WHERE object_id = $1;
`

const RecipeUpdate = `
UPDATE recipe
SET title = $2, vegetarian = $3, ethnicity = $4, source = $5
WHERE object_id = $1
RETURNING id;
`

const RecipeDelete = `
DELETE FROM recipe
WHERE object_id = $1
RETURNING id, title, vegetarian, ethnicity, source;
`
