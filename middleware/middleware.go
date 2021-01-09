package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

type Middleware interface {
	Type() *Type
	Name() string
	Method() Method
	Pattern() Pattern
	Handler() func(c *c.Context)
}

type (
	Type struct {
		t string
	}

	Method string

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

func SetMWData(c *c.Context, name string, mwData map[string]interface{}) {
	c.MWData[name] = mwData
}

func GetMWData(c *c.Context, name string) map[string]interface{} {
	val, have := c.MWData[name]
	if have {
		return val.(map[string]interface{})
	}
	return nil
}
