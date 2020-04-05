package postgres

import (
	"database/sql"
	"fmt"
	// This is how it says to do it in the docs
	_ "github.com/lib/pq"

	"github.com/bpross/mealz/config"
)

const (
	// SSLDisable defines the disable SSLMode
	SSLDisable = "disable"
)

// Connect attempts to open and validate a connection to postgres with the provided configuration
func Connect(c *config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		c.User(),
		c.Password(),
		c.Host(),
		c.Database(),
		c.SSLMode(),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
