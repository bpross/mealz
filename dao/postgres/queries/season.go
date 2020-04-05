package queries

// SeasonGet queries for a specific season
const SeasonGet = `
SELECT id
FROM season
WHERE title = $1;
`

// SeasonAssociate associates a recipe to a season
const SeasonAssociate = `
INSERT INTO season_to_recipe (season_id, recipe_id)
VALUES ((SELECT id FROM season WHERE title = $2), (SELECT id FROM recipe WHERE object_id = $1))
RETURNING id;
`

// GetRecipeSeasons retrieves all of the seasons for a given recipe
const GetRecipeSeasons = `
SELECT s.id, s.title FROM season s
JOIN season_to_recipe sr ON sr.season_id = s.id
JOIN recipe r ON r.id = sr.recipe_id WHERE r.object_id = $1;
`
