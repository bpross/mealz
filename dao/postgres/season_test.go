package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	"github.com/bpross/mealz/dao/postgres/queries"
	mealzpb "github.com/bpross/mealz/proto"
)

var _ = Describe("Seasons", func() {
	var (
		ctx    context.Context
		logger *log.Logger
		db     *sql.DB
		mock   sqlmock.Sqlmock
		s      *Seasons
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		var dbErr error
		db, mock, dbErr = sqlmock.New()
		Expect(dbErr).To(BeNil())

		s = NewSeasons(db, logger)
	})

	AfterEach(func() {
		db.Close()
	})

	Context("Associate", func() {
		var (
			seasons        []string
			recipeObjectID bson.ObjectId
			err            error
			mockAssociates map[string]*sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			seasons = []string{"winter", "spring", "summer", "fall"}
			recipeObjectID = bson.NewObjectId()

			rObj := []byte(recipeObjectID)
			mockAssociates = make(map[string]*sqlmock.ExpectedQuery)
			for _, season := range seasons {
				mockAssociates[season] = mock.ExpectQuery(PrepareQueryRegex(queries.SeasonAssociate)).WithArgs(rObj, season)
			}
		})

		JustBeforeEach(func() {
			err = s.Associate(ctx, recipeObjectID, seasons)
		})

		Context("without seasons", func() {
			BeforeEach(func() {
				seasons = []string{}
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("must specify at least one season"))
			})
		})

		Context("with season that does not exist", func() {
			BeforeEach(func() {
				mockAssociates["winter"].WillReturnError(sql.ErrNoRows)
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown season: winter or recipe: sql: no rows in result set"))
			})
		})

		Context("with unknown error", func() {
			BeforeEach(func() {
				mockAssociates["winter"].WillReturnError(fmt.Errorf("test-error"))
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to associate season to recipe: test-error"))
			})
		})

		Context("with one success and one failure", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(42)
				mockAssociates["winter"].WillReturnRows(rows)
				mockAssociates["spring"].WillReturnError(fmt.Errorf("test-error"))
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to associate season to recipe: test-error"))
			})
		})

		Context("with all success", func() {
			BeforeEach(func() {
				winterRows := sqlmock.NewRows([]string{"id"}).
					AddRow(42)
				springRows := sqlmock.NewRows([]string{"id"}).
					AddRow(43)
				summerRows := sqlmock.NewRows([]string{"id"}).
					AddRow(44)
				fallRows := sqlmock.NewRows([]string{"id"}).
					AddRow(45)
				mockAssociates["winter"].WillReturnRows(winterRows)
				mockAssociates["spring"].WillReturnRows(springRows)
				mockAssociates["summer"].WillReturnRows(summerRows)
				mockAssociates["fall"].WillReturnRows(fallRows)
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Get", func() {
		var (
			season string
			id     int
			err    error

			mockGet *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			season = "winter"
			mockGet = mock.ExpectQuery(PrepareQueryRegex(queries.SeasonGet)).
				WithArgs(season)
		})

		JustBeforeEach(func() {
			id, err = s.Get(ctx, season)
		})

		Context("with season that does not exist", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(sql.ErrNoRows)
			})

			It("should NOT return a season", func() {
				Expect(id).To(Equal(0))
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown season: sql: no rows in result set"))
			})
		})

		Context("with unknown error", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should NOT return a season", func() {
				Expect(id).To(Equal(0))
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to retrieve season: test-error"))
			})
		})

		Context("with get success", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(42)
				mockGet.WillReturnRows(rows)
			})

			It("should return a season", func() {
				Expect(id).To(Equal(42))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Context("GetRecipeSeasons", func() {
		var (
			recipeObjectID bson.ObjectId
			mockGet        *sqlmock.ExpectedQuery

			seasons []mealzpb.Season
			err     error
		)

		BeforeEach(func() {
			recipeObjectID = bson.NewObjectId()
			mockGet = mock.ExpectQuery(PrepareQueryRegex(queries.GetRecipeSeasons)).
				WithArgs([]byte(recipeObjectID))
		})

		JustBeforeEach(func() {
			seasons, err = s.GetRecipeSeasons(ctx, recipeObjectID)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(sql.ErrNoRows)
			})

			It("should NOT return seasons", func() {
				Expect(seasons).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown recipe: sql: no rows in result set"))
			})
		})

		Context("with unknown error", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should NOT return seasons", func() {
				Expect(seasons).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to get seasons for recipe: test-error"))
			})
		})

		Context("with no rows returned", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "title"})
				mockGet.WillReturnRows(rows)
			})

			It("should return empty seasons", func() {
				Expect(seasons).To(BeEmpty())
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("multiple rows returned", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "title"}).
					AddRow(1, "winter").
					AddRow(2, "spring").
					AddRow(3, "fall")
				mockGet.WillReturnRows(rows)
			})

			It("should return empty seasons", func() {
				expected := []mealzpb.Season{
					mealzpb.Season_WINTER,
					mealzpb.Season_SPRING,
					mealzpb.Season_FALL,
				}
				Expect(seasons).To(Equal(expected))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
