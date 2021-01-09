package middleware

import (
	c "github.com/billcoding/flygo/context"
	l "github.com/billcoding/flygo/log"
)

type logger struct {
	logger  l.Logger
	handler func(c *c.Context)
}

func (l *logger) Type() *Type {
	return TypeBefore
}

func (l *logger) Name() string {
	return "StdLogger"
}

func (l *logger) Method() Method {
	return MethodAny
}

func (l *logger) Pattern() Pattern {
	return PatternNoRoute
}

func (l *logger) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		c.Chain()
		l.logger.Info("[%s]%s => %d", c.Request.Method, c.Request.URL.Path, c.Render().Code)
	}
}

func StdLogger() Middleware {
	return &logger{
		logger: l.New("[StdLogger]"),
	}
}
