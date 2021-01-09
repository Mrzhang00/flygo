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

func (h *header) Type() *Type {
	return TypeBefore
}

func (h *header) Name() string {
	return "Header"
}

func (h *header) Method() Method {
	return MethodAny
}

func (h *header) Pattern() Pattern {
	return PatternAny
}

func (h *header) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		for k, v := range h.headers {
			c.Header().Set(k, v)
		}
		c.Chain()
	}
}

func Header() Middleware {
	return &header{map[string]string{
		headers.Server:  fmt.Sprintf("golang/flygo (%s)", runtime.Version()),
		headers.XServer: fmt.Sprintf("golang/flygo (%s)", runtime.Version()),
	}}
}
