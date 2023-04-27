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
	sanitized = regexp.MustCompile(`四十二(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "42-")
	sanitized = regexp.MustCompile(`四十一(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "41-")
	sanitized = regexp.MustCompile(`四十(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "40-")
	sanitized = regexp.MustCompile(`三十九(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "39-")
	sanitized = regexp.MustCompile(`三十八(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "38-")
	sanitized = regexp.MustCompile(`三十七(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "37-")
	sanitized = regexp.MustCompile(`三十六(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "36-")
	sanitized = regexp.MustCompile(`三十五(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "35-")
	sanitized = regexp.MustCompile(`三十四(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "34-")
	sanitized = regexp.MustCompile(`三十三(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "33-")
	sanitized = regexp.MustCompile(`三十二(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "32-")
	sanitized = regexp.MustCompile(`三十一(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "31-")
	sanitized = regexp.MustCompile(`三十(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "30-")
	sanitized = regexp.MustCompile(`二十九(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "29-")
	sanitized = regexp.MustCompile(`二十八(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "28-")
	sanitized = regexp.MustCompile(`二十七(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "27-")
	sanitized = regexp.MustCompile(`二十六(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "26-")
	sanitized = regexp.MustCompile(`二十五(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "25-")
	sanitized = regexp.MustCompile(`二十四(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "24-")
	sanitized = regexp.MustCompile(`二十三(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "23-")
	sanitized = regexp.MustCompile(`二十二(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "22-")
	sanitized = regexp.MustCompile(`二十一(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "21-")
	sanitized = regexp.MustCompile(`二十(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "20-")
	sanitized = regexp.MustCompile(`十九(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "19-")
	sanitized = regexp.MustCompile(`十八(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "18-")
	sanitized = regexp.MustCompile(`十七(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "17-")
	sanitized = regexp.MustCompile(`十六(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "16-")
	sanitized = regexp.MustCompile(`十五(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "15-")
	sanitized = regexp.MustCompile(`十四(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "14-")
	sanitized = regexp.MustCompile(`十三(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "13-")
	sanitized = regexp.MustCompile(`十二(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "12-")
	sanitized = regexp.MustCompile(`十一(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "11-")
	sanitized = regexp.MustCompile(`十(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "10-")
	sanitized = regexp.MustCompile(`九(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "9-")
	sanitized = regexp.MustCompile(`八(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "8-")
	sanitized = regexp.MustCompile(`七(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "7-")
	sanitized = regexp.MustCompile(`六(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "6-")
	sanitized = regexp.MustCompile(`五(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "5-")
	sanitized = regexp.MustCompile(`四(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "4-")
	sanitized = regexp.MustCompile(`三(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "3-")
	sanitized = regexp.MustCompile(`二(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "2-")
	sanitized = regexp.MustCompile(`一(?:(?:丁目)|(?:番地))`).ReplaceAllString(sanitized, "1-")

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
