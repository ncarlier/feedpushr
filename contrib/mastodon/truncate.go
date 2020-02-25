package main

func truncate(text string, length int) string {
	r := []rune(text)
	if len(r) > length {
		nbCharToTruncate := len(r) - length
		end := len(text) - nbCharToTruncate
		return string(text[:end])
	}
	return text
}
