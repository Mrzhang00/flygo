package middleware

import (
	c "github.com/billcoding/flygo/context"
	"net/http"
)

// Middleware interface
type Middleware interface {
	Type() *Type
	Name() string
	Method() Method
	Pattern() Pattern
	Handler() func(c *c.Context)
}

type (
	// Type type
	Type struct{ t string }
	// Method type
	Method string
	// Pattern type
	Pattern string
)

var (
	// TypeBefore type
	TypeBefore = &Type{t: "BEFORE"}
	// TypeHandler type
	TypeHandler = &Type{t: "HANDLER"}
	// TypeAfter type
	TypeAfter = &Type{t: "AFTER"}
)

const (
	// PatternNoRoute pattern
	PatternNoRoute = Pattern("")
	// PatternAny pattern
	PatternAny = Pattern("/*")
	// MethodAny method
	MethodAny = Method("*")
	// MethodGet method
	MethodGet = Method(http.MethodGet)
	// MethodPost method
	MethodPost = Method(http.MethodPost)
	// MethodPut method
	MethodPut = Method(http.MethodPut)
	// MethodDelete method
	MethodDelete = Method(http.MethodDelete)
	// MethodPatch method
	MethodPatch = Method(http.MethodPatch)
)

// SetMWData set middleware data
func SetMWData(c *c.Context, name string, mwData map[string]interface{}) {
	c.MWData[name] = mwData
}

// GetMWData get middleware data
func GetMWData(c *c.Context, name string) map[string]interface{} {
	val, have := c.MWData[name]
	if have {
		return val.(map[string]interface{})
	}
	return nil
}
