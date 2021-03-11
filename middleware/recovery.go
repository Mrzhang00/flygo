package middleware

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/mime"
	"net/http"
	"runtime"
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
	return RecoveryWithConfig("code", 500, "message", handlers...)
}

// RecoveryWithConfig return new recovery
func RecoveryWithConfig(codeName string, codeVal int, msgName string, handlers ...func(c *c.Context)) Middleware {
	var handler func(ctx *c.Context)
	if len(handlers) > 0 && handlers[0] != nil {
		handler = handlers[0]
	} else {
		handler = func(ctx *c.Context) {
			defer func() {
				if re := recover(); re != nil {
					fmt.Println(fmt.Sprintf("[Recovered]%v", re))
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					fmt.Printf("%s\n", string(buf[:n]))
					message := ""
					switch re.(type) {
					case string:
						message = fmt.Sprintf("%v", re)
					case error:
						message = re.(error).Error()
					default:
						message = fmt.Sprintf("%v", re)
					}
					ctx.Render(c.RenderBuilder().Buffer([]byte(fmt.Sprintf(`{"%s":%d,"%s":"%s"}`, codeName, codeVal, msgName, message))).ContentType(mime.JSON).Code(http.StatusInternalServerError).Build())
				}
			}()
			ctx.Chain()
		}
	}
	return &recovery{handler}
}
