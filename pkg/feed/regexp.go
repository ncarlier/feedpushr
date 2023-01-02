package feed

import "regexp"

// ValidFeedContentType is a REGEXP used to validate valid feed Content-Type
var ValidFeedContentType = regexp.MustCompile(`^(application|text)/(\w+\+)?xml`)
