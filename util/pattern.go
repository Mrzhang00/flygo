package util

import "regexp"

// TrimSpecialChars string
func TrimSpecialChars(str string) string {
	re := regexp.MustCompile(`[^/\w-._{}]`)
	return re.ReplaceAllString(str, "")
}
