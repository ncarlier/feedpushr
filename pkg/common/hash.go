package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Hash create a hash of a list of values
func Hash(values ...string) string {
	value := strings.Join(values, "|")
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}
