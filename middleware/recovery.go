package middleware

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/mime"
	"net/http"
)

type recovery struct {
	handler func(c *c.Context)
}

// Type implements
func (r *recovery) Type() *Type {
	return TypeBefore
}

// Name implements
func (r *recovery) Name() string {
	return "Recovered"
}

// Method implements
func (r *recovery) Method() Method {
	return MethodAny
}

// Pattern implements
func (r *recovery) Pattern() Pattern {
	return PatternNoRoute
}

// Handler implements
func (r *recovery) Handler() func(c *c.Context) {
	return r.handler
}

// Recovery return new recovery
func Recovery(handlers ...func(c *c.Context)) Middleware {
	if len(handlers) > 0 {
		return &recovery{handlers[0]}
	}
	return &recovery{recoveryHandler}
}

var recoveryHandler = func(ctx *c.Context) {
	defer func() {
		if re := recover(); re != nil {
			_ = fmt.Errorf("[Recovered]%v\n", re)
			ctx.Render(c.RenderBuilder().Buffer([]byte(fmt.Sprintf(`{"code":500,"msg":"%s"}`, re))).ContentType(mime.JSON).Code(http.StatusInternalServerError).Build())
		}
	}()
	ctx.Chain()
}
