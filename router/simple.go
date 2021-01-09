package router

import c "github.com/billcoding/flygo/context"

type Simple struct {
	Method  string
	Pattern string
	Handler func(c *c.Context)
}
