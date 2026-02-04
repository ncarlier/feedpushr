package auth

import (
	"net/http"
	"strings"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/rs/zerolog/log"
)

// AuthMiddleware creates an authentication middleware
func AuthMiddleware(authn Authenticator, whitelist ...string) api.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip OPTIONS requests (CORS preflight)
			if r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			// Check if path is whitelisted
			if isWhiteListed(r.URL.Path, whitelist) {
				next.ServeHTTP(w, r)
				return
			}

			// Validate authentication
			if !authn.Validate(r, w) {
				log.Ctx(r.Context()).Info().Msg("authentication failed")
				api.NewResponse(w).Unauthorized("Unauthorized")
				return
			}

			// Proceed
			next.ServeHTTP(w, r)
		})
	}
}

// isWhiteListed checks if a path is in the whitelist
func isWhiteListed(path string, whitelist []string) bool {
	for _, allowed := range whitelist {
		if path == allowed || strings.HasPrefix(path, allowed) {
			return true
		}
	}
	return false
}

// NoAuthMiddleware is a pass-through middleware when no authentication is configured
func NoAuthMiddleware() api.Middleware {
	return func(next http.Handler) http.Handler {
		return next
	}
}

// GetAuthMiddleware returns the appropriate authentication middleware
func GetAuthMiddleware(authn Authenticator, whitelist ...string) api.Middleware {
	if authn == nil {
		return NoAuthMiddleware()
	}
	return AuthMiddleware(authn, whitelist...)
}
