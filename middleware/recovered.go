package middleware

import (
	"fmt"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"log"
	"net/http"
	"os"
)

//Defne recovered struct
type recovered struct {
}

//Type
func (r *recovered) Type() *Type {
	return TypeBefore
}

//Name
func (r *recovered) Name() string {
	return "Recovered"
}

//Method
func (r *recovered) Method() Method {
	return MethodAny
}

//Pattern
func (r *recovered) Pattern() Pattern {
	return PatternAny
}

//Handler
func (r *recovered) Handler() func(c *c.Context) {
	logger := log.New(os.Stderr, "[Recovered]", log.LstdFlags)
	return func(c *c.Context) {
		defer func() {
			if re := recover(); re != nil {
				logger.Printf("%v\n", re)
				c.Header().Set(headers.MIME, mime.JSON)
				c.WriteHeader(http.StatusInternalServerError)
				c.Write([]byte(fmt.Sprintf(`{"code":500,"msg":"%s"}`, re)))
			}
		}()
		c.Chain()
	}
}

//Recovered
func Recovered() Middleware {
	return &recovered{}
}
