package httpserver

import (
	"log"
	"net/http"

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
