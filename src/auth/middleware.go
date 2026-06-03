package auth

import (
	"context"
	"eco-map-app/additional"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			additional.PrintError("/api/auth", fmt.Errorf("No auth hearder in middleware"))
			http.Error(w, `{"error": "authorization header required"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if expiresAt, exists := TokenBlacklist[tokenString]; exists {
			if time.Now().Before(expiresAt) {
				additional.PrintError("/api/auth", fmt.Errorf("Token revoked in middleware"))
				http.Error(w, `{"error": "token has been revoked"}`, http.StatusUnauthorized)
				return
			}
			delete(TokenBlacklist, tokenString)
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			additional.PrintError("/api/auth", fmt.Errorf("Invalid token in middleware"))
			http.Error(w, `{"error": "invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
