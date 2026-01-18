package middleware

import (
	"context"
	"net/http"
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/pkg/utils"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey string

const UserContextKey contextKey = "user"

// SessionClaims represents the session data stored in context
type SessionClaims struct {
	SessionID uuid.UUID
	UserID    int
}

func AuthMiddleware(logger *zap.Logger, userRepo repository.UsersRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			
			// Get Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("missing authorization header")
				utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "missing authorization token")
				return
			}

			// Check if it starts with Bearer
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				logger.Warn("invalid authorization header format")
				utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "invalid authorization format")
				return
			}

			tokenString := parts[1]

			// Parse UUID token
			sessionID, err := uuid.Parse(tokenString)
			if err != nil {
				logger.Warn("invalid token format", zap.Error(err))
				utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "invalid token format")
				return
			}

			// Validate session in database
			session, err := userRepo.ValidateSession(ctx, sessionID)
			if err != nil {
				logger.Warn("invalid or expired session", zap.String("session_id", sessionID.String()), zap.Error(err))
				utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "invalid or expired session")
				return
			}

			// Create session claims
			claims := &SessionClaims{
				SessionID: session.ID,
				UserID:    session.UserID,
			}

			// Add claims to context
			ctx = context.WithValue(ctx, UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Helper function to get session claims from context
func GetUserFromContext(ctx context.Context) (*SessionClaims, bool) {
	claims, ok := ctx.Value(UserContextKey).(*SessionClaims)
	return claims, ok
}
