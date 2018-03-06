package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// HTTPOutputProvider HTTP output provider
type HTTPOutputProvider struct {
	uri string
}

func newHTTPOutputProvider(uri string) *HTTPOutputProvider {
	return &HTTPOutputProvider{
		uri: uri,
	}
}

// Send article to HTTP endpoint.
func (op *HTTPOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	resp, err := http.Post(op.uri, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	return nil
}
