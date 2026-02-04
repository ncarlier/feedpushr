package handler

import (
	"errors"
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
)

// ExploreResponse represents an explore result
type ExploreResponse struct {
	Title   string  `json:"title"`
	Desc    string  `json:"desc"`
	XMLURL  string  `json:"xmlurl"`
	HTMLURL *string `json:"htmlurl,omitempty"`
}

// ExploreFeeds handles GET /v2/explore
func (h *Handlers) ExploreFeeds(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	q := req.QueryParam("q")
	if q == "" {
		res.BadRequest(errors.New("missing query parameter"))
		return
	}

	results, err := h.explorer.Search(q)
	if err != nil {
		res.InternalError(err)
		return
	}

	response := make([]ExploreResponse, 0, len(*results))
	for _, result := range *results {
		htmlURL := &result.HTMLURL
		response = append(response, ExploreResponse{
			Title:   result.Title,
			Desc:    result.Desc,
			XMLURL:  result.XMLURL,
			HTMLURL: htmlURL,
		})
	}

	res.OK(response)
}
