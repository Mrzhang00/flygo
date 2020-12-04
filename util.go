package flygo

import (
	"regexp"
	"strings"
)

//trim quotas
func trim(pattern string) string {
	//trim quotas
	pattern = "/" + strings.TrimLeft(pattern, "/")
	reg := regexp.MustCompile("[^/*{}a-zA-Z0-9.]")
	pattern = reg.ReplaceAllString(pattern, "")
	//trim double `*`
	for strings.Contains(pattern, "**") {
		pattern = strings.ReplaceAll(pattern, "**", "*")
	}
	return pattern
}
