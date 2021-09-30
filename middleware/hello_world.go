package middleware

import (
	"fmt"
	"github.com/billcoding/flygo/context"
)

type helloWorld struct {
}

// HelloWorld return new helloWorld
func HelloWorld() *helloWorld {
	return &helloWorld{}
}

// Name implements
func (h *helloWorld) Name() string {
	return "helloWorld"
}

// Type implements
func (h *helloWorld) Type() *Type {
	return TypeBefore
}

// Method implements
func (h *helloWorld) Method() Method {
	return MethodAny
}

// Pattern implements
func (h *helloWorld) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (h *helloWorld) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		fmt.Println("hello world")
	}
}
