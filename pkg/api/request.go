package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	defaultPageSize = 10
	maxPageSize     = 100
)

// Request represents a helper for parsing HTTP requests
type Request struct {
	r *http.Request
}

// NewRequest creates a new request helper
func NewRequest(r *http.Request) *Request {
	return &Request{r: r}
}

// PathParam extracts a path parameter by name from the URL
// This assumes the router stores path params in the request context
func (r *Request) PathParam(name string) string {
	return r.r.PathValue(name)
}

// QueryParam extracts a query parameter by name
func (r *Request) QueryParam(name string) string {
	return r.r.URL.Query().Get(name)
}

// QueryParams extracts multiple query parameters by name
func (r *Request) QueryParams(name string) []string {
	return r.r.URL.Query()[name]
}

// QueryParamInt extracts an integer query parameter
func (r *Request) QueryParamInt(name string, defaultValue int) int {
	value := r.QueryParam(name)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// QueryParamBool extracts a boolean query parameter
func (r *Request) QueryParamBool(name string, defaultValue bool) bool {
	value := r.QueryParam(name)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

// Page extracts pagination parameters from query params
func (r *Request) Page() (page, size int) {
	page = r.QueryParamInt("page", 0)
	size = r.QueryParamInt("size", defaultPageSize)

	if page < 0 {
		page = 0
	}
	if size < 1 {
		size = defaultPageSize
	}
	if size > maxPageSize {
		size = maxPageSize
	}

	return page, size
}

// Decode decodes the request body into the given value
func (r *Request) Decode(v interface{}) error {
	defer r.r.Body.Close()
	return json.NewDecoder(r.r.Body).Decode(v)
}

// Body returns the request body as bytes
func (r *Request) Body() ([]byte, error) {
	defer r.r.Body.Close()
	return io.ReadAll(r.r.Body)
}

// LogRequest logs the incoming HTTP request
func LogRequest(r *http.Request) {
	log.Debug().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("remote", r.RemoteAddr).
		Msg("incoming request")
}
