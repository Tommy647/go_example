package httpserver

import (
	"log"
	"net/http"

	"github.com/Tommy647/go_example/internal/customgreeter"

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
		_, _ = w.Write([]byte(g.Greet(r.Context(), c.Foo)))
		_, _ = w.Write([]byte(g.Greet(r.Context(), c.Subject)))
		for i := range c.Roles {
			_, _ = w.Write([]byte(g.Greet(r.Context(), c.Roles[i])))
		}
	})
}

// HandleGreet as a http request
func HandleGreet() http.Handler {
	g := customgreeter.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("greeting http request")
		c := jwt.GetUser(r.Context())
		greetings, ok := r.URL.Query()["greeting"]
		if !ok || len(greetings[0]) < 1 {
			log.Println("Url Param 'greeting' is missing")
			return
		}
		greeting := greetings[0]
		if c == nil {
			_, _ = w.Write([]byte(g.Greet(r.Context(), greeting, "")))
			return
		}
		_, _ = w.Write([]byte(g.Greet(r.Context(), greeting, c.Foo)))
		_, _ = w.Write([]byte(g.Greet(r.Context(), greeting, c.Subject)))
		for i := range c.Roles {
			_, _ = w.Write([]byte(g.Greet(r.Context(), greeting, c.Roles[i])))
		}
	})
}
