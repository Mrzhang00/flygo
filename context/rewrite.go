package context

import (
	"github.com/billcoding/flygo/headers"
	"net/http"
)

func (c *Context) Redirect(url string) {
	c.render = RenderBuilder().Header(http.Header{
		headers.Location: []string{url},
	}).Code(http.StatusTemporaryRedirect).Build()
}
