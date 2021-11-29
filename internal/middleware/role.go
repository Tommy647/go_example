package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tommy647/go_example/internal/jwt"
)

// WithRole to check whether the user has the "barista" role
func WithRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking for the user role")
		u := jwt.GetUser(r.Context())
		if findRole(u.Roles, "barista") {
			next.ServeHTTP(w, r)
		}
	})
}

// findRole will return true if the provided []string contains a string == to the second argument
func findRole(in []string, v string) bool {
	for i := range in {
		if strings.EqualFold(in[i], v) {
			return true
		}
	}
	return false
}
