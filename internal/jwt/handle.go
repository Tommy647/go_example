package jwt

import (
	"net/http"

	_jwt "github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"github.com/Tommy647/go_example/internal/logger"
)

// HandleNewToken http response
func HandleNewToken(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.Context(), "new token")
	// get a new token, this is just placeholder code, user data would be filled in from the SSO service
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, New())
	// sign the token
	ss, err := token.SignedString(getSecret())
	if err != nil {
		logger.Error(r.Context(), "signing key", zap.Error(err))
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	// return the response
	_, _ = w.Write([]byte(ss))
}
