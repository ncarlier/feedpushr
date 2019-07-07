package common

import "errors"

var (
	// ErrFeedNotFound is returned when a feed is not found in the DB.
	ErrFeedNotFound = errors.New("feed not found")
	// ErrFilterNotFound is returned when a filter is not found in the DB.
	ErrFilterNotFound = errors.New("filter not found")
	// ErrOutputNotFound is returned when a output is not found in the DB.
	ErrOutputNotFound = errors.New("output not found")
)
