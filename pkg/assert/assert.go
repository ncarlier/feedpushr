package assert

import "testing"

// Nil assert that an object is nil
func Nil(t *testing.T, actual interface{}, message string) {
	if message == "" {
		message = "Nil assertion failed"
	}
	if actual != nil {
		t.Fatalf("%s - actual: %s", message, actual)
	}
}

// NotNil assert that an object is not nil
func NotNil(t *testing.T, actual interface{}, message string) {
	if message == "" {
		message = "Not nil assertion failed"
	}
	if actual == nil {
		t.Fatalf("%s - actual: nil", message)
	}
}

// Equal assert that an object is equal to an expected value
func Equal(t *testing.T, expected interface{}, actual interface{}, message string) {
	if message == "" {
		message = "Equal assertion failed"
	}
	if actual != expected {
		t.Fatalf("%s - expected: %s, actual: %s", message, expected, actual)
	}
}

// NotEqual assert that an object is not equal to an expected value
func NotEqual(t *testing.T, expected interface{}, actual interface{}, message string) {
	if message == "" {
		message = "Not equal assertion failed"
	}
	if actual == expected {
		t.Fatalf("%s - unexpected: %s, actual: %s", message, expected, actual)
	}
}

// True assert that an expression is true
func True(t *testing.T, expression bool, message string) {
	if message == "" {
		message = "Expression is not true"
	}
	if !expression {
		t.Fatalf("%s : %v", message, expression)
	}
}
