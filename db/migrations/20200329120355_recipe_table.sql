-- +goose Up
CREATE TYPE ethnicity AS ENUM ('AMERICAN', 'CARIBBEAN', 'CHINESE', 'ASIAN', 'ITALIAN', 'THAI', 'CAJUN', 'JAPANESE', 'MOROCCAN');

CREATE TABLE recipe (
    id SERIAL PRIMARY KEY,
    object_id bytea NOT NULL,
    vegetarian boolean DEFAULT FALSE,
    ethnicity ethnicity,
    source varchar (1024)
);

CREATE UNIQUE INDEX recipe_object_id_idx ON recipe (object_id);

-- +goose Down
DROP TABLE recipe;
DROP TYPE ethnicity;
