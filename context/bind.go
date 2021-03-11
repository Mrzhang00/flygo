package context

import (
	"encoding/json"
	v "github.com/billcoding/validator"
	"io/ioutil"
)

// Bind struct ptr
func (ctx *Context) Bind(structPtr interface{}) {
	readAll, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(readAll, structPtr)
	if err != nil {
		panic(err)
	}
}

// Validate struct ptr
func (ctx *Context) Validate(structPtr interface{}, call func()) {
	ctx.ValidateWithParams(structPtr, "Parameters is invalid", 500, call)
}

// ValidateWithParams struct ptr
func (ctx *Context) ValidateWithParams(structPtr interface{}, message string, code int, call func()) {
	result := v.New(structPtr).Validate()
	if result.Passed {
		call()
	} else {
		msg := result.Messages()
		if msg == "" {
			msg = message
		}
		ctx.JSON(map[string]interface{}{"message": msg, "code": code})
	}
}

// BindAndValidate struct ptr
func (ctx *Context) BindAndValidate(structPtr interface{}, call func()) {
	ctx.Bind(structPtr)
	ctx.Validate(structPtr, call)
}

// BindAndValidateWithParams struct ptr
func (ctx *Context) BindAndValidateWithParams(structPtr interface{}, message string, code int, call func()) {
	ctx.Bind(structPtr)
	ctx.ValidateWithParams(structPtr, message, code, call)
}
