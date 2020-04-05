package config

import (
	"errors"
	"os"
)

const (
	// KeyPostgresHost defines the key in the environment for the postgres host
	KeyPostgresHost = "POSTGRES_HOST"

	// KeyPostgresUser defines the key in the environment for the postgres user
	KeyPostgresUser = "POSTGRES_USER"

	// KeyPostgresPassword defines the key in the environment for the postgres password
	KeyPostgresPassword = "POSTGRES_PASSWORD"

	// KeyPostgresSSLMode defines the key in the environment for the postgres sslmode
	KeyPostgresSSLMode = "POSTGRES_SSLMODE"

	// KeyPostgresDatabase defines the key in the environment for the postgres database
	KeyPostgresDatabase = "POSTGRES_DATABASE"
)

// PostgresConfig defines the configuration for a postgres database
type PostgresConfig struct {
	host     string
	user     string
	password string
	sslMode  string
	database string
}

// NewPostgresConfig returns a PostgresConfig from the environment, or an error if there is a problem
// loading any of the keys from the environment
func NewPostgresConfig() (*PostgresConfig, error) {
	c := new(PostgresConfig)
	var ok bool
	c.host, ok = os.LookupEnv(KeyPostgresHost)
	if !ok {
		return nil, errors.New("host not set")
	}

	c.user, ok = os.LookupEnv(KeyPostgresUser)
	if !ok {
		return nil, errors.New("user not set")
	}

	c.password, ok = os.LookupEnv(KeyPostgresPassword)
	if !ok {
		return nil, errors.New("password not set")
	}

	c.sslMode, ok = os.LookupEnv(KeyPostgresSSLMode)
	if !ok {
		return nil, errors.New("sslmode not set")
	}

	c.database, ok = os.LookupEnv(KeyPostgresDatabase)
	if !ok {
		return nil, errors.New("database not set")
	}

	return c, nil
}

// Host returns the configured host
func (c *PostgresConfig) Host() string { return c.host }

// User returns the configured user
func (c *PostgresConfig) User() string { return c.user }

// Password returns the configured password
func (c *PostgresConfig) Password() string { return c.password }

//SSLMode returns the configured SSLMode
func (c *PostgresConfig) SSLMode() string { return c.sslMode }

//Database returns the configured database
func (c *PostgresConfig) Database() string { return c.database }
