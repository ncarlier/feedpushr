package controller

import (
	"io/fs"
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/assets"
)

// NewUIHandler creates a handler to serve the UI assets
func NewUIHandler() http.Handler {
	fsys, err := fs.Sub(assets.Content, "content")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(fsys))
}

// NewRedirectHandler creates a redirect handler
func NewRedirectHandler(to string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, http.StatusMovedPermanently)
	}
}
