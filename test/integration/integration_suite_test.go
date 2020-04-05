package integration

import (
	"database/sql"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bpross/mealz/config"
	"github.com/bpross/mealz/postgres"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var db *sql.DB

var _ = BeforeSuite(func() {
	var (
		configErr error
		dbErr     error
	)
	c, configErr := config.NewPostgresConfig()
	Expect(configErr).To(BeNil())

	db, dbErr = postgres.Connect(c)
	Expect(dbErr).To(BeNil())
})

var _ = AfterSuite(func() {
	db.Close()
})
