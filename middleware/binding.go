package middleware

import (
	bind "github.com/billcoding/flygo/binding"
	c "github.com/billcoding/flygo/context"
	"reflect"
)

//Define binding struct
type binding struct {
}

//Binding
func Binding() Middleware {
	return &binding{}
}

//Type
func (b *binding) Type() *Type {
	return TypeBefore
}

//Name
func (b *binding) Name() string {
	return "Binding"
}

//Method
func (b *binding) Method() Method {
	return MethodAny
}

//Pattern
func (b *binding) Pattern() Pattern {
	return PatternNoRoute
}

//Handler
func (b *binding) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		bindingData, have := c.MWData["Binding"]
		if !have {
			return
		}
		if reflect.TypeOf(bindingData).Kind() != reflect.Ptr &&
			reflect.TypeOf(bindingData).Elem().Kind() != reflect.Struct {
			bind.New(bindingData).Bind(c.Request)
		}
		c.Chain()
	}
}
