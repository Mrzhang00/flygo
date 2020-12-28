package middleware

import (
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"net/http"
	"strings"
)

//Define cors struct
type cors struct {
	origin  string
	methods []string
	header  http.Header
}

//Cors
func Cors() *cors {
	return &cors{
		origin:  "*",
		methods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		header:  make(http.Header, 0),
	}
}

//Name
func (cs *cors) Name() string {
	return "Cors"
}

//Type
func (cs *cors) Type() *Type {
	return TypeBefore
}

//Method
func (cs *cors) Method() Method {
	return MethodAny
}

//Pattern
func (cs *cors) Pattern() Pattern {
	return PatternAny
}

//Handler
func (cs *cors) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		cs.header.Set(headers.AccessControlAllowOrigin, cs.origin)
		cs.header.Set(headers.AccessControlAllowMethods, strings.Join(cs.methods, ","))
		for k, v := range cs.header {
			for _, vv := range v {
				c.Header().Add(k, vv)
			}
		}
		calls.True(c.Request.Method != http.MethodHead && c.Request.Method != http.MethodOptions, func() {
			c.Chain()
		})
	}
}

//Origin
func (cs *cors) Origin(origin string) *cors {
	cs.origin = origin
	return cs
}

//Methods
func (cs *cors) Methods(methods ...string) *cors {
	cs.methods = methods
	return cs
}
