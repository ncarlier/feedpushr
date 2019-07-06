package model

const mask = "########"

// MaskSecret mask secret string
func MaskSecret(secret string) string {
	l := len(secret)
	if l > 8 {
		return string(secret[:3] + mask + secret[l-3:])
	}
	return mask
}
