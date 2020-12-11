package rest

import c "github.com/billcoding/flygo/context"

//Define Controller interface
type Controller interface {
	Prefix() string //Prefix
	// /GET: {Prefix}/{id}
	GET() func(c *c.Context) //Get one service
	// /GET: {Prefix}
	GETS() func(c *c.Context) //Get list service
	// /POST: {Prefix}
	POST() func(c *c.Context) //Post service
	// /PUT: {Prefix}
	PUT() func(c *c.Context) //Put service
	// /DELETE: {Prefix}/{id}
	DELETE() func(c *c.Context) //Delete service
}
