package fn

// Truncate shorten text content
func Truncate(length int, text string) string {
	r := []rune(text)
	if len(r) > length {
		return string(text[:length])
	}
	return text
}
