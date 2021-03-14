package util

import (
	"fmt"
	"regexp"
	"strings"
)

// TrimSpecialChars string
func TrimSpecialChars(str string) string {
	str = TrimLeftAndRight(str)
	re := regexp.MustCompile(`[^/\w-._{}]`)
	return re.ReplaceAllString(str, "")
}

func TrimLeftAndRight(pattern string) string {
	pattern = strings.TrimLeft(pattern, "//")
	pattern = strings.TrimRight(pattern, "//")
	pattern = strings.TrimLeft(pattern, "/")
	pattern = strings.TrimRight(pattern, "/")
	pattern = "/" + pattern
	return pattern
}

func TrimPattern(pattern string) string {
	pattern = TrimLeftAndRight(pattern)
	re := regexp.MustCompile(`[^/\w-._*]`)
	np := re.ReplaceAllString(pattern, "")
	np = strings.ReplaceAll(np, "**", "*")
	np = strings.ReplaceAll(np, "*", `[\w-._/]+`)
	return fmt.Sprintf(`^%s$`, np)
}
