package flygo

import "net/http"

//Redirect url
func (c *Context) Redirect(url string) {
	c.SetHeader("Location", url)
	c.ResponseWriter.WriteHeader(http.StatusTemporaryRedirect)
}
