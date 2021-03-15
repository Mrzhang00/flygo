package middleware

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/billcoding/flygo/util"
)

type logger struct {
	logger  log.Logger
	handler func(ctx *context.Context)
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
	return PatternAny
}

// Handler implements
func (l *logger) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		ctx.Chain()
		l.logger.Info("[%s]%s => %d", ctx.Request.Method, util.TrimLeftAndRight(ctx.Request.URL.Path), ctx.Rendered().Code)
	}
}

// StdLogger return logger
func StdLogger() Middleware {
	return &logger{
		logger: log.New("[StdLogger]"),
	}
}
