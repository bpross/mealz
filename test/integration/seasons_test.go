package integration

import (
	"context"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"labix.org/v2/mgo/bson"

	dao "github.com/bpross/mealz/dao/postgres"
	mealzpb "github.com/bpross/mealz/proto"
)

var _ = Describe("Recipes Integration", func() {
	var (
		ctx    context.Context
		logger *log.Logger
		r      *dao.Recipes
		s      *dao.Seasons
	)

	BeforeEach(func() {
		ctx = context.Background()

		logger = log.New()
		logger.Out = ioutil.Discard

		r = dao.NewRecipes(db, logger)
		s = dao.NewSeasons(db, logger)
	})

	Context("Associate", func() {
		var (
			recipeObjectID bson.ObjectId
			seasons        []string
			err            error
		)

		JustBeforeEach(func() {
			err = s.Associate(ctx, recipeObjectID, seasons)
		})

		Context("with recipe that does not exist", func() {
			BeforeEach(func() {
				recipeObjectID = bson.NewObjectId()
				seasons = []string{"winter"}
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

				seasons = []string{"winter", "spring", "summer", "fall"}
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Get", func() {
		var (
			season string

			id  int
			err error
		)

		JustBeforeEach(func() {
			id, err = s.Get(ctx, season)
		})

		Context("with season that does not exist", func() {
			BeforeEach(func() {
				season = "test"
			})

			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			It("should return zero id", func() {
				Expect(id).To(Equal(0))
			})
		})

		Context("with a season that exists", func() {
			BeforeEach(func() {
				season = "winter"
			})

			It("should NOT return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return an non-zero id", func() {
				Expect(id).NotTo(Equal(0))
			})
		})
	})

	Context("GetRecipeSeasons", func() {
		var (
			recipeObjectID bson.ObjectId

			seasons []mealzpb.Season
			err     error
		)

		JustBeforeEach(func() {
			seasons, err = s.GetRecipeSeasons(ctx, recipeObjectID)
		})

		Context("with recipe that does not exist", func() {
			Context("with recipe that does not exist", func() {
				BeforeEach(func() {
					recipeObjectID = bson.NewObjectId()
				})

				It("should NOT return an error", func() {
					Expect(err).To(BeNil())
				})

				It("should NOT return seasons", func() {
					Expect(seasons).To(BeEmpty())
				})
			})
		})

		Context("with a recipe that exists", func() {
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

			Context("with no associations", func() {

				It("should not return an error", func() {
					Expect(err).To(BeNil())
				})

				It("should NOT return seasons", func() {
					Expect(seasons).To(BeEmpty())
				})
			})

			Context("with associations", func() {
				BeforeEach(func() {
					associateSeasons := []string{"winter", "spring", "summer", "fall"}
					associateErr := s.Associate(ctx, recipeObjectID, associateSeasons)
					Expect(associateErr).To(BeNil())
				})

				It("should not return an error", func() {
					Expect(err).To(BeNil())
				})

				It("should return seasons", func() {
					expected := []mealzpb.Season{
						mealzpb.Season_WINTER,
						mealzpb.Season_SPRING,
						mealzpb.Season_SUMMER,
						mealzpb.Season_FALL,
					}
					Expect(seasons).To(Equal(expected))
				})
			})
		})
	})
})
