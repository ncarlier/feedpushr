package controller

import (
	"net/http"
	"net/url"

	"github.com/goadesign/goa"
)

// Redirect to other path
func Redirect(to string) goa.MuxHandler {
	return func(rw http.ResponseWriter, req *http.Request, v url.Values) {
		http.Redirect(rw, req, to, http.StatusMovedPermanently)
	}
}
