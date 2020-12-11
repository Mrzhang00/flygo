package context

import (
	"github.com/billcoding/flygo/headers"
	"net/http"
)

//Redirect
func (c *Context) Redirect(url string) {
	c.Header().Set(headers.Location, url)
	c.WriteHeader(http.StatusTemporaryRedirect)
}
