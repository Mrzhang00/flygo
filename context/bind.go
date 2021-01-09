package context

import (
	bind "github.com/billcoding/binding"
	"github.com/billcoding/calls"
	vali "github.com/billcoding/validator"
)

// Bind struct
func (c *Context) Bind(structPtr interface{}) {
	bind.New(structPtr).BindReq(c.Request)
}

// BindWithType struct
func (c *Context) BindWithType(structPtr interface{}, typ *bind.Type) {
	bind.NewWithType(structPtr, typ).BindReq(c.Request)
}

// Validate struct
func (c *Context) Validate(structPtr interface{}, call func()) {
	c.ValidateWithParams(structPtr, "Parameters is invalid", 500, call)
}

// ValidateWithParams struct
func (c *Context) ValidateWithParams(structPtr interface{}, message string, code int, call func()) {
	result := vali.New(structPtr).Validate()
	calls.True(result.Passed, call)
	calls.False(result.Passed, func() {
		msg := result.Messages()
		if msg == "" {
			msg = message
		}
		c.JSON(map[string]interface{}{"message": msg, "code": code})
	})
}

// BindAndValidate struct
func (c *Context) BindAndValidate(structPtr interface{}, call func()) {
	bind.New(structPtr).BindReq(c.Request)
	c.Validate(structPtr, call)
}

// BindWithParamsAndValidate struct
func (c *Context) BindWithParamsAndValidate(structPtr interface{}, typ *bind.Type, call func()) {
	bind.NewWithType(structPtr, typ).BindReq(c.Request)
	c.Validate(structPtr, call)
}

// BindWithParamsAndValidateWithParams struct
func (c *Context) BindWithParamsAndValidateWithParams(structPtr interface{}, typ *bind.Type, message string, code int, call func()) {
	bind.NewWithType(structPtr, typ).BindReq(c.Request)
	c.ValidateWithParams(structPtr, message, code, call)
}
