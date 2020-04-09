package helper

import (
	"strings"
)

// IsEmptyString test if a string pointer is nil or empty
func IsEmptyString(s *string) bool {
	return s == nil || len(strings.TrimSpace(*s)) == 0
}
