package auth

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt/v5"
	jwtRequest "github.com/golang-jwt/jwt/v5/request"
	"github.com/ncarlier/readflow/pkg/oidc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// JWTAuthenticator authenticator use to handle OIDC jWT
type JWTAuthenticator struct {
	issuer   string
	username string
	keystore *oidc.Keystore
	logger   zerolog.Logger
}

// NewJWTAuthenticator create new JWT authenticator
func NewJWTAuthenticator(issuer, username string) (*JWTAuthenticator, error) {
	if _, err := url.ParseRequestURI(issuer); err != nil {
		return nil, fmt.Errorf("invalid issuer URL: %w", err)
	}
	client, err := oidc.NewOIDCClient(issuer, "", "")
	if err != nil {
		return nil, err
	}

	return &JWTAuthenticator{
		issuer:   issuer,
		username: username,
		keystore: client.Keystore,
		logger:   log.With().Str("component", "jwt-autenticator").Logger(),
	}, nil
}

// Validate HTTP request credentials
func (j *JWTAuthenticator) Validate(req *http.Request, res http.ResponseWriter) bool {
	res.Header().Set("WWW-Authenticate", `Bearer realm="Restricted"`)

	token, err := jwtRequest.ParseFromRequest(req, jwtRequest.OAuth2Extractor, func(token *jwt.Token) (i interface{}, e error) {
		if id, ok := token.Header["kid"]; ok {
			return j.keystore.GetKey(id.(string))
		}
		return nil, errors.New("kid header not found in token")
	})

	if err != nil {
		j.logger.Info().Err(err).Msg("unable to extract JWT")
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return j.username == "*" || j.username == getFirstClaimValueFor(claims, "preferred_username", "sub")
	}
	return false
}

// Issuer of the authenticator
func (j *JWTAuthenticator) Issuer() string {
	return j.issuer
}

func getFirstClaimValueFor(claims jwt.MapClaims, names ...string) string {
	for _, name := range names {
		if value, found := claims[name]; found {
			return value.(string)
		}
	}
	return ""
}
