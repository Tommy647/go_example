package dbgreeter

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGreet_Greet(t *testing.T) {
	// expectedQuery in our tests - note this is a regex not a string
	var expectedQuery = `^` + strings.ReplaceAll(query, "$", `\$`) + `$`

	tests := []struct {
		name   string
		in     string
		expect string
	}{
		{
			name:   "should return the given name if nothing in the database",
			in:     "Tom",
			expect: "Hello, Tom!",
		}, {
			name:   "should return a replaced name if we find one",
			in:     "Orson",
			expect: "Hello, Ralphy!",
		}, {
			name:   "should return the given name if we error in the query",
			in:     "Kurt",
			expect: "Hello, Kurt!",
		}, {
			name:   "should return the given name if we error in the rows",
			in:     "Cath",
			expect: "Hello, Cath!",
		}, {
			name:   "should return the given name if we error in the scan",
			in:     "Becca",
			expect: "Hello, Becca!",
		},
	}

	// create a mock database for our unit tests, we are doing this outside of our
	// db: is a mocked out DB instance
	// mock: allows us to create expectations and returns
	// err: if anything went wrong
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error()) // stop the tests
	}

	// mock out some database rows to return
	columns := []string{"to"}

	// set up the mocks
	// expect a query with the input "Tom" and return an empty row set - NOTE: the expectedSQL is regex, not a string
	mock.ExpectQuery(expectedQuery). // The query
						WithArgs("Tom").                         // which arguments are expected
						WillReturnRows(sqlmock.NewRows(columns)) // return no rows

	// expect a query with the input "Orson" and return a row set with "Ralphy"
	mock.
		ExpectQuery(expectedQuery).                               // The query
		WithArgs("Orson").                                        // which arguments are expected
		WillReturnRows(sqlmock.NewRows(columns).AddRow("Ralphy")) // return a rows

	// expect a query with the input "Kurt" and error
	mock.
		ExpectQuery(expectedQuery).         // The query
		WithArgs("Kurt").                   // which arguments are expected
		WillReturnError(errors.New("oops")) // return an error

	// expect a query with the input "Cath" and return a row set with an error
	mock.
		ExpectQuery(expectedQuery). // The query
		WithArgs("Cath").           // which arguments are expected
		WillReturnRows(
			sqlmock.NewRows(columns).AddRow("Cathy").RowError(0, errors.New("oops")),
		) // return a rows

	// expect a query with the input "Boo" and return a row set with an error
	mock.
		ExpectQuery(expectedQuery). // The query
		WithArgs("Becca").          // which arguments are expected
		WillReturnRows(
			sqlmock.NewRows(append(columns, "wrong")).AddRow("Becca", "Boo"), // too many values
		) // return a rows

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greet{
				db: db,
			}

			got := g.Greet(context.Background(), tt.in)
			assert.Equal(t, tt.expect, got)
		})
	}
	// ensure all the mocks we defined have been called
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGreet_CoffeeGreet(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect string
	}{
		{
			name:   "should return the price from the mocked DB for an espresso",
			in:     "espresso",
			expect: "160"},
		{
			name:   "should return the price from the mocked DB for a macchiato",
			in:     "macchiato",
			expect: "180"},
		{
			name:   "should return a free coffee from a string when coffee is not in DB",
			in:     "latte",
			expect: "Free Coffee served from strings"},
	}

	// create a mock database for our unit tests, we are doing this outside of our for loop to
	// share among all our test cases
	// db: is a mocked out DB instance
	// mock: allows us to create expectations and returns
	// err: if anything went wrong
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error()) // stop the tests
	}

	columns := []string{"price"}

	// Expects a query with the input "espresso" and returns the price 160
	mock.ExpectQuery(`SELECT "price" FROM "public"\."coffee" WHERE "type" = \$1 LIMIT 1`).
		WithArgs("espresso").
		WillReturnRows(
			sqlmock.NewRows(columns).AddRow("160"),
		)

	// Expects a query with the input "macchiato" and returns the price 180
	mock.ExpectQuery(`SELECT "price" FROM "public"\."coffee" WHERE "type" = \$1 LIMIT 1`).
		WithArgs("macchiato").
		WillReturnRows(
			sqlmock.NewRows(columns).AddRow("180"),
		)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := &Greet{
				db: db,
			}
			got := g.CoffeeGreet(context.Background(), tc.in)
			assert.Equal(t, tc.expect, got, fmt.Sprintf("%#v", got))
		})
	}
}
