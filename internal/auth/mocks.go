package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func NewMockAuthenticator() Authenticator {
	return &TestAuthenticator{}
}

type TestAuthenticator struct{}

const secret = "test"

func (a *TestAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString, nil
}

func (a *TestAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
