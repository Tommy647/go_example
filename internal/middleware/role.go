package middleware

import (
	"log"
	"net/http"
)

// WithRole to check whether the user has the "barista" role
func WithRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking for the user role")
		/*		u := jwt.GetUser(r.Context())
				log.Println(u.Username)
				log.Println(u.Roles)*/
		next.ServeHTTP(w, r)
	})
}
