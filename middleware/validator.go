package middleware

import (
	c "github.com/billcoding/flygo/context"
	vali "github.com/billcoding/flygo/validator"
	"reflect"
)

//Define validator struct
type validator struct {
	defaultMsg  string
	defaultCode int
}

//Validator
func Validator(defaultMsg string, defaultCode int) Middleware {
	return &validator{
		defaultMsg:  defaultMsg,
		defaultCode: defaultCode,
	}
}

//Type
func (v *validator) Type() *Type {
	return TypeBefore
}

//Name
func (v *validator) Name() string {
	return "Validator"
}

//Method
func (v *validator) Method() Method {
	return MethodAny
}

//Pattern
func (v *validator) Pattern() Pattern {
	return PatternNoRoute
}

//Handler
func (v *validator) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		validatorData, have := c.MWData["Validator"]
		if !have {
			return
		}
		if reflect.TypeOf(validatorData).Kind() != reflect.Ptr &&
			reflect.TypeOf(validatorData).Elem().Kind() != reflect.Struct {
			result := vali.New(validatorData).Validate()
			if !result.Passed {
				msg := result.Messages()
				if msg == "" {
					msg = v.defaultMsg
				}
				c.JSON(map[string]interface{}{"message": msg, "code": v.defaultCode})
				return
			}
		}
		c.Chain()
	}
}
