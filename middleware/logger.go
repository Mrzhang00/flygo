package middleware

import (
	c "github.com/billcoding/flygo/context"
	l "github.com/billcoding/flygo/log"
)

type logger struct {
	logger  l.Logger
	handler func(c *c.Context)
}

// Type implements
func (l *logger) Type() *Type {
	return TypeBefore
}

// Name implements
func (l *logger) Name() string {
	return "StdLogger"
}

// Method implements
func (l *logger) Method() Method {
	return MethodAny
}

// Pattern implements
func (l *logger) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (l *logger) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		c.Chain()
		l.logger.Info("[%s]%s => %d", c.Request.Method, c.Request.URL.Path, c.Rendered().Code)
	}
}

// StdLogger return logger
func StdLogger() Middleware {
	return &logger{
		logger: l.New("[StdLogger]"),
	}
}
