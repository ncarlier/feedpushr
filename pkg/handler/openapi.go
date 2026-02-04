package handler

import (
	_ "embed"
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
)

//go:embed openapi.json
var openapi []byte

// GetOpenAPI handles GET /v2/openapi.json
func (h *Handlers) GetOpenAPI(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)
	res.Raw("application/json", openapi)
}
