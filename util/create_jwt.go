package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)
/*
type Claim struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}*/

type Claims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJWT(secret string, data Claims) (string, error) {
	var usr Claims

	usr.UserID = data.UserID
	usr.Email = data.Email
	usr.Role = data.Role
	usr.IssuedAt = jwt.NewNumericDate(time.Now())
	usr.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, usr)

	return token.SignedString([]byte(secret))
}