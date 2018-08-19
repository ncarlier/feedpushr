package assets

import (
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/goadesign/goa"
	"github.com/rakyll/statik/fs"
)

var instance http.FileSystem
var once sync.Once

// GetFS return assets file system instance
func GetFS() http.FileSystem {
	once.Do(func() {
		var err error
		instance, err = fs.New()
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}

// Handler to fetch assets from the virtual file system
func Handler() goa.MuxHandler {
	h := http.FileServer(GetFS())
	return func(rw http.ResponseWriter, req *http.Request, v url.Values) {
		h.ServeHTTP(rw, req)
	}
}
