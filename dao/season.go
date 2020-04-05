package dao

import (
	"context"

	"labix.org/v2/mgo/bson"

	mealzpb "github.com/bpross/mealz/proto"
)

// Seasons defines the interface a datastore must implement to support seasons
type Seasons interface {
	Associate(context.Context, bson.ObjectId, []string) error
	Get(context.Context, string) (int, error)
	GetRecipeSeasons(context.Context, bson.ObjectId) ([]mealzpb.Season, error)
}
