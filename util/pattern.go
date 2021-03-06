package util

import (
	"regexp"
	"strings"
)

// TrimSpecialChars string
func TrimSpecialChars(str string) string {
	str = strings.TrimRight(str, "/")
	re := regexp.MustCompile(`[^/\w-._{}]`)
	return re.ReplaceAllString(str, "")
}
