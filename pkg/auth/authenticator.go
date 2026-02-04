package auth

import (
	"net/http"
	"strings"
)

// Authenticator interface defines authentication methods
type Authenticator interface {
	Validate(r *http.Request, w http.ResponseWriter) bool
	Issuer() string
}

const noAuth = "none"

// NewAuthenticator creates new authenticator
func NewAuthenticator(uri, username string) (Authenticator, error) {
	if uri == noAuth || uri == "" {
		return nil, nil
	}

	if strings.HasPrefix(uri, "https://") {
		return NewJWTAuthenticator(uri, username)
	}

	return NewHtpasswdFromFile(uri, username)
}
