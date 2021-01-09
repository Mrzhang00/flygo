package middleware

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"runtime"
)

type header struct {
	headers map[string]string
}

// Type implements
func (h *header) Type() *Type {
	return TypeBefore
}

// Name implements
func (h *header) Name() string {
	return "Header"
}

// Method implements
func (h *header) Method() Method {
	return MethodAny
}

// Pattern implements
func (h *header) Pattern() Pattern {
	return PatternAny
}

// Handler implements
func (h *header) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		for k, v := range h.headers {
			c.Header().Set(k, v)
		}
		c.Chain()
	}
}

// Header return header
func Header() Middleware {
	return &header{map[string]string{
		headers.Server:  fmt.Sprintf("golang/flygo (%s)", runtime.Version()),
		headers.XServer: fmt.Sprintf("golang/flygo (%s)", runtime.Version()),
	}}
}
