package middleware

import (
	"github.com/billcoding/flygo/context"
	"github.com/sirupsen/logrus"
)

type logger struct {
	logger  *logrus.Logger
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
		l.logger.WithFields(map[string]interface{}{
			"Host":       ctx.Request.Host,
			"Proto":      ctx.Request.Proto,
			"Method":     ctx.Request.Method,
			"URL":        ctx.Request.URL,
			"RequestURI": ctx.Request.RequestURI,
			"RemoteAddr": ctx.Request.RemoteAddr,
			"Header":     ctx.Request.Header,
			"Form":       ctx.Request.Form,
			"PostForm":   ctx.Request.PostForm,
		}).Infof("Requested %s", ctx.Request.RequestURI)
		ctx.Chain()
	}
}

// StdLogger return logger
func StdLogger() Middleware {
	return &logger{
		logger: logrus.StandardLogger(),
	}
}
