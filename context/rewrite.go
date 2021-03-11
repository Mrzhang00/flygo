package context

import (
	"github.com/billcoding/flygo/headers"
	"net/http"
)

// Redirect url
func (ctx *Context) Redirect(url string) {
	ctx.render = RenderBuilder().Header(http.Header{
		headers.Location: []string{url},
	}).Code(http.StatusTemporaryRedirect).Build()
}
