package handler

import (
	"net/http"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/output"
)

// CreateOutputRequest represents a create output request
type CreateOutputRequest struct {
	Name      string                 `json:"name"`
	Alias     string                 `json:"alias,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty"`
	Condition string                 `json:"condition,omitempty"`
}

// UpdateOutputRequest represents an update output request
type UpdateOutputRequest struct {
	Alias     string                 `json:"alias,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty"`
	Condition string                 `json:"condition,omitempty"`
	Enabled   bool                   `json:"enabled"`
}

// CreateOutput handles POST /v2/outputs
func (h *Handlers) CreateOutput(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Decode request payload
	var payload CreateOutputRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Create output definition
	def := output.NewBuilder().
		Alias(stringPtr(payload.Alias)).
		Spec(payload.Name).
		Props(payload.Props).
		Condition(stringPtr(payload.Condition)).
		Enable(false).
		NewID().
		Build()

	// Register output processor
	processor, err := h.outputs.AddOutputProcessor(def)
	if err != nil {
		res.InternalError(err)
		return
	}

	// Persist output definition
	def, err = h.db.SaveOutput(r.Context(), processor.GetDef())
	if err != nil {
		// cleanup previous created processor
		_def := processor.GetDef()
		h.outputs.RemoveOutputProcessor(&_def)
		res.InternalError(err)
		return
	}

	res.Created(def)
}

// UpdateOutput handles PUT /v2/outputs/{id}
func (h *Handlers) UpdateOutput(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve output ID from path
	id := req.PathParam("id")

	// Decode request payload
	var payload UpdateOutputRequest
	if err := req.Decode(&payload); err != nil {
		res.BadRequest(err)
		return
	}

	// Retrieve output processor
	processor, err := h.outputs.GetOutputProcessor(id)
	if err != nil {
		if err == common.ErrOutputNotFound {
			res.NotFound()
			return
		}
		res.InternalError(err)
		return
	}

	// Build updated output definition
	update := output.NewBuilder().
		From(processor.GetDef()).
		Alias(stringPtr(payload.Alias)).
		Props(payload.Props).
		Condition(stringPtr(payload.Condition)).
		Enable(payload.Enabled).
		Build()

	// Update output processor
	processor, err = h.outputs.UpdateOutputProcessor(update)
	if err != nil {
		res.InternalError(err)
		return
	}

	// Save output with updated definition
	def, err := h.db.SaveOutput(r.Context(), processor.GetDef())
	if err != nil {
		res.InternalError(err)
		return
	}

	res.OK(def)
}

// DeleteOutput handles DELETE /v2/outputs/{id}
func (h *Handlers) DeleteOutput(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	// Retrieve output ID from path
	id := req.PathParam("id")

	// Remove output processor
	def := &model.OutputDef{ID: id}
	if err := h.outputs.RemoveOutputProcessor(def); err != nil {
		res.NotFound()
		return
	}

	// Delete output from database
	if _, err := h.db.DeleteOutput(r.Context(), def.ID); err != nil {
		res.InternalError(err)
		return
	}

	res.NoContent()
}

// GetOutput handles GET /v2/outputs/{id}
func (h *Handlers) GetOutput(w http.ResponseWriter, r *http.Request) {
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

	res.OK(processor.GetDef())
}

// ListOutputs handles GET /v2/outputs
func (h *Handlers) ListOutputs(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	// Retrieve output definitions
	outputs := h.outputs.GetOutputDefs()
	result := make([]interface{}, 0, len(outputs))
	for _, def := range outputs {
		result = append(result, def)
	}

	res.OK(result)
}

// GetOutputSpecs handles GET /v2/outputs/_specs
func (h *Handlers) GetOutputSpecs(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	// Retrieve output specifications
	specs := h.outputs.GetAvailableOutputs()
	result := make([]map[string]interface{}, 0, len(specs))

	for _, spec := range specs {
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

// stringPtr returns a pointer to a string if not empty, otherwise nil
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
