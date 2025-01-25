package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/goadesign/goa"
)

const noAuth = "none"

func isWhiteListed(path string, whitelist []string) bool {
	for _, allowed := range whitelist {
		if path == allowed {
			return true
		}
	}
	return false
}

// NewMiddleware creates a static auth middleware.
func NewMiddleware(authn Authenticator, whitelist ...string) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
			if req.Method != "OPTIONS" && !isWhiteListed(req.URL.Path, whitelist) && !authn.Validate(req, res) {
				goa.LogInfo(ctx, "authentication failed")
				return goa.ErrUnauthorized("Unauthorized")
			}
			// Proceed
			return h(ctx, res, req)
		}
	}
}

// NewAuthenticator create new authenticator
func NewAuthenticator(uri, username string) (Authenticator, error) {
	if uri == noAuth {
		return nil, nil
	}

	if strings.HasPrefix(uri, "https://") {
		return NewJWTAuthenticator(uri, username)
	}

	return NewHtpasswdFromFile(uri, username)
}
