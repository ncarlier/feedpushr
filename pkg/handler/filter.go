package handler

import (
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/filter"
)

// CreateFilterRequest represents a create filter request
type CreateFilterRequest struct {
	Name      string                 `json:"name"`
	Alias     string                 `json:"alias,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty"`
	Condition string                 `json:"condition,omitempty"`
}

// UpdateFilterRequest represents an update filter request
type UpdateFilterRequest struct {
	Alias     *string                `json:"alias,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty"`
	Condition *string                `json:"condition,omitempty"`
	Enabled   bool                   `json:"enabled"`
}

// GetFilterSpecs handles GET /v2/filters/_specs
func (h *Handlers) GetFilterSpecs(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	result := make([]map[string]interface{}, 0, len(h.filterSpecs))
	for _, spec := range h.filterSpecs {
		props := make([]map[string]interface{}, 0, len(spec.PropsSpec))
		for _, prop := range spec.PropsSpec {
			props = append(props, map[string]interface{}{
				"name":    prop.Name,
				"desc":    prop.Desc,
				"type":    prop.Type.String(),
				"options": prop.Options,
			})
		}

		result = append(result, map[string]interface{}{
			"name":  spec.Name,
			"desc":  spec.Desc,
			"props": props,
		})
	}

	res.OK(result)
}

// CreateFilter handles POST /v2/outputs/{id}/filters
func (h *Handlers) CreateFilter(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve output ID from path
	id := req.PathParam("id")

	// Retrieve output processor
	processor, err := h.outputs.GetOutputProcessor(id)
	if err != nil {
		res.NotFound()
		return
	}

	// Decode request payload
	var payload CreateFilterRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Build filter definition
	def := filter.NewBuilder().
		Alias(stringPtr(payload.Alias)).
		Spec(payload.Name).
		Props(payload.Props).
		Condition(stringPtr(payload.Condition)).
		Enable(false).
		NewID().
		Build()

	// Add filter to processor
	f, err := processor.Filters.Add(def)
	if err != nil {
		res.InternalError(err)
		return
	}

	// Save output with new filter
	_, err = h.db.SaveOutput(r.Context(), processor.GetDef())
	if err != nil {
		res.InternalError(err)
		return
	}

	res.Created(f.GetDef())
}

// UpdateFilter handles PUT /v2/outputs/{id}/filters/{fid}
func (h *Handlers) UpdateFilter(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve output ID from path
	id := req.PathParam("id")

	// Retrieve output processor
	processor, err := h.outputs.GetOutputProcessor(id)
	if err != nil {
		res.NotFound()
		return
	}

	// Retrieve filter ID from path
	fid := req.PathParam("fid")

	// Retrieve existing filter
	processorFilter, err := processor.Filters.Get(fid)
	if err != nil {
		res.NotFound()
		return
	}

	// Decode request payload
	var payload UpdateFilterRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Build updated filter definition
	def := filter.NewBuilder().
		From(processorFilter.GetDef()).
		Alias(payload.Alias).
		Props(payload.Props).
		Condition(payload.Condition).
		Enable(payload.Enabled).
		ID(fid).
		Build()

	// Update filter in processor
	processorFilter, err = processor.Filters.Update(fid, def)
	if err != nil {
		res.InternalError(err)
		return
	}

	// Save output with updated filter
	_, err = h.db.SaveOutput(r.Context(), processor.GetDef())
	if err != nil {
		res.InternalError(err)
		return
	}

	res.OK(processorFilter.GetDef())
}

// DeleteFilter handles DELETE /v2/outputs/{id}/filters/{fid}
func (h *Handlers) DeleteFilter(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve output ID from path
	id := req.PathParam("id")

	// Retrieve output processor
	processor, err := h.outputs.GetOutputProcessor(id)
	if err != nil {
		res.NotFound()
		return
	}

	// Retrieve filter ID from path
	fid := req.PathParam("fid")

	// Remove filter from processor
	if err := processor.Filters.Remove(fid); err != nil {
		if err == common.ErrFilterNotFound {
			res.NotFound()
		} else {
			res.InternalError(err)
		}
		return
	}

	// Save output without the deleted filter
	_, err = h.db.SaveOutput(r.Context(), processor.GetDef())
	if err != nil {
		res.InternalError(err)
		return
	}

	res.NoContent()
}
