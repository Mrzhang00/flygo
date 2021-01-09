package rest

import c "github.com/billcoding/flygo/context"

// Controller interface
type Controller interface {
	// Prefix route pattern prefix
	Prefix() string
	// GET route
	GET() func(c *c.Context)
	// GETS route
	GETS() func(c *c.Context)
	// POST route
	POST() func(c *c.Context)
	// PUT route
	PUT() func(c *c.Context)
	// DELETE route
	DELETE() func(c *c.Context)
}
