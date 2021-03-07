package context

import (
	"encoding/json"
	v "github.com/billcoding/validator"
	"io/ioutil"
)

// Bind struct ptr
func (c *Context) Bind(structPtr interface{}) {
	readAll, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.logger.Error("%v", err)
	} else {
		jerr := json.Unmarshal(readAll, structPtr)
		if jerr != nil {
			c.logger.Error("%v", err)
		}
	}
}

// Validate struct ptr
func (c *Context) Validate(structPtr interface{}, call func()) {
	c.ValidateWithParams(structPtr, "Parameters is invalid", 500, call)
}

// ValidateWithParams struct ptr
func (c *Context) ValidateWithParams(structPtr interface{}, message string, code int, call func()) {
	result := v.New(structPtr).Validate()
	if result.Passed {
		call()
	} else {
		msg := result.Messages()
		if msg == "" {
			msg = message
		}
		c.JSON(map[string]interface{}{"message": msg, "code": code})
	}
}

// BindAndValidate struct ptr
func (c *Context) BindAndValidate(structPtr interface{}, call func()) {
	c.Bind(structPtr)
	c.Validate(structPtr, call)
}

// BindAndValidateWithParams struct ptr
func (c *Context) BindAndValidateWithParams(structPtr interface{}, message string, code int, call func()) {
	c.Bind(structPtr)
	c.ValidateWithParams(structPtr, message, code, call)
}
