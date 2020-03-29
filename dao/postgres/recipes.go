package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	"github.com/bpross/mealz/dao"
)

// NewRecipes returns Recipes with the supplied configurations
func NewRecipes(db *sql.DB, logger log.Logger) {
	return &Recipes{
		db:     db,
		logger: logger,
	}
}

// Recipes implements the dao.Recipes interface with a postgres backend
type Recipes struct {
	db     *sql.DB
	logger *log.Logger
}

// Delete deletes the recipe given the recipeObjectID
func (r *Recipes) Delete(ctx context.Context, recipeObjectID bson.ObjectId) (mealzpb.Recipe, error) {
	return nil, errors.New("unimplemented")
}

// Get retrieves the recipe given the recipeObjectID
func (r *Recipes) Get(ctx context.Context, recipeObjectID bson.ObjectId) (mealzpb.Recipe, error) {
	return nil, errors.New("unimplemented")
}

// Insert inserts the supplied recipe information into the datastore
func (r *Recipes) Insert(ctx context.Context, recipe mealzpb.Recipe) (mealzpb.Recipe, error) {
	return nil, errors.New("unimplemented")
}

// Update updates the supplied recipe information in the database
func (r *Recipes) Update(ctx context.Context, recipe mealzpb.Recipe) (mealzpb.Recipe, error) {
	return nil, errors.New("unimplemented")
}
