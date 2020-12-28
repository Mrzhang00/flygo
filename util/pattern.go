package util

import "regexp"

//TrimSpecialChars
func TrimSpecialChars(str string) string {
	re := regexp.MustCompile(`[^/\w-._{}]`)
	return re.ReplaceAllString(str, "")
}
