// Package dbgreeter implements our greeter interface using a database to replace names if we get a match
package dbgreeter

import (
	"context"
	"database/sql"
	"log"

	"github.com/Tommy647/go_example/internal/greeter"
)

// query to get a name replacement
const query = `SELECT "to" FROM "public"."name" WHERE "from" = $1 LIMIT 1`

// New returns a new instance of our database greeter
func New(db *sql.DB) *DBGreeter {
	return &DBGreeter{
		db: db,
	}
}

// DBGreeter our database greeter
type DBGreeter struct {
	db *sql.DB
}

// Greet provides our hello request, checks the DB to see
// if `in` exists, and replaces with the DB.from
func (g *DBGreeter) Greet(ctx context.Context, in string) string {
	// create an instance of our basic greeter to reuse
	basicGreeter := greeter.New()
	rows, err := g.db.QueryContext(ctx, query, in)
	if err != nil {
		// log out the error and continue with the default behaviour
		log.Println("query error", err.Error())
		return basicGreeter.Greet(ctx, in)
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
			return basicGreeter.Greet(ctx, in)
		}
	}

	// no need to rows.Close if rows.Next returned false, just check for errors
	if err := rows.Err(); err != nil {
		log.Println("row error", err.Error())
		return basicGreeter.Greet(ctx, in)
	}

	// use our original greeter to handle the final string
	return basicGreeter.Greet(ctx, to)
}
