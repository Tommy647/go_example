package dbgreeter

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGreeter_HelloGreet(t *testing.T) {
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
			g := &Greeter{
				db: db,
			}

			got := g.HelloGreet(context.Background(), tt.in)
			assert.Equal(t, tt.expect, got)
		})
	}
	// ensure all the mocks we defined have been called
	assert.NoError(t, mock.ExpectationsWereMet())
}
