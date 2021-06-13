package auth

import (
	"context"
	"net/http"
	"net/url"

	"github.com/goadesign/goa"
)

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
func NewAuthenticator(uri, subject string) (Authenticator, error) {
	if uri == "none" {
		return nil, nil
	}
	_, err := url.ParseRequestURI(uri)
	if err == nil {
		return NewJWTAuthenticator(uri, subject)
	}
	return NewHtpasswdFromFile(uri, subject)
}
