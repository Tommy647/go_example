package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Tommy647/go_example/internal/jwt"
)

// ErrInvalidToken has been provided in the headers
var ErrInvalidToken = errors.New("invalid token")

// WithAuth checks the JWT as middleware
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello, I am middleware, checking for a JWT token")
		tokenString, err := getAuthToken(r)
		if err != nil {
			log.Println("getting token", err.Error())
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(err.Error()))
		}

		claims, err := jwt.GetClaims(tokenString)
		if err != nil {
			log.Println("oops! no valid token present after parsing")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("valid auth token required"))
			return
		}

		ctx := jwt.WithUser(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Println("hello again, middleware here, checked for a JWT token finished calling my child request")
	})
}

func getAuthToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		log.Println("oops! no token present")
		return "", ErrInvalidToken
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")
	if tokenString == auth {
		log.Println("oops! no valid token present")
		return "", ErrInvalidToken
	}
	return tokenString, nil
}
