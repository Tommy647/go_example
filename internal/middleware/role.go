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
		if findRole(u, "admin") && findRole(u, "db") {
			next.ServeHTTP(w, r)
			return
		}
		log.Printf("no authorized to get a coffee: %s", u.Username)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("valid role required"))
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
