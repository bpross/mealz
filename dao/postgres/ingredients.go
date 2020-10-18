package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	dao "github.com/bpross/mealz/dao"
	"github.com/bpross/mealz/dao/postgres/queries"
	mealzpb "github.com/bpross/mealz/proto"
)

type pgUnitOfMeasure string

const (
	teaspoon   pgUnitOfMeasure = "TEASPOON"
	tablespoon pgUnitOfMeasure = "TABLESPOON"
	pound      pgUnitOfMeasure = "POUND"
	ounce      pgUnitOfMeasure = "OUNCE"
	cup        pgUnitOfMeasure = "CUP"
	piece      pgUnitOfMeasure = "PIECE"
)

// NewIngredients returns Ingredients with the supplied configurations
func NewIngredients(db *sql.DB, logger *log.Logger) *Ingredients {
	return &Ingredients{
		db:     db,
		logger: logger,
	}
}

// Ingredients implements the dao.Ingredients interface with a postgres backend
type Ingredients struct {
	db     *sql.DB
	logger *log.Logger
}

// Associate associates the supplied ingredients with the recipe
func (i *Ingredients) Associate(ctx context.Context, recipeObjectID bson.ObjectId, ingredients ...*dao.IngredientAssociation) error {
	if len(ingredients) == 0 {
		return fmt.Errorf("must specify at least one ingredient")
	}
	logger := i.logger.WithField("recipe_object_id", recipeObjectID.Hex())

	// Using a transaction here, given the unknown amount of ingredients. Ultimately its going to be low-ish
	// but better to safe
	var (
		err error
		tx  *sql.Tx
	)
	tx, err = i.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return
			}
		}
		err = tx.Commit()
		if err != nil {
			return
		}
	}()

	objB := []byte(recipeObjectID)
	for _, ingredientAssociation := range ingredients {
		logger.WithField("ingredient_object_id", ingredientAssociation.ObjectID.Hex()).Infoln("associating")
		_, err = i.db.ExecContext(
			ctx,
			queries.IngredientAssociate,
			[]byte(ingredientAssociation.ObjectID),
			objB,
			ingredientAssociation.UnitOfMeasure.String(),
			ingredientAssociation.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

// Insert inserts an ingredient record in the database
func (i *Ingredients) Insert(ctx context.Context, vegetarian bool, title string) (*mealzpb.Ingredient, error) {
	ingredientObjectID := bson.NewObjectId()
	logger := i.logger.WithField("ingredient_object_id", ingredientObjectID.Hex())
	logger.Info("inserting")

	var id int
	err := i.db.QueryRowContext(ctx, queries.IngredientInsert, ingredientObjectID, vegetarian, title).Scan(&id)
	if err != nil {
		errStr := fmt.Sprintf("failed to insert ingredient: %s", err.Error())
		return nil, errors.New(errStr)
	}

	ingredient := &mealzpb.Ingredient{
		ObjectId:   []byte(ingredientObjectID),
		Title:      title,
		Vegetarian: vegetarian,
	}

	return ingredient, nil
}

// Get retrieves an ingredient record from the database
func (i *Ingredients) Get(ctx context.Context, ingredientObjectID bson.ObjectId) (*mealzpb.Ingredient, error) {
	logger := i.logger.WithField("ingredient_object_id", ingredientObjectID.Hex())
	var id int
	ingredient := new(mealzpb.Ingredient)
	ingredient.ObjectId = []byte(ingredientObjectID)

	logger.Info("retrieving")
	err := i.db.QueryRowContext(ctx, queries.IngredientGet, ingredient.ObjectId).Scan(
		&id,
		&ingredient.Vegetarian,
		&ingredient.Title,
	)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown ingredient: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to get ingredient: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	return ingredient, nil
}

// Delete removes an ingredient record from the database
func (i *Ingredients) Delete(ctx context.Context, ingredientObjectID bson.ObjectId) (*mealzpb.Ingredient, error) {
	logger := i.logger.WithField("ingredient_object_id", ingredientObjectID.Hex())
	var id int
	ingredient := new(mealzpb.Ingredient)
	ingredient.ObjectId = []byte(ingredientObjectID)

	logger.Infoln("deleting")
	err := i.db.QueryRowContext(ctx, queries.IngredientDelete, ingredient.ObjectId).Scan(
		&id,
		&ingredient.Vegetarian,
		&ingredient.Title,
	)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown ingredient: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to delete ingredient: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	return ingredient, nil
}

func pgUnitOfMeasureToProto(u pgUnitOfMeasure) mealzpb.UnitOfMeasure {
	switch u {
	case teaspoon:
		return mealzpb.UnitOfMeasure_TEASPOON
	case tablespoon:
		return mealzpb.UnitOfMeasure_TABLESPOON
	case pound:
		return mealzpb.UnitOfMeasure_POUND
	case ounce:
		return mealzpb.UnitOfMeasure_OUNCE
	case cup:
		return mealzpb.UnitOfMeasure_CUP
	case piece:
		return mealzpb.UnitOfMeasure_PIECE
	default:
		return mealzpb.UnitOfMeasure_UNKNOWN_UNITOFMEASURE
	}
}
