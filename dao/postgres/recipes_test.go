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

var _ = Describe("Recipes", func() {
	var (
		ctx    context.Context
		logger *log.Logger
		db     *sql.DB
		mock   sqlmock.Sqlmock
		r      *Recipes
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		var dbErr error
		db, mock, dbErr = sqlmock.New()
		Expect(dbErr).To(BeNil())

		r = NewRecipes(db, logger)
	})

	AfterEach(func() {
		db.Close()
	})

	Context("Insert", func() {
		var (
			recipe    *mealzpb.Recipe
			retRecipe *mealzpb.Recipe
			err       error
		)

		JustBeforeEach(func() {
			retRecipe, err = r.Insert(ctx, recipe)
		})

		Context("with supplied object_id", func() {
			BeforeEach(func() {
				recipe = &mealzpb.Recipe{
					ObjectId: []byte("test"),
				}
			})

			It("should NOT return a recipe", func() {
				Expect(retRecipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("cannot specify object_id on insert"))
			})
		})

		Context("without object_id", func() {
			var mockInsert *sqlmock.ExpectedQuery
			BeforeEach(func() {
				recipe = &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
				mockInsert = mock.ExpectQuery(PrepareQueryRegex(queries.RecipeInsert)).
					WithArgs(BeObjectID(), recipe.Title, recipe.Vegetarian, recipe.Ethnicity.String(), recipe.Source)
			})

			Context("with DB error", func() {
				BeforeEach(func() {
					mockInsert.WillReturnError(fmt.Errorf("test-error"))
				})

				It("should NOT return a recipe", func() {
					Expect(retRecipe).To(BeNil())
				})

				It("should return an error", func() {
					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("failed to insert recipe: test-error"))
				})
			})

			Context("with DB success", func() {
				BeforeEach(func() {
					rows := sqlmock.NewRows([]string{"id"}).AddRow(42)
					mockInsert.WillReturnRows(rows)
				})

				It("should return a recipe", func() {
					Expect(retRecipe).To(Equal(recipe))
				})

				It("should NOT return an error", func() {
					Expect(err).To(BeNil())
				})
			})
		})
	})

	Context("Get", func() {
		var (
			recipeObjectID bson.ObjectId
			recipe         *mealzpb.Recipe
			err            error

			mockGet *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			recipeObjectID = bson.NewObjectId()
			mockGet = mock.ExpectQuery(PrepareQueryRegex(queries.RecipeGet)).
				WithArgs([]byte(recipeObjectID))
		})

		JustBeforeEach(func() {
			recipe, err = r.Get(ctx, recipeObjectID)
		})

		Context("with recipeObjectID that does not exist", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(sql.ErrNoRows)
			})

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
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

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to retrieve recipe: test-error"))
			})
		})

		Context("with get success", func() {
			var retRecipe *mealzpb.Recipe
			BeforeEach(func() {
				retRecipe = &mealzpb.Recipe{
					ObjectId:   []byte(recipeObjectID),
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
				rows := sqlmock.NewRows([]string{"id", "title", "vegetarian", "ethnicity", "source"}).
					AddRow(42, "test-recipe", "true", "AMERICAN", "test-source")
				mockGet.WillReturnRows(rows)
			})

			It("should return a recipe", func() {
				Expect(retRecipe).To(Equal(recipe))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Update", func() {
		var (
			recipe    *mealzpb.Recipe
			retRecipe *mealzpb.Recipe
			err       error

			mockUpdate *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			recipe = &mealzpb.Recipe{
				ObjectId:   []byte(bson.NewObjectId()),
				Title:      "test-recipe",
				Vegetarian: true,
				Ethnicity:  mealzpb.Ethnicity_AMERICAN,
				Source:     "test-source",
			}
			mockUpdate = mock.ExpectQuery(PrepareQueryRegex(queries.RecipeUpdate)).
				WithArgs(recipe.ObjectId, recipe.Title, recipe.Vegetarian, recipe.Ethnicity.String(), recipe.Source)
		})

		JustBeforeEach(func() {
			retRecipe, err = r.Update(ctx, recipe)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				mockUpdate.WillReturnError(sql.ErrNoRows)
			})

			It("should NOT return a recipe", func() {
				Expect(retRecipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown recipe: sql: no rows in result set"))
			})
		})

		Context("with unknown error", func() {
			BeforeEach(func() {
				mockUpdate.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should NOT return a recipe", func() {
				Expect(retRecipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to update recipe: test-error"))
			})
		})

		Context("with update success", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(42)
				mockUpdate.WillReturnRows(rows)
			})

			It("should return a recipe", func() {
				Expect(retRecipe).To(Equal(recipe))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Delete", func() {
		var (
			recipeObjectID bson.ObjectId
			recipe         *mealzpb.Recipe
			err            error

			mockDelete *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			recipeObjectID = bson.NewObjectId()
			mockDelete = mock.ExpectQuery(PrepareQueryRegex(queries.RecipeDelete)).
				WithArgs([]byte(recipeObjectID))
		})

		JustBeforeEach(func() {
			recipe, err = r.Delete(ctx, recipeObjectID)
		})

		Context("with recipeObjectID that does not exist", func() {
			BeforeEach(func() {
				mockDelete.WillReturnError(sql.ErrNoRows)
			})

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown recipe: sql: no rows in result set"))
			})
		})

		Context("with unknown error", func() {
			BeforeEach(func() {
				mockDelete.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to delete recipe: test-error"))
			})
		})

		Context("with delete success", func() {
			var retRecipe *mealzpb.Recipe
			BeforeEach(func() {
				retRecipe = &mealzpb.Recipe{
					ObjectId:   []byte(recipeObjectID),
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
				rows := sqlmock.NewRows([]string{"id", "title", "vegetarian", "ethnicity", "source"}).
					AddRow(42, "test-recipe", "true", "AMERICAN", "test-source")
				mockDelete.WillReturnRows(rows)
			})

			It("should return a recipe", func() {
				Expect(retRecipe).To(Equal(recipe))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
