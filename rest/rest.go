package rest

import c "github.com/billcoding/flygo/context"

type Controller interface {
	Prefix() string
	GET() func(c *c.Context)
	GETS() func(c *c.Context)
	POST() func(c *c.Context)
	PUT() func(c *c.Context)
	DELETE() func(c *c.Context)
}
