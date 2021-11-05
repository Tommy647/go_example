package jwt

import (
	"log"
	"net/http"

	_jwt "github.com/golang-jwt/jwt/v4"
)

// HandleNewToken http response
func HandleNewToken(w http.ResponseWriter, r *http.Request) {
	// get a new token, this is just placeholder code, user data would be filled in from the SSO service
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, New())
	// sign the token
	ss, err := token.SignedString(getSecret())
	if err != nil {
		log.Println("error signing key:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	// return the response
	_, _ = w.Write([]byte(ss))
}
