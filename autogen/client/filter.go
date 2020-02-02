// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": filter Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v2/design
// --out=/home/nicolas/workspace/feedpushr/autogen
// --version=v1.4.3

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateFilterPayload is the filter create action payload.
type CreateFilterPayload struct {
	// Alias of the filter
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Conditional expression of the output
	Condition string `form:"condition" json:"condition" yaml:"condition" xml:"condition"`
	// Name of the filter
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Filter properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
}

// CreateFilterPath computes a request path to the create action of filter.
func CreateFilterPath() string {

	return fmt.Sprintf("/v1/filters")
}

// Create a new filter
func (c *Client) CreateFilter(ctx context.Context, path string, payload *CreateFilterPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateFilterRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateFilterRequest create the request corresponding to the create action endpoint of the filter resource.
func (c *Client) NewCreateFilterRequest(ctx context.Context, path string, payload *CreateFilterPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteFilterPath computes a request path to the delete action of filter.
func DeleteFilterPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/filters/%s", param0)
}

// Delete a filter
func (c *Client) DeleteFilter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteFilterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteFilterRequest create the request corresponding to the delete action endpoint of the filter resource.
func (c *Client) NewDeleteFilterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetFilterPath computes a request path to the get action of filter.
func GetFilterPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/filters/%s", param0)
}

// Retrieve filter with given ID
func (c *Client) GetFilter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetFilterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetFilterRequest create the request corresponding to the get action endpoint of the filter resource.
func (c *Client) NewGetFilterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListFilterPath computes a request path to the list action of filter.
func ListFilterPath() string {

	return fmt.Sprintf("/v1/filters")
}

// Retrieve all filters definitions
func (c *Client) ListFilter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListFilterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFilterRequest create the request corresponding to the list action endpoint of the filter resource.
func (c *Client) NewListFilterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SpecsFilterPath computes a request path to the specs action of filter.
func SpecsFilterPath() string {

	return fmt.Sprintf("/v1/filters/_specs")
}

// Retrieve all filter types available
func (c *Client) SpecsFilter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSpecsFilterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSpecsFilterRequest create the request corresponding to the specs action endpoint of the filter resource.
func (c *Client) NewSpecsFilterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateFilterPayload is the filter update action payload.
type UpdateFilterPayload struct {
	// Alias of the filter
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Conditional expression of the output
	Condition *string `form:"condition,omitempty" json:"condition,omitempty" yaml:"condition,omitempty" xml:"condition,omitempty"`
	// Filter status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// Filter properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
}

// UpdateFilterPath computes a request path to the update action of filter.
func UpdateFilterPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/v1/filters/%s", param0)
}

// Update a filter
func (c *Client) UpdateFilter(ctx context.Context, path string, payload *UpdateFilterPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateFilterRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateFilterRequest create the request corresponding to the update action endpoint of the filter resource.
func (c *Client) NewUpdateFilterRequest(ctx context.Context, path string, payload *UpdateFilterPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
