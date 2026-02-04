package api

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Response represents a standard API response
type Response struct {
	writer http.ResponseWriter
}

// NewResponse creates a new response helper
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{writer: w}
}

// JSON writes a JSON response with the given status code
func (r *Response) JSON(status int, data interface{}) {
	r.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.writer.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(r.writer).Encode(data); err != nil {
			log.Error().Err(err).Msg("failed to encode JSON response")
		}
	}
}

// OK writes a 200 OK JSON response
func (r *Response) OK(data interface{}) {
	r.JSON(http.StatusOK, data)
}

// Created writes a 201 Created JSON response
func (r *Response) Created(data interface{}) {
	r.JSON(http.StatusCreated, data)
}

// Accepted writes a 202 Accepted response
func (r *Response) Accepted(data ...interface{}) {
	if len(data) > 0 {
		r.JSON(http.StatusAccepted, data[0])
	} else {
		r.writer.WriteHeader(http.StatusAccepted)
	}
}

// NoContent writes a 204 No Content response
func (r *Response) NoContent() {
	r.writer.WriteHeader(http.StatusNoContent)
}

// BadRequest writes a 400 Bad Request JSON response
func (r *Response) BadRequest(err error) {
	r.Error(http.StatusBadRequest, err.Error())
}

// Unauthorized writes a 401 Unauthorized JSON response
func (r *Response) Unauthorized(message string) {
	r.Error(http.StatusUnauthorized, message)
}

// NotFound writes a 404 Not Found JSON response
func (r *Response) NotFound() {
	r.Error(http.StatusNotFound, "resource not found")
}

// InternalError writes a 500 Internal Server Error JSON response
func (r *Response) InternalError(err error) {
	log.Error().Err(err).Msg("internal server error")
	r.Error(http.StatusInternalServerError, err.Error())
}

// Error writes an error JSON response with the given status code
func (r *Response) Error(status int, message string) {
	r.JSON(status, map[string]string{"error": message})
}

// Raw writes raw bytes with the given content type
func (r *Response) Raw(contentType string, data []byte) {
	r.writer.Header().Set("Content-Type", contentType)
	r.writer.WriteHeader(http.StatusOK)
	if _, err := r.writer.Write(data); err != nil {
		log.Error().Err(err).Msg("failed to write raw response")
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
