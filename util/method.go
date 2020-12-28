package util

import "net/http"

var allMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodPatch,
	http.MethodHead,
	http.MethodOptions,
}

var routeMethods = map[string]int{
	"*":                0,
	http.MethodGet:     0,
	http.MethodPost:    0,
	http.MethodPut:     0,
	http.MethodDelete:  0,
	http.MethodPatch:   0,
	http.MethodHead:    0,
	http.MethodOptions: 0,
}

var requestMethods = map[string]int{
	http.MethodGet:     0,
	http.MethodHead:    0,
	http.MethodPost:    0,
	http.MethodPut:     0,
	http.MethodPatch:   0,
	http.MethodDelete:  0,
	http.MethodConnect: 0,
	http.MethodOptions: 0,
	http.MethodTrace:   0,
}

//RouteSupport
func RouteSupport(method string) bool {
	_, have := routeMethods[method]
	return have
}

//RequestSupport
func RequestSupport(method string) bool {
	_, have := requestMethods[method]
	return have
}

//AllMethods
func AllMethods() []string {
	return allMethods
}
