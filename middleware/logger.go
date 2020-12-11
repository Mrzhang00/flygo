package middleware

import (
	c "github.com/billcoding/flygo/context"
	l "github.com/billcoding/flygo/log"
)

//Define logger struct
type logger struct {
	l.Logger
}

//Type
func (l *logger) Type() *Type {
	return TypeBefore
}

//Name
func (l *logger) Name() string {
	return "Logger"
}

//Method
func (l *logger) Method() Method {
	return MethodAny
}

//Pattern
func (l *logger) Pattern() Pattern {
	return PatternAny
}

//Handler
func (l *logger) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		l.Logger.Info("%s => %s", c.Request.Method, c.Request.URL.Path)
		c.Chain()
	}
}

func StdLogger() Middleware {
	return &logger{
		Logger: l.New("[Middleware:Logger]"),
	}
}
