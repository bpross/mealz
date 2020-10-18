package dao

import (
	"context"

	"labix.org/v2/mgo/bson"

	mealzpb "github.com/bpross/mealz/proto"
)

// IngredientAssociation represents data used to associate an ingredient with a recipe
type IngredientAssociation struct {
	ObjectID      bson.ObjectId
	UnitOfMeasure mealzpb.UnitOfMeasure
	Amount        float64
}

// Ingredients defines the interface a datastore must implement to support ingredients
type Ingredients interface {
	Associate(context.Context, bson.ObjectId, ...*IngredientAssociation) error
	Insert(context.Context, bool, string) (*mealzpb.Ingredient, error)
	Get(context.Context, bson.ObjectId) (*mealzpb.Ingredient, error)
	Delete(context.Context, bson.ObjectId) (*mealzpb.Ingredient, error)
}
