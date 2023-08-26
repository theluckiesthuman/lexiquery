package utils

import "testing"

func TestSanitizeForURL(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello, World!",
			expected: "Hello-World",
		},
		{
			input:    "This is a test string.",
			expected: "This-is-a-test-string",
		},
		{
			input:    "1234",
			expected: "1234",
		},
		{
			input:    "This is a test string with spaces and special characters!@#$%^&*()_+",
			expected: "This-is-a-test-string-with-spaces-and-special-characters",
		},
		{
			input:    "!@#$%^&*()_+",
			expected: "",
		},
	}

	for _, tc := range testCases {
		actual := SanitizeForURL(tc.input)
		if actual != tc.expected {
			t.Errorf("SanitizeForURL(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
