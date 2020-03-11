package fn

import "fmt"

// Tweet truncate text if needed and put the link after
func Tweet(text string, suffix string) string {
	text = Truncate(270-len(suffix), text)
	return fmt.Sprintf("%s\n%s", text, suffix)
}
