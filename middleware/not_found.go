package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

//Define notFound struct
type notFound struct {
}

//Name
func (n *notFound) Name() string {
	return "NotFound"
}

//Method
func (n *notFound) Type() *Type {
	return TypeAfter
}

//Method
func (n *notFound) Method() Method {
	return MethodAny
}

//Pattern
func (n *notFound) Pattern() Pattern {
	return PatternAny
}

//Handler
func (n *notFound) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		if !c.Wrote {
			c.WriteHeader(http.StatusNotFound)
			c.Write([]byte("404 Not Found"))
		} else {
			c.Chain()
		}
	}
}

//NotFound
func NotFound() Middleware {
	return &notFound{}
}
