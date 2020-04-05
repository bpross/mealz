package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	"github.com/bpross/mealz/dao/postgres/queries"
	mealzpb "github.com/bpross/mealz/proto"
)

type pgEthnicity string

const (
	american  pgEthnicity = "AMERICAN"
	caribbean pgEthnicity = "CARIBBEAN"
	chinese   pgEthnicity = "CHINESE"
	asian     pgEthnicity = "ASIAN"
	italian   pgEthnicity = "ITALIAN"
	thai      pgEthnicity = "THAI"
	cajun     pgEthnicity = "CAJUN"
	japanese  pgEthnicity = "JAPANESE"
	moroccan  pgEthnicity = "MOROCCAN"
)

// NewRecipes returns Recipes with the supplied configurations
func NewRecipes(db *sql.DB, logger *log.Logger) *Recipes {
	return &Recipes{
		db:     db,
		logger: logger,
	}
}

// Recipes implements the dao.Recipes interface with a postgres backend
type Recipes struct {
	// TODO bpross: make a timing DB implementation
	db     *sql.DB
	logger *log.Logger
}

// Delete deletes the recipe given the recipeObjectID
func (r *Recipes) Delete(ctx context.Context, recipeObjectID bson.ObjectId) (*mealzpb.Recipe, error) {
	logger := r.logger.WithField("recipe_object_id", recipeObjectID.Hex())
	var (
		id        int
		ethnicity pgEthnicity
	)
	recipe := new(mealzpb.Recipe)
	recipe.ObjectId = []byte(recipeObjectID)

	logger.Infoln("deleting")
	err := r.db.QueryRowContext(ctx, queries.RecipeDelete, recipe.ObjectId).Scan(
		&id,
		&recipe.Title,
		&recipe.Vegetarian,
		&ethnicity,
		&recipe.Source,
	)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown recipe: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to delete recipe: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	recipe.Ethnicity = pgEthnicityToProto(ethnicity)

	return recipe, nil
}

// Get retrieves the recipe given the recipeObjectID
func (r *Recipes) Get(ctx context.Context, recipeObjectID bson.ObjectId) (*mealzpb.Recipe, error) {
	logger := r.logger.WithField("recipe_object_id", recipeObjectID.Hex())
	var (
		id        int
		ethnicity pgEthnicity
	)
	recipe := new(mealzpb.Recipe)
	recipe.ObjectId = []byte(recipeObjectID)

	logger.Infoln("retrieving")
	err := r.db.QueryRowContext(ctx, queries.RecipeGet, recipe.ObjectId).Scan(
		&id,
		&recipe.Title,
		&recipe.Vegetarian,
		&ethnicity,
		&recipe.Source,
	)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown recipe: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to retrieve recipe: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	recipe.Ethnicity = pgEthnicityToProto(ethnicity)

	return recipe, nil
}

// Insert inserts the supplied recipe information into the datastore
func (r *Recipes) Insert(ctx context.Context, recipe *mealzpb.Recipe) (*mealzpb.Recipe, error) {
	if recipe.ObjectId != nil {
		return nil, errors.New("cannot specify object_id on insert")
	}

	// Create objectID
	recipeObjectID := bson.NewObjectId()
	recipe.ObjectId = []byte(recipeObjectID)
	logger := r.logger.WithField("recipe_object_id", recipeObjectID.Hex())

	logger.Infoln("inserting")
	var id int
	err := r.db.QueryRowContext(ctx, queries.RecipeInsert, recipe.ObjectId, recipe.Title, recipe.Vegetarian, recipe.Ethnicity.String(), recipe.Source).Scan(&id)
	if err != nil {
		// Write code to wrap errors
		errStr := fmt.Sprintf("failed to insert recipe: %s", err.Error())
		return nil, errors.New(errStr)
	}
	return recipe, nil
}

// Update updates the supplied recipe information in the database
func (r *Recipes) Update(ctx context.Context, recipe *mealzpb.Recipe) (*mealzpb.Recipe, error) {
	logger := r.logger.WithField("recipe_object_id", recipe.ObjectId)
	logger.Infoln("updating")
	var id int

	err := r.db.QueryRowContext(ctx, queries.RecipeUpdate, recipe.ObjectId, recipe.Title, recipe.Vegetarian, recipe.Ethnicity.String(), recipe.Source).Scan(&id)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown recipe: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to update recipe: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	return recipe, nil
}

func pgEthnicityToProto(e pgEthnicity) mealzpb.Ethnicity {
	switch e {
	case american:
		return mealzpb.Ethnicity_AMERICAN
	case caribbean:
		return mealzpb.Ethnicity_CARIBBEAN
	case chinese:
		return mealzpb.Ethnicity_CHINESE
	case asian:
		return mealzpb.Ethnicity_ASIAN
	case italian:
		return mealzpb.Ethnicity_ITALIAN
	case thai:
		return mealzpb.Ethnicity_THAI
	case cajun:
		return mealzpb.Ethnicity_CAJUN
	case japanese:
		return mealzpb.Ethnicity_JAPANESE
	case moroccan:
		return mealzpb.Ethnicity_MOROCCAN
	default:
		return mealzpb.Ethnicity_UNKNOWN_ETHNICITY
	}
}
