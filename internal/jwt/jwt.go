package jwt

import (
	"errors"
	"log"
	"time"

	_jwt "github.com/golang-jwt/jwt/v4"
)

// CTXUser to namespace our user on the context
const CTXUser ctxUserKey = `user`

// ErrInvalidToken for authorisation
var ErrInvalidToken = errors.New("invalid token provided")

type (
	// ctxUserKey context namespace our user value
	ctxUserKey string
	// CustomClaims our jwt token as Go
	CustomClaims struct {
		Username string   `json:"username"`
		Roles    []string `json:"roles"`
		DB       bool     `json:"db"`
		_jwt.RegisteredClaims
	}
)

// New token dataset
func New() CustomClaims {
	return CustomClaims{
		Username: "gus",
		Roles:    []string{"operator", "admin", "barista"},
		DB:       true,
		RegisteredClaims: _jwt.RegisteredClaims{
			Issuer:    "test",
			Subject:   "somebody",
			Audience:  []string{"somebody_else"},
			ExpiresAt: _jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			NotBefore: _jwt.NewNumericDate(time.Now()),
			IssuedAt:  _jwt.NewNumericDate(time.Now()),
			ID:        "1",
		},
	}
}

// GetClaims from a JWT token
func GetClaims(tokenString string) (*CustomClaims, error) {
	c := &CustomClaims{}
	token, err := _jwt.ParseWithClaims(tokenString, c, func(token *_jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if err != nil {
		log.Println("oops! error parsing token", err.Error())
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		log.Println("oops! no valid token present after parsing")
		return nil, ErrInvalidToken
	}
	return claims, nil
}
