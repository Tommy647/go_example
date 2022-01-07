// Package dbgreeter implements our greeter interface using a database to replace names if we get a match
package dbgreeter

import (
	"context"
	"database/sql"
	"log"

	"go.uber.org/zap"

	"github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/logger"
)

const (
	// query to get a name replacement
	query = `SELECT "to" FROM "public"."name" WHERE "from" = $1 LIMIT 1`

	// query to get the price of the requested coffee
	queryCoffee = `SELECT "price" FROM "public"."coffee" WHERE "coffee_type" = $1 LIMIT 1`

	// query to get the unit_price of the requested fruit
	queryFruit = `SELECT "unit_price" FROM "public"."fruit" WHERE "name" = $1 LIMIT 1`
)

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
	logger.Info(ctx, "database greet called", zap.String("in", in))
	// create an instance of our basic greeter to reuse
	basicGreeter := greeter.New()
	rows, err := g.db.QueryContext(ctx, query, in)
	if err != nil {
		// logger out the error and continue with the default behaviour
		logger.Error(ctx, "greet query", zap.Error(err))
		log.Println(">>> Failed to get the DB")
		return basicGreeter.Greet(ctx, in)
	}

	// placeholder for database value
	to := in
	// While we have rows - we are only expecting one
	for rows.Next() {
		// scan the data from our row into our placeholder
		if err := rows.Scan(&to); err != nil {
			logger.Error(ctx, "greet scan", zap.Error(err))
			// if rows.Scan errors, we need to close
			if err := rows.Close(); err != nil {
				logger.Error(ctx, "greet row close", zap.Error(err))
			}
			return basicGreeter.Greet(ctx, in)
		}
	}

	// no need to rows.Close if rows.Next returned false, just check for errors
	if err := rows.Err(); err != nil {
		logger.Error(ctx, "greet row", zap.Error(err))
		return basicGreeter.Greet(ctx, in)
	}

	// use our original greeter to handle the final string
	return basicGreeter.Greet(ctx, to)
}

// CoffeeGreet takes the context and a type of coffee to serve its price if found on DB
// or server one free otherwise
func (g *DBGreeter) CoffeeGreet(ctx context.Context, coffeeType string) string {
	logger.Info(ctx, "database coffeeGreet called", zap.String("in", coffeeType))
	// We grab a new instance of a BasicGreeter
	basicGreeter := greeter.New()

	rows, err := g.db.QueryContext(ctx, queryCoffee, coffeeType)
	if err != nil { // If we don't find the requested kind of coffee we return a basicGreeter
		logger.Error(ctx, "coffeeGreet query", zap.Error(err))
		log.Println("sorry can't get you a coffee from db")
		return basicGreeter.CoffeeGreet(ctx, coffeeType)
	}
	// placeholder for value from DB
	out := coffeeType
	for rows.Next() {
		if err := rows.Scan(&out); err != nil { // If there is a problem scanning the rows we return a basicGreeter
			logger.Error(ctx, "coffeeGreet scan", zap.Error(err))
			if err = rows.Close(); err != nil {
				logger.Error(ctx, "greet row close", zap.Error(err))
			}
			return basicGreeter.CoffeeGreet(ctx, coffeeType)
		}
		log.Println("coffee found in db")
		return out // We return with the value from the DB
	}
	// no need to rows.Close if rows.Next returned false, just check for errors
	if err := rows.Err(); err != nil {
		logger.Error(ctx, "coffeeGreet row", zap.Error(err))
		return basicGreeter.CoffeeGreet(ctx, coffeeType)
	}
	// We return with a basicGreeter
	return basicGreeter.CoffeeGreet(ctx, coffeeType)
}

func (g *DBGreeter) FruitGreet(ctx context.Context, fruit string) string {
	logger.Info(ctx, "building the response for a Fruit request")
	// We grab a new instance of a BasicGreeter
	basicGreeter := greeter.New()
	rows, err := g.db.QueryContext(ctx, queryFruit, fruit)
	if err != nil {
		logger.Error(ctx, "fruitGreet query", zap.Error(err))
		log.Println("sorry can't get you that fruit from DB")
		return basicGreeter.FruitGreet(ctx, fruit)
	}

	out := fruit
	for rows.Next() {
		if err := rows.Scan(&out); err != nil {
			logger.Error(ctx, "fruitGreet scan", zap.Error(err))
			if err = rows.Close(); err != nil {
				logger.Error(ctx, "FruitGreet row close", zap.Error(err))
			}
			return basicGreeter.FruitGreet(ctx, fruit)
		}
		log.Println("fruit found in DB")
		return out
	}
	if err := rows.Err(); err != nil {
		logger.Error(ctx, "fruitGreet row", zap.Error(err))
	}
	return basicGreeter.FruitGreet(ctx, fruit)
}
