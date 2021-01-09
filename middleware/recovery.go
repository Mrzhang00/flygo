package middleware

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"net/http"
)

type recovery struct {
	handler func(c *c.Context)
}

func (r *recovery) Type() *Type {
	return TypeBefore
}

func (r *recovery) Name() string {
	return "Recovered"
}

func (r *recovery) Method() Method {
	return MethodAny
}

func (r *recovery) Pattern() Pattern {
	return PatternNoRoute
}

func (r *recovery) Handler() func(c *c.Context) {
	return r.handler
}

func Recovery(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 {
		return &recovery{handlers[0]}
	}
	return &recovery{recoveryHandler}
}

var recoveryHandler = func(c *c.Context) {
	defer func() {
		if re := recover(); re != nil {
			_ = fmt.Errorf("[Recovered]%v\n", re)
			c.Header().Set(headers.MIME, mime.JSON)
			c.WriteCode(http.StatusInternalServerError)
			c.Write([]byte(fmt.Sprintf(`{"code":500,"msg":"%s"}`, re)))
		}
	}()
	c.Chain()
}
