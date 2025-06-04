package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func newJWTToken(id string, name string, isAdmin bool) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &JWTClaim{
		id,
		name,
		isAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
