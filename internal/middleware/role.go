package middleware

import (
	"log"
	"net/http"

	"github.com/Tommy647/go_example/internal/jwt"
)

// WithRole to check whether the user has the "barista" role
func WithRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking for the user role")
		u := jwt.GetUser(r.Context())

		if u.Roles[2] != "barista" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("appropriate role required"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*
// containRole will return true if the provided []string contains a string == to the second argument
func containRole(in []string, v string) bool {
	for _, n := range in {

		if !strings.EqualFold(n, v) {
			return true
		}
	}
	return false
}*/
