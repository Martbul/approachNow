package middleware

import (
	"net/http"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/auth/utils"
)

// JWTAuthMiddleware checks for a valid JWT in the Authorization header
func JWTAuthMiddleware(log hclog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Token format: Bearer <token>
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Log the user's email (optional)
		log.Info("Authenticated user", "email", claims.Email)

		// Pass the request to the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
