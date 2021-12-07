package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // special: we need to include this package here to ensure the drivers load, but we do not need the code
)

const (
	// environment variables
	dbHost     = `DB_HOST`     // database host
	dbPort     = `DB_PORT`     // database port
	dbUser     = `DB_USER`     // database user
	dbPassword = `DB_PASSWORD` // database password
	dbDbname   = `DB_DBNAME`   // database name
)

// NewConnection to a postgres database
func NewConnection() (*sql.DB, error) {
	return sql.Open("postgres", getPostgresConnection())
}

// getPostgresConnection string we need to open the connection
// gets connection details from the environment for now @todo: replace with viper
func getPostgresConnection() string {
	return fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv(dbHost),
		os.Getenv(dbPort),
		os.Getenv(dbUser),
		os.Getenv(dbPassword),
		os.Getenv(dbDbname),
	)
}
