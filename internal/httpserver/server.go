package httpserver

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

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
	log.Println("handleCoffee: opening a connection to the DB")
	gDB := dbgreeter.New(dbConn)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("coffee http request")
		u := jwt.GetUser(r.Context())
		if !(findRole(u, "barista") && findRole(u, "db")) {
			g := greeter.New()
			_, _ = w.Write([]byte(g.CoffeeGreet(r.Context(), "")))
			return
		}
		_, _ = w.Write([]byte(gDB.CoffeeGreet(r.Context(), "Espresso")))
	})
}

func findRole(u *jwt.CustomClaims, role string) bool {
	for _, v := range u.Roles {
		if strings.EqualFold(v, role) {
			return true
		}
	}
	return false
}
