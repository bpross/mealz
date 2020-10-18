// +build integration

package integration

import (
	"context"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	dao "github.com/bpross/mealz/dao"
	postgres "github.com/bpross/mealz/dao/postgres"
	mealzpb "github.com/bpross/mealz/proto"
)

var _ = Describe("Ingredients Integration", func() {
	var (
		ctx    context.Context
		logger *log.Logger
		i      *postgres.Ingredients
		r      *postgres.Recipes
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		i = postgres.NewIngredients(db, logger)
		r = postgres.NewRecipes(db, logger)
	})

	Context("Insert", func() {
		var (
			vegetarian bool
			title      string

			ingredient *mealzpb.Ingredient
			err        error
		)

		JustBeforeEach(func() {
			ingredient, err = i.Insert(ctx, vegetarian, title)
		})

		Context("with unique object_id", func() {
			BeforeEach(func() {
				title = "test-ingredient"
				vegetarian = true
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return an ingredient", func() {
				Expect(ingredient.Vegetarian).To(Equal(vegetarian))
				Expect(ingredient.Title).To(Equal(title))
			})
		})
	})

	Context("Get", func() {
		var (
			vegetarian         bool
			title              string
			ingredientObjectID bson.ObjectId

			ingredient *mealzpb.Ingredient
			err        error
		)

		JustBeforeEach(func() {
			ingredient, err = i.Get(ctx, ingredientObjectID)
		})

		Context("with an ingredient that does not exist", func() {
			BeforeEach(func() {
				ingredientObjectID = bson.NewObjectId()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})
		})

		Context("with an ingredient that exists", func() {
			var retIngredient *mealzpb.Ingredient
			BeforeEach(func() {
				vegetarian = true
				title = "test-ingredient"

				var insertErr error
				retIngredient, insertErr = i.Insert(ctx, vegetarian, title)
				Expect(insertErr).To(BeNil())
				Expect(retIngredient.Vegetarian).To(Equal(vegetarian))
				Expect(retIngredient.Title).To(Equal(title))

				ingredientObjectID = bson.ObjectId(retIngredient.ObjectId)
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return an ingredient", func() {
				Expect(ingredient).To(Equal(retIngredient))
			})
		})
	})

	Context("Delete", func() {
		var (
			vegetarian         bool
			title              string
			ingredientObjectID bson.ObjectId

			ingredient *mealzpb.Ingredient
			err        error
		)

		JustBeforeEach(func() {
			ingredient, err = i.Delete(ctx, ingredientObjectID)
		})

		Context("with an ingredient that does not exist", func() {
			BeforeEach(func() {
				ingredientObjectID = bson.NewObjectId()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return an ingredient", func() {
				Expect(ingredient).To(BeNil())
			})
		})

		Context("with an ingredient that exists", func() {
			var retIngredient *mealzpb.Ingredient
			BeforeEach(func() {
				vegetarian = true
				title = "test-ingredient"

				var insertErr error
				retIngredient, insertErr = i.Insert(ctx, vegetarian, title)
				Expect(insertErr).To(BeNil())
				Expect(retIngredient.Vegetarian).To(Equal(vegetarian))
				Expect(retIngredient.Title).To(Equal(title))

				ingredientObjectID = bson.ObjectId(retIngredient.ObjectId)
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return an ingredient", func() {
				Expect(ingredient).To(Equal(retIngredient))
			})
		})
	})

	Context("Associate", func() {
		var (
			recipeObjectID bson.ObjectId
			ingredients    []*dao.IngredientAssociation
			err            error
		)

		JustBeforeEach(func() {
			err = i.Associate(ctx, recipeObjectID, ingredients...)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				recipeObjectID = bson.NewObjectId()
				ingredients = []*dao.IngredientAssociation{
					&dao.IngredientAssociation{
						ObjectID: bson.NewObjectId(),
					},
				}
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})

		Context("with a recipe that does exist", func() {
			BeforeEach(func() {
				insertRecipe := &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}

				retRecipe, insertErr := r.Insert(ctx, insertRecipe)
				Expect(insertErr).To(BeNil())
				Expect(retRecipe).To(Equal(insertRecipe))

				recipeObjectID = bson.ObjectId(retRecipe.ObjectId)
			})

			Context("with no ingredients", func() {
				BeforeEach(func() {
					ingredients = []*dao.IngredientAssociation{}
				})

				It("should return an error", func() {
					Expect(err).NotTo(BeNil())
				})
			})

			Context("with ingredients", func() {
				BeforeEach(func() {
					numIngredients := 6 // number of measurements
					uOfM := []mealzpb.UnitOfMeasure{
						mealzpb.UnitOfMeasure_TEASPOON,
						mealzpb.UnitOfMeasure_TABLESPOON,
						mealzpb.UnitOfMeasure_POUND,
						mealzpb.UnitOfMeasure_OUNCE,
						mealzpb.UnitOfMeasure_CUP,
						mealzpb.UnitOfMeasure_PIECE,
					}
					for k := 0; k < numIngredients; k++ {
						title := fmt.Sprintf("test-ingredient-%d", k)
						vegetarian := true
						ingredient, insertErr := i.Insert(ctx, vegetarian, title)
						Expect(insertErr).To(BeNil())
						ia := &dao.IngredientAssociation{
							ObjectID:      bson.ObjectId(ingredient.ObjectId),
							UnitOfMeasure: uOfM[k],
							Amount:        1.0,
						}
						ingredients = append(ingredients, ia)
					}
				})

				It("should NOT return an error", func() {
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
