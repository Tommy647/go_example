package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Tommy647/go_example/internal/router"
)

type Store struct {
	db *sql.DB
}

func (s Store) GetRouterByID(id string) (router.Router, error) {
	var rtr router.Router

	row := s.db.QueryRow(
		`SELECT id, hostname, vendor, mgmtip FROM routers WHERE id=$1;`,
		id,
	)
	err := row.Scan(&rtr.ID, &rtr.Hostname, &rtr.Vendor, &rtr.MgmtIP)
	if err != nil {
		log.Println(err.Error())
		return router.Router{}, err
	}
	return rtr, nil
}

func (s Store) InsertRouter(rtr router.Router) (router.Router, error) {
	_, err := s.db.Exec(
		"INSERT INTO routers (id, hostname, vendor, mgmtip) VALUES ($1, $2, $3, $4)",
		rtr.ID, rtr.Hostname, rtr.Vendor, rtr.MgmtIP,
	)
	if err != nil {
		log.Println(err.Error())
		return router.Router{}, err
	}

	return rtr, nil
}

func (s Store) DeleteRouter(id string) error {
	_, err := s.db.Exec(
		`DELETE FROM routers WHERE id=$1`, id,
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
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
