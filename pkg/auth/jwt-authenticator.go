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

type JWTAuthenticator struct {
	issuer   string
	subject  string
	keystore *oidc.Keystore
	logger   zerolog.Logger
}

func NewJWTAuthenticator(issuer, subject string) (*JWTAuthenticator, error) {
	cfg, err := oidc.GetOIDCConfiguration(issuer)
	if err != nil {
		return nil, err
	}
	return &JWTAuthenticator{
		issuer:   issuer,
		subject:  subject,
		keystore: oidc.NewOIDCKeystore(cfg),
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
		sub, ok := claims["sub"]
		return ok && (j.subject == "*" || sub == j.subject)
	}
	return false
}

func (j *JWTAuthenticator) Issuer() string {
	return j.issuer
}
