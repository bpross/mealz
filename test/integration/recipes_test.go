// +build integration

package integration

import (
	"context"
	"database/sql"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	"github.com/bpross/mealz/config"
	dao "github.com/bpross/mealz/dao/postgres"
	"github.com/bpross/mealz/postgres"
	mealzpb "github.com/bpross/mealz/proto"
)

var _ = Describe("Recipes Integration", func() {
	var (
		ctx    context.Context
		c      *config.PostgresConfig
		logger *log.Logger
		db     *sql.DB
		r      *dao.Recipes
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		var configErr error
		c, configErr = config.NewPostgresConfig()
		Expect(configErr).To(BeNil())

		var dbErr error
		db, dbErr = postgres.Connect(c)
		Expect(dbErr).To(BeNil())

		r = dao.NewRecipes(db, logger)
	})

	AfterEach(func() {
		db.Close()
	})

	Context("Delete", func() {
		var (
			recipeObjectID bson.ObjectId
			recipe         *mealzpb.Recipe
			err            error
		)

		JustBeforeEach(func() {
			recipe, err = r.Delete(ctx, recipeObjectID)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				recipeObjectID = bson.NewObjectId()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
			})
		})

		Context("with recipe that does exist", func() {
			var retRecipe *mealzpb.Recipe

			BeforeEach(func() {
				insertRecipe := &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}

				var insertErr error
				retRecipe, insertErr = r.Insert(ctx, insertRecipe)
				Expect(insertErr).To(BeNil())
				Expect(retRecipe).To(Equal(insertRecipe))

				recipeObjectID = bson.ObjectId(retRecipe.ObjectId)
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a recipe", func() {
				Expect(recipe).To(Equal(retRecipe))
			})
		})
	})

	Context("Get", func() {
		var (
			recipeObjectID bson.ObjectId
			recipe         *mealzpb.Recipe
			err            error
		)

		JustBeforeEach(func() {
			recipe, err = r.Get(ctx, recipeObjectID)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				recipeObjectID = bson.NewObjectId()
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return a recipe", func() {
				Expect(recipe).To(BeNil())
			})
		})

		Context("with recipe that does exist", func() {
			var retRecipe *mealzpb.Recipe

			BeforeEach(func() {
				insertRecipe := &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}

				var insertErr error
				retRecipe, insertErr = r.Insert(ctx, insertRecipe)
				Expect(insertErr).To(BeNil())
				Expect(retRecipe).To(Equal(insertRecipe))

				recipeObjectID = bson.ObjectId(retRecipe.ObjectId)
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a recipe", func() {
				Expect(recipe).To(Equal(retRecipe))
			})
		})
	})

	Context("Insert", func() {
		var (
			recipe, retRecipe *mealzpb.Recipe
			err               error
		)

		JustBeforeEach(func() {
			retRecipe, err = r.Insert(ctx, recipe)
		})

		Context("with unique object_id", func() {
			BeforeEach(func() {
				recipe = &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a recipe", func() {
				Expect(retRecipe).To(Equal(recipe))
			})
		})

		Context("with bad ethnicity", func() {
			BeforeEach(func() {
				recipe = &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_UNKNOWN_ETHNICITY,
					Source:     "test-source",
				}
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return a recipe", func() {
				Expect(retRecipe).To(BeNil())
			})
		})
	})

	Context("Update", func() {
		var (
			recipe, retRecipe *mealzpb.Recipe
			err               error
		)

		JustBeforeEach(func() {
			retRecipe, err = r.Update(ctx, recipe)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				recipe = &mealzpb.Recipe{
					ObjectId:   []byte(bson.NewObjectId()),
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should NOT return a recipe", func() {
				Expect(retRecipe).To(BeNil())
			})
		})

		Context("with recipe that does exist", func() {
			BeforeEach(func() {
				insertRecipe := &mealzpb.Recipe{
					Title:      "test-recipe",
					Vegetarian: true,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}

				var insertErr error
				retRecipe, insertErr = r.Insert(ctx, insertRecipe)
				Expect(insertErr).To(BeNil())
				Expect(retRecipe).To(Equal(insertRecipe))

				recipe = &mealzpb.Recipe{
					ObjectId:   insertRecipe.ObjectId,
					Title:      "NEW RECIPE",
					Vegetarian: false,
					Ethnicity:  mealzpb.Ethnicity_AMERICAN,
					Source:     "test-source",
				}
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a recipe", func() {
				Expect(recipe).To(Equal(retRecipe))
			})
		})
	})
})
