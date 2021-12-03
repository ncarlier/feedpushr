package controller

import (
	"io/fs"
	"net/http"
	"net/url"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/pkg/assets"
)

// UIHandler to fetch assets from the virtual file system
func UIHandler() goa.MuxHandler {
	fsys, err := fs.Sub(assets.Content, "content")
	if err != nil {
		panic(err)
	}
	h := http.FileServer(http.FS(fsys))
	return func(rw http.ResponseWriter, req *http.Request, v url.Values) {
		h.ServeHTTP(rw, req)
	}
}
