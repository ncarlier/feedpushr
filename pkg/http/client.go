package http

import (
	"net"
	"net/http"
	"time"
)

// DefaultTimeout is the default HTTP request timeout
const DefaultTimeout = time.Duration(5 * time.Second)

// DefaultTransport is the default HTTP client transport
var DefaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   5 * time.Second,  // Default is 30s.
		KeepAlive: 15 * time.Second, // Default is 30s.
	}).DialContext,
	// ForceAttemptHTTP2: true, // Fail with Reddit?!?
	MaxIdleConns:          50,               // Default is 100
	IdleConnTimeout:       20 * time.Second, // Default is 90s.
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// DefaultClient is the default HTTP client with timeout
var DefaultClient = &http.Client{
	Timeout:   DefaultTimeout,
	Transport: DefaultTransport,
}

// New HTTP client
func New(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:   timeout,
		Transport: DefaultTransport,
	}
}
