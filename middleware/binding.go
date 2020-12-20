package middleware

import (
	c "github.com/billcoding/flygo/context"
)

//Define header struct
type binding struct {
}

//Type
func (b *binding) Type() *Type {
	return TypeBefore
}

//Name
func (b *binding) Name() string {
	panic("implement me")
}

//Method
func (b *binding) Method() Method {
	panic("implement me")
}

//Pattern
func (b *binding) Pattern() Pattern {
	panic("implement me")
}

//Handler
func (b *binding) Handler() func(c *c.Context) {
	panic("implement me")
}
