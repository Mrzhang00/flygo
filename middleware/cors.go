package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"net/http"
	"strings"
)

type cors struct {
	origin       string
	methods      []string
	allowHeaders []string
	header       http.Header
}

// Cors return new cors
func Cors() *cors {
	return &cors{
		origin:       "*",
		methods:      strings.Split("GET,POST,DELETE,PUT,PATCH,HEAD,OPTIONS", ","),
		allowHeaders: make([]string, 0),
		header:       make(http.Header, 0),
	}
}

// Name implements
func (cs *cors) Name() string {
	return "Cors"
}

// Type implements
func (cs *cors) Type() *Type {
	return TypeBefore
}

// Method implements
func (cs *cors) Method() Method {
	return MethodAny
}

// Pattern implements
func (cs *cors) Pattern() Pattern {
	return PatternAny
}

// Handler implements
func (cs *cors) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		cs.header.Set(headers.Allow, strings.Join(cs.methods, ","))
		cs.header.Set(headers.AccessControlAllowHeaders, strings.Join(cs.allowHeaders, ","))
		cs.header.Set(headers.AccessControlAllowOrigin, cs.origin)
		cs.header.Set(headers.AccessControlAllowMethods, strings.Join(cs.methods, ","))
		for k, v := range cs.header {
			for _, vv := range v {
				c.Header().Add(k, vv)
			}
		}
		if c.Request.Method != http.MethodHead && c.Request.Method != http.MethodOptions {
			c.Chain()
		}
	}
}

// Origin cors
func (cs *cors) Origin(origin string) *cors {
	cs.origin = origin
	return cs
}

// Methods cors
func (cs *cors) Methods(methods ...string) *cors {
	cs.methods = methods
	return cs
}

// AllowHeaders cors
func (cs *cors) AllowHeaders(headers ...string) *cors {
	cs.allowHeaders = headers
	return cs
}
