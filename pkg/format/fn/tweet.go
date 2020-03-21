package fn

import "fmt"

// Tweet truncate text if needed and put the link after
func Tweet(text string, suffix string) string {
	toTruncate := 270 - len(suffix)
	if toTruncate > 0 {
		text = Truncate(toTruncate, text)
	} else {
		// suffix too long, text is sacrificed
		return suffix
	}
	return fmt.Sprintf("%s\n%s", text, suffix)
}
