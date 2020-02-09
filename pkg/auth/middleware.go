package auth

import (
	"net/http"
	"strings"

	"context"

	"github.com/goadesign/goa"
)

func isWhiteListed(path string, whitelist []string) bool {
	for _, prefix := range whitelist {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

// NewMiddleware creates a static auth middleware.
func NewMiddleware(authn Authenticator, whitelist ...string) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			if !isWhiteListed(req.URL.Path, whitelist) && !authn.Validate(req) {
				goa.LogInfo(ctx, "failed auth")
				rw.Header().Set("WWW-Authenticate", `Basic realm="Ah ah ah, you didn't say the magic word"`)
				return goa.ErrUnauthorized("invalid credentials")
			}
			// Proceed
			return h(ctx, rw, req)
		}
	}
}
