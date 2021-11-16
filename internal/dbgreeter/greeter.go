// Package dbgreeter implements our greeter interface using a database to replace names if we get a match
package dbgreeter

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Tommy647/go_example/internal/greeter"
)

const (
	// query to get a name replacement
	query       = `SELECT "to" FROM "public"."name" WHERE "from" = $1 LIMIT 1`
	queryCoffee = `SELECT "price" FROM "public"."coffee" WHERE "type" = $1 LIMIT 1`
)

const (
	// environment variable names
	envGreeter = `GREETER`     // which greeter to use
	dbHost     = `DB_HOST`     // database host
	dbPort     = `DB_PORT`     // database port
	dbUser     = `DB_USER`     // database user
	dbPassword = `DB_PASSWORD` // database password
	dbDbname   = `DB_DBNAME`   // database name
)

// Greet our database greeter
type Greet struct {
	db *sql.DB
}

// New returns a new instance of our database greeter
func New(db *sql.DB) *Greet {
	return &Greet{
		db: db,
	}
}

// Greet provides our hello request, checks the DB to see
// if `in` exists, and replaces with the DB.from
func (g *Greet) Greet(ctx context.Context, in string) string {
	rows, err := g.db.QueryContext(ctx, query, in)
	if err != nil {
		// log out the error and continue with the default behaviour
		log.Println("query error", err.Error())
		return (greeter.Greet{}).Greet(ctx, in)
	}
	// placeholder for database value
	to := in
	// While we have rows - we are only expecting one
	for rows.Next() {
		// scan the data from our row into our placeholder
		if err := rows.Scan(&to); err != nil {
			log.Println("scan error", err.Error())
			// if rows.Scan errors, we need to close
			if err := rows.Close(); err != nil {
				log.Println("row close error", err.Error())
			}
			return (greeter.Greet{}).Greet(ctx, in)
		}
	}

	// no need to rows.Close if rows.Next returned false, just check for errors
	if err := rows.Err(); err != nil {
		log.Println("row error", err.Error())
		return (greeter.Greet{}).Greet(ctx, in)
	}

	// use our original greeter to handle the final string
	return (greeter.Greet{}).Greet(ctx, to)
}

// CoffeeGreet provides our coffee request, looks for `in` in the DB and gets the price
// if that kind of coffee exists, otherwise an error message is returned
func (g *Greet) CoffeeGreet(ctx context.Context, in string) string {

/*	db, err := sql.Open("postgres", getPostgresConnection())
	if err != nil {
		panic("database" + err.Error())
	}*/

	rows, err := g.db.QueryContext(ctx, queryCoffee, in)
	if err != nil {
		// log out the error and continue with the default behaviour
		log.Println("query error", err.Error())
		return (greeter.Greet{}).CoffeeGreet(ctx, in)
	}
	// placeholder for database value
	var out = "initial value"
	// While we have rows - we are only expecting one
	for rows.Next() {
		// scan the data from our row into our placeholder
		if err := rows.Scan(&out); err != nil {
			log.Println("scan error", err.Error())
			// if rows.Scan errors, we need out close
			if err := rows.Close(); err != nil {
				log.Println("row close error", err.Error())
			}
			return (greeter.Greet{}).CoffeeGreet(ctx, in)
		}
	}

	// no need out rows.Close if rows.Next returned false, just check for errors
	if err := rows.Err(); err != nil {
		log.Println("row error", err.Error())
		return (greeter.Greet{}).CoffeeGreet(ctx, in)
	}

	// use our original greeter out handle the final string
	return out
}

// getPostgresConnection string we need to open the connection
// gets connection details from the environment for now @todo: replace with viper
func getPostgresConnection() string {
	host := os.Getenv(dbHost)
	port, err := strconv.Atoi(os.Getenv(dbPort))
	if err != nil {
		panic("port must be a number " + err.Error())
	}
	user := os.Getenv(dbUser)
	password := os.Getenv(dbPassword)
	dbname := os.Getenv(dbDbname)

	return fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname,
	)
}
