package flygo

import (
	"regexp"
)

//match path
func (c *Context) matchPath(regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(c.ParsedRequestURI)
}

//c matches
func (c *Context) contextMatches(regex string) []string {
	return findMatches(c.ParsedRequestURI, regex)
}

//find matches
func findMatches(str, regex string) []string {
	reg := regexp.MustCompile(regex)
	return reg.FindAllString(str, -1)
}

//is variable route
func isVariableRoute(pattern string) bool {
	matched, _ := regexp.MatchString(`{[^{]+}`, pattern)
	return matched
}
