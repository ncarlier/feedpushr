package common

import (
	"io"
	"mime"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

// IsEmptyString test if a string pointer is nil or empty
func IsEmptyString(s *string) bool {
	return s == nil || len(strings.TrimSpace(*s)) == 0
}

// GetNormalizedBodyFromResponse get body reader from HTTP response using UTF-8
func GetNormalizedBodyFromResponse(res *http.Response) (io.Reader, error) {
	contentType := res.Header.Get("Content-Type")
	return getNormalizedBody(contentType, res.Body)
}

// GetNormalizedBodyFromRequest get body reader from HTTP request using UTF-8
func GetNormalizedBodyFromRequest(req *http.Request) (io.Reader, error) {
	contentType := req.Header.Get("Content-Type")
	return getNormalizedBody(contentType, req.Body)
}

func getNormalizedBody(contentType string, body io.ReadCloser) (io.Reader, error) {
	_, params, err := mime.ParseMediaType(contentType)
	if err == nil {
		if enc, found := params["charset"]; found {
			enc = strings.ToLower(enc)
			if enc != "utf-8" && enc != "utf8" && enc != "" {
				return charset.NewReader(body, contentType)
			}
		}
	}
	return body, nil
}
