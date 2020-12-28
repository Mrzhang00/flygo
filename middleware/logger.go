package middleware

import (
	c "github.com/billcoding/flygo/context"
	l "github.com/billcoding/flygo/log"
)

//Define logger struct
type logger struct {
	logger  l.Logger
	handler func(c *c.Context)
}

//Type
func (l *logger) Type() *Type {
	return TypeBefore
}

//Name
func (l *logger) Name() string {
	return "StdLogger"
}

//Method
func (l *logger) Method() Method {
	return MethodAny
}

//Pattern
func (l *logger) Pattern() Pattern {
	return PatternNoRoute
}

//Handler
func (l *logger) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		c.Chain()
		l.logger.Info("[%s]%s => %d", c.Request.Method, c.Request.URL.Path, c.Render().Code)
	}
}

//StdLogger
func StdLogger() Middleware {
	return &logger{
		logger: l.New("[StdLogger]"),
	}
}
