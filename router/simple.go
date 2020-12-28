package router

import c "github.com/billcoding/flygo/context"

//Define Simple struct
type Simple struct {
	Method  string             //Method
	Pattern string             //Pattern
	Handler func(c *c.Context) //context
}
