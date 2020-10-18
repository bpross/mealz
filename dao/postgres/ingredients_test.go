package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	"github.com/bpross/mealz/dao"
	"github.com/bpross/mealz/dao/postgres/queries"
	mealzpb "github.com/bpross/mealz/proto"
)

var _ = Describe("Ingredients", func() {
	var (
		ctx    context.Context
		logger *log.Logger
		db     *sql.DB
		mock   sqlmock.Sqlmock
		i      *Ingredients
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		var dbErr error
		db, mock, dbErr = sqlmock.New()
		Expect(dbErr).To(BeNil())

		i = NewIngredients(db, logger)
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Associate", func() {
		var (
			recipeObjectID bson.ObjectId
			ingredients    []*dao.IngredientAssociation

			err error
		)

		BeforeEach(func() {
			recipeObjectID = bson.NewObjectId()
			ingredients = make([]*dao.IngredientAssociation, 0)
		})

		JustBeforeEach(func() {
			err = i.Associate(ctx, recipeObjectID, ingredients...)
		})

		Context("without ingredients", func() {
			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("must specify at least one ingredient"))
			})
		})

		Context("with ingredients", func() {
			var mockBegin *sqlmock.ExpectedBegin
			BeforeEach(func() {
				ingredient := &dao.IngredientAssociation{
					ObjectID:      bson.NewObjectId(),
					UnitOfMeasure: mealzpb.UnitOfMeasure_TABLESPOON,
					Amount:        1.0,
				}

				ingredients = append(ingredients, ingredient)

				mockBegin = mock.ExpectBegin()
			})

			Context("with Begin error", func() {
				BeforeEach(func() {
					mockBegin.WillReturnError(errors.New("test-error"))
				})

				It("should return an error", func() {
					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("test-error"))
				})
			})

			Context("with Begin success", func() {
				var mockAssociate *sqlmock.ExpectedExec
				BeforeEach(func() {
					ingredient := ingredients[0]
					mockAssociate = mock.ExpectExec(PrepareQueryRegex(queries.IngredientAssociate)).WithArgs([]byte(ingredient.ObjectID), []byte(recipeObjectID), ingredient.UnitOfMeasure.String(), ingredient.Amount)
				})

				Context("with associate error", func() {
					BeforeEach(func() {
						mockAssociate.WillReturnError(fmt.Errorf("test-error"))
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
						Expect(err.Error()).To(Equal("test-error"))
					})
				})

				Context("with associate success", func() {
					BeforeEach(func() {
						lastInsertID := int64(42)
						rowsAffected := int64(1)
						result := sqlmock.NewResult(lastInsertID, rowsAffected)
						mockAssociate.WillReturnResult(result)
					})

					It("should NOT return an error", func() {
						Expect(err).To(BeNil())
					})
				})
			})
		})
	}) // Associate

	Describe("Insert", func() {
		var (
			vegetarian bool
			title      string

			result *mealzpb.Ingredient
			err    error

			mockInsert *sqlmock.ExpectedQuery
		)
		BeforeEach(func() {
			vegetarian = true
			title = "test ingredient"
			mockInsert = mock.ExpectQuery(PrepareQueryRegex(queries.IngredientInsert)).WithArgs(BeObjectID(), vegetarian, title)
		})

		JustBeforeEach(func() {
			result, err = i.Insert(ctx, vegetarian, title)
		})

		Context("with Insert failure", func() {
			BeforeEach(func() {
				mockInsert.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should NOT return result", func() {
				Expect(result).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("failed to insert ingredient: test-error"))
			})
		})

		Context("with Insert success", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(42)
				mockInsert.WillReturnRows(rows)
			})

			It("should return an ingredient", func() {
				Expect(result).NotTo(BeNil())
				Expect(result.Title).To(Equal(title))
				Expect(result.Vegetarian).To(Equal(vegetarian))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get", func() {
		var (
			ingredientObjectID bson.ObjectId
			ingredient         *mealzpb.Ingredient
			err                error

			mockGet *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			ingredientObjectID = bson.NewObjectId()
			mockGet = mock.ExpectQuery(PrepareQueryRegex(queries.IngredientGet)).
				WithArgs([]byte(ingredientObjectID))
		})

		JustBeforeEach(func() {
			ingredient, err = i.Get(ctx, ingredientObjectID)
		})

		Context("with an ingredient that does not exist", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(sql.ErrNoRows)
			})

			It("should not return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown ingredient: sql: no rows in result set"))
			})
		})

		Context("with an unknown error", func() {
			BeforeEach(func() {
				mockGet.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should not return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to get ingredient: test-error"))
			})
		})

		Context("with get success", func() {
			var retIngredient *mealzpb.Ingredient
			BeforeEach(func() {
				retIngredient = &mealzpb.Ingredient{
					ObjectId:   []byte(ingredientObjectID),
					Title:      "test-ingredient",
					Vegetarian: true,
				}
				rows := sqlmock.NewRows([]string{"id", "vegetarian", "title"}).
					AddRow(42, "true", "test-ingredient")
				mockGet.WillReturnRows(rows)
			})

			It("should return an ingredient", func() {
				Expect(ingredient).To(Equal(retIngredient))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		var (
			ingredientObjectID bson.ObjectId
			ingredient         *mealzpb.Ingredient
			err                error

			mockDelete *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
			ingredientObjectID = bson.NewObjectId()
			mockDelete = mock.ExpectQuery(PrepareQueryRegex(queries.IngredientDelete)).
				WithArgs([]byte(ingredientObjectID))
		})

		JustBeforeEach(func() {
			ingredient, err = i.Delete(ctx, ingredientObjectID)
		})

		Context("with an ingredient that does not exist", func() {
			BeforeEach(func() {
				mockDelete.WillReturnError(sql.ErrNoRows)
			})

			It("should not return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unknown ingredient: sql: no rows in result set"))
			})
		})

		Context("with an unknown error", func() {
			BeforeEach(func() {
				mockDelete.WillReturnError(fmt.Errorf("test-error"))
			})

			It("should not return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("unable to delete ingredient: test-error"))
			})
		})

		Context("with get success", func() {
			var retIngredient *mealzpb.Ingredient
			BeforeEach(func() {
				retIngredient = &mealzpb.Ingredient{
					ObjectId:   []byte(ingredientObjectID),
					Vegetarian: true,
					Title:      "test-ingredient",
				}
				rows := sqlmock.NewRows([]string{"id", "vegetarian", "title"}).
					AddRow(42, "true", "test-ingredient")
				mockDelete.WillReturnRows(rows)
			})

			It("should return an ingredient", func() {
				Expect(ingredient).To(Equal(retIngredient))
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
