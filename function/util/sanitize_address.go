package util

import (
	"golang.org/x/text/width"
	"regexp"
	"strings"
)

// SanitizeAddress returns sanitized address for geocoding
func SanitizeAddress(address string) string {
	sanitized := width.Fold.String(address)

	sanitized = strings.ReplaceAll(sanitized, "−", "-")

	// Normalize Japanese street number(丁目,番地,号)
	sanitized = regexp.MustCompile(`([0-9]+)ー([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)(?:番地)?の([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "$1-")

	// Remove building name after street name
	sanitized = regexp.MustCompile(`([0-9]+(?:-[0-9]+)?(?:-[0-9]+)?)[^条線]*$`).ReplaceAllString(sanitized, "$1")

	return sanitized
}
