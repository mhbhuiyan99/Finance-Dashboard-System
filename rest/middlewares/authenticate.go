package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ClaimsKey contextKey = "claims"

type Claims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		accessToken := headerArr[1]

		token, err := jwt.ParseWithClaims(
			accessToken,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(m.cnf.JwtSecretKey), nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}






