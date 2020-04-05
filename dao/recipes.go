package dao

import (
	"context"

	"labix.org/v2/mgo/bson"

	mealzpb "github.com/bpross/mealz/proto"
)

// Recipes defines the interface that a datstore must implement to support recipes
type Recipes interface {
	Delete(context.Context, bson.ObjectId) (*mealzpb.Recipe, error)
	Get(context.Context, bson.ObjectId) (*mealzpb.Recipe, error)
	Insert(context.Context, *mealzpb.Recipe) (*mealzpb.Recipe, error)
	Update(context.Context, *mealzpb.Recipe) (*mealzpb.Recipe, error)
}
