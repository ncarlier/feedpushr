package config

import "strings"

// ArrayFlags contains an array of command flags
type ArrayFlags struct {
	items    []string
	fallback []string
}

// Values return the values of a flag array
func (i *ArrayFlags) Values() []string {
	if len(i.items) == 0 {
		return i.fallback
	}
	return i.items
}

// String return the string value of a flag array
func (i *ArrayFlags) String() string {
	return strings.Join(i.Values(), ",")
}

// Set is used to add a value to the flag array
func (i *ArrayFlags) Set(value string) error {
	i.items = append(i.items, value)
	return nil
}
