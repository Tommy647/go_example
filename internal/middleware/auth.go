package middleware

import (
	"errors"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/Tommy647/go_example/internal/jwt"
	"github.com/Tommy647/go_example/internal/logger"
)

// ErrInvalidToken has been provided in the headers
var ErrInvalidToken = errors.New("invalid token")

// WithAuth checks the JWT as middleware
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Context(), "hello, I am middleware, checking for a JWT token")
		tokenString, err := getAuthToken(r)
		if err != nil {
			logger.Error(r.Context(), "getting jwt token", zap.Error(err))
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		claims, err := jwt.GetClaims(tokenString)
		if err != nil {
			logger.Error(r.Context(), "oops! no valid token present after parsing")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("valid auth token required"))
			return
		}

		ctx := jwt.WithUser(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		logger.Info(r.Context(), "hello again, middleware here, checked for a JWT token finished calling my child request")
	})
}

func getAuthToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		logger.Warn(r.Context(), "oops! no token present")
		return "", ErrInvalidToken
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")
	if tokenString == auth {
		logger.Warn(r.Context(), "oops! no valid token present")
		return "", ErrInvalidToken
	}
	return tokenString, nil
}
