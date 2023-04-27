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
	sanitized = regexp.MustCompile(`一(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "1-")
	sanitized = regexp.MustCompile(`二(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "2-")
	sanitized = regexp.MustCompile(`三(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "3-")
	sanitized = regexp.MustCompile(`四(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "4-")
	sanitized = regexp.MustCompile(`五(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "5-")
	sanitized = regexp.MustCompile(`六(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "6-")
	sanitized = regexp.MustCompile(`七(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "7-")
	sanitized = regexp.MustCompile(`八(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "8-")
	sanitized = regexp.MustCompile(`九(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "9-")
	sanitized = regexp.MustCompile(`([0-9]+)ー([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)ー([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)(?:番地|-)?の([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "$1-")
	sanitized = regexp.MustCompile(`-([^0-9]|$)`).ReplaceAllString(sanitized, "$1")

	sanitized = regexp.MustCompile(`イオン.+$`).ReplaceAllString(sanitized, "")

	// Remove building name after street name
	sanitized = regexp.MustCompile(`([0-9]+(?:-[0-9]+)?(?:-[0-9]+)?)[^-0-9条線].*$`).ReplaceAllString(sanitized, "$1")

	sanitized = strings.TrimSpace(sanitized)

	return sanitized
}
