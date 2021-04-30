package common

import "errors"

var (
	// ErrFeedAlreadyExists is returned when a feed is already exists in the DB.
	ErrFeedAlreadyExists = errors.New("feed already exists")
	// ErrFeedNotFound is returned when a feed is not found in the DB.
	ErrFeedNotFound = errors.New("feed not found")
	// ErrFilterNotFound is returned when a filter is not found in the DB.
	ErrFilterNotFound = errors.New("filter not found")
	// ErrOutputNotFound is returned when a output is not found in the DB.
	ErrOutputNotFound = errors.New("output not found")
	// ErrFeedQuotaExceeded is returned when feed quota is exceeded.
	ErrFeedQuotaExceeded = errors.New("feed quota exceeded")
	// ErrOutputQuotaExceeded is returned when output quota is exceeded.
	ErrOutputQuotaExceeded = errors.New("output quota exceeded")
)
