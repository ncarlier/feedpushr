package common

import (
	"io"
	"mime"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

// GetNormalizedBody get body reader using UTF-8
func GetNormalizedBody(res *http.Response) (io.Reader, error) {
	contentType := res.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	if err == nil {
		if enc, found := params["charset"]; found {
			enc = strings.ToLower(enc)
			if enc != "utf-8" && enc != "utf8" && enc != "" {
				return charset.NewReader(res.Body, contentType)
			}
		}
	}
	return res.Body, nil
}
