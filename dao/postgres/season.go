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

type season string

const (
	winter season = "winter"
	spring season = "spring"
	summer season = "summer"
	fall   season = "fall"
)

// NewSeasons returns Seasons with the supplied configurations
func NewSeasons(db *sql.DB, logger *log.Logger) *Seasons {
	return &Seasons{
		db:     db,
		logger: logger,
	}
}

// Seasons implements the dao.Seasons interface with a postgres backend
type Seasons struct {
	db     *sql.DB
	logger *log.Logger
}

// Associate adds seasons to recipes
func (s *Seasons) Associate(ctx context.Context, recipeObjectID bson.ObjectId, seasons []string) error {
	if len(seasons) == 0 {
		return fmt.Errorf("must specify at least one season")
	}

	logger := s.logger.WithField("recipe_object_id", recipeObjectID)

	var id int

	// This could go in a transaction, but the number of seasons is fixed at 4, so not really worth it
	objB := []byte(recipeObjectID)
	for _, season := range seasons {
		logger.WithField("season", season).Infoln("associating")
		err := s.db.QueryRowContext(ctx, queries.SeasonAssociate, objB, season).Scan(&id)
		if err != nil {
			var errStr string
			if err == sql.ErrNoRows {
				errStr = fmt.Sprintf("unknown season: %s or recipe: %s", season, err.Error())
			} else {
				errStr = fmt.Sprintf("unable to associate season to recipe: %s", err.Error())
			}
			return errors.New(errStr)
		}
	}

	return nil
}

// Get returns the id of the season
func (s *Seasons) Get(ctx context.Context, season string) (int, error) {
	logger := s.logger.WithField("season", season)
	var id int

	logger.Info("retrieving")
	err := s.db.QueryRowContext(ctx, queries.SeasonGet, season).Scan(&id)

	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown season: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to retrieve season: %s", err.Error())
		}
		return 0, errors.New(errStr)
	}

	return id, nil
}

// GetRecipeSeasons returns the seasons for the recipe represented by the object_id
func (s *Seasons) GetRecipeSeasons(ctx context.Context, recipeObjectID bson.ObjectId) ([]mealzpb.Season, error) {
	logger := s.logger.WithField("recipe_object_id", recipeObjectID)
	logger.Infoln("getting seasons")

	objB := []byte(recipeObjectID)
	rows, err := s.db.QueryContext(ctx, queries.GetRecipeSeasons, objB)
	if err != nil {
		var errStr string
		if err == sql.ErrNoRows {
			errStr = fmt.Sprintf("unknown recipe: %s", err.Error())
		} else {
			errStr = fmt.Sprintf("unable to get seasons for recipe: %s", err.Error())
		}
		return nil, errors.New(errStr)
	}

	defer rows.Close()
	seasons := make([]mealzpb.Season, 0)
	for rows.Next() {
		var (
			id int
			s  season
		)
		err = rows.Scan(&id, &s)
		if err != nil {
			errStr := fmt.Sprintf("error reading season row: %s", err.Error())
			return nil, errors.New(errStr)
		}

		seasons = append(seasons, (pgSeasonToProto(s)))
	}

	return seasons, nil
}

func pgSeasonToProto(s season) mealzpb.Season {
	switch s {
	case winter:
		return mealzpb.Season_WINTER
	case spring:
		return mealzpb.Season_SPRING
	case summer:
		return mealzpb.Season_SUMMER
	case fall:
		return mealzpb.Season_FALL
	default:
		return mealzpb.Season_UNKNOWN_SEASON
	}
}
