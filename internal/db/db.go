package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Tommy647/go_example/internal/router"
)

type Store struct {
	db *sql.DB
}

func (s Store) GetRouterByID(id string) (router.Router, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) InsertRouter(router router.Router) (router.Router, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) DeleteRouter(id string) error {
	//TODO implement me
	panic("implement me")
}

func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}
