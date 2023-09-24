package auth

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	jwtRequest "github.com/dgrijalva/jwt-go/request"
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
	cfg, err := oidc.GetOIDCConfiguration(issuer)
	if err != nil {
		return nil, err
	}
	keystore, err := oidc.NewOIDCKeystore(cfg)
	if err != nil {
		return nil, err
	}
	go keystore.Start()
	return &JWTAuthenticator{
		issuer:   issuer,
		username: username,
		keystore: keystore,
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
