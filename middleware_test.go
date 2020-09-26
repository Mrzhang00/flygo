package flygo

import (
	"fmt"
	"testing"
)

//Test login middleware
func TestLoginMiddleware(t *testing.T) {
	app := NewApp()
	lm := &loginMiddleware{}
	app.Use(lm)
	app.AfterInterceptor("/**", func(c *Context) {
		aa := c.GetMiddlewareData(lm.Name(), "aaaa")
		c.SetDone(false).Text(fmt.Sprintf("%v", aa))
	})
	app.Run()
}

//Test logging middleware
func TestLoggingMiddleware(t *testing.T) {
	app := NewApp()
	app.UseFilter(&loggingMiddleware{})
	app.Run()
}

//Test
func TestValidateMiddleware(t *testing.T) {
	app := NewApp()
	app.UseInterceptor(&validateMiddleware{})
	app.Run()
}

type loginMiddleware struct {
}

func (l *loginMiddleware) Name() string {
	return "loginMiddleware"
}

func (l *loginMiddleware) Method() string {
	return "GET"
}

func (l *loginMiddleware) Pattern() string {
	return "/middleware/login"
}

func (l *loginMiddleware) Process() Handler {
	return func(c *Context) {
		//c.Text("login middleware")
		c.SetMiddlewareData(l.Name(), "aaaa", "login")
		c.Text("success")
	}
}

func (l *loginMiddleware) Fields() []*Field {
	return nil
}

type loggingMiddleware struct {
}

func (*loggingMiddleware) Name() string {
	return "loggingMiddleware"
}

func (*loggingMiddleware) Type() string {
	return "AFTER"
}

func (*loggingMiddleware) Pattern() string {
	return "/**"
}

func (*loggingMiddleware) Process() FilterHandler {
	return func(c *FilterContext) {
		fmt.Println("logging")
	}
}

type validateMiddleware struct {
}

func (*validateMiddleware) Name() string {
	return "validateMiddleware"
}

func (*validateMiddleware) Type() string {
	return "BEFORE"
}

func (*validateMiddleware) Pattern() string {
	return "/**"
}

func (*validateMiddleware) Process() InterceptorHandler {
	return func(c *Context) {
		if len(c.GetParameters("name")) <= 0 {
			c.Text("Name is empty")
			return
		}
	}
}
