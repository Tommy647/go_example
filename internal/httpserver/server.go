package httpserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Tommy647/go_example/internal/dbgreeter"

	"github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/jwt"
)

// HelloResponse to http requests // @todo: fixme
type HelloResponse struct{}

// HandleHello as a http request
func HandleHello() http.Handler {
	g := greeter.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello http request")
		c := jwt.GetUser(r.Context())
		if c == nil {
			_, _ = w.Write([]byte(g.Greet(r.Context(), "")))
			return
		}
		_, _ = w.Write([]byte(g.Greet(r.Context(), c.Username)))
		_, _ = w.Write([]byte(g.Greet(r.Context(), c.Subject)))
		for i := range c.Roles {
			_, _ = w.Write([]byte(g.Greet(r.Context(), c.Roles[i])))
		}
	})
}

type CoffeeResponse struct{}

func HandleCoffee(dbConn *sql.DB) http.Handler {
	g := dbgreeter.New(dbConn)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("coffee http request")
		u := jwt.GetUser(r.Context())
		switch u {
		case nil:
			_, _ = w.Write([]byte(g.CoffeeGreet(r.Context(), "no CustomClaim available")))
		default:
			if u.Username == "gus" {
				// espresso is ignored at the moment @todo: Fix it to accept a DBConn
				_, _ = w.Write([]byte(g.CoffeeGreet(r.Context(), "espresso")))

			}
			// macchiato is ignored at the moment @todo: fix it to accept a DBConn
			_, _ = w.Write([]byte(g.CoffeeGreet(r.Context(), "macchiato")))
		}
	})
}
