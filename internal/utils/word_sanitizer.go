package utils

import (
	"net/url"
	"regexp"
	"strings"
)

func SanitizeForURL(word string) string {
	// Trim leading and trailing spaces
	word = strings.TrimSpace(word)
	
	// Replace spaces with hyphens
	word = strings.ReplaceAll(word, " ", "-")

	// Remove non-alphanumeric characters except hyphens
	reg := regexp.MustCompile("[^a-zA-Z0-9-]+")
	word = reg.ReplaceAllString(word, "")

	// Encode the sanitized word
	sanitizedURL := url.QueryEscape(word)

	return sanitizedURL
}
