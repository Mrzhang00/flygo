package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

//Define Middleware interface
type Middleware interface {
	Type() *Type                 //The type for middleware
	Name() string                //The name for middleware
	Method() Method              //The method for middleware
	Pattern() Pattern            //The pattern for middleware
	Handler() func(c *c.Context) //The handler for middleware
}

type (
	//Define Type struct
	Type struct {
		t string
	}

	//Define Method type
	Method string

	//Define Pattern type
	Pattern string
)

var (
	TypeBefore  = &Type{t: "BEFORE"}
	TypeHandler = &Type{t: "HANDLER"}
	TypeAfter   = &Type{t: "AFTER"}
)

const (
	PatternNoRoute = Pattern("")
	PatternAny     = Pattern("/*")
	MethodAny      = Method("*")
	MethodGet      = Method(http.MethodGet)
	MethodPost     = Method(http.MethodPost)
	MethodPut      = Method(http.MethodPut)
	MethodDelete   = Method(http.MethodDelete)
	MethodPatch    = Method(http.MethodPatch)
)

//SetMWData
func SetMWData(c *c.Context, name string, mwData map[string]interface{}) {
	c.MWData[name] = mwData
}

//GetMWData
func GetMWData(c *c.Context, name string) map[string]interface{} {
	val, have := c.MWData[name]
	if have {
		return val.(map[string]interface{})
	}
	return nil
}
