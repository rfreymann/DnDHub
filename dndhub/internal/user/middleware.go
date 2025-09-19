package user

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "changeme"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid claims", http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "invalid subject claim", http.StatusUnauthorized)
			return
		}

		// inject user_id into context
		ctx := context.WithValue(r.Context(), userIDKey, sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID extracts user_id from context (helper)
func GetUserID(r *http.Request) (string, bool) {
	uid, ok := r.Context().Value(userIDKey).(string)
	return uid, ok
}
