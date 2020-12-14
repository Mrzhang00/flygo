package context

import (
	"github.com/billcoding/flygo/mime"
	"net/http"
)

//Define builder struct
type builder struct {
	r *Render
}

//Build
func (r *builder) Build() *Render {
	return r.r
}

//RenderBuilder
func RenderBuilder() *builder {
	return &builder{r: &Render{Code: 200, ContentType: mime.TEXT}}
}

//DefaultBuild
func (r *builder) DefaultBuild() *Render {
	return r.Header(http.Header{}).Cookies(make([]*http.Cookie, 0)).ContentType(mime.TEXT).Build()
}

//Buffer
func (r *builder) Buffer(buffer []byte) *builder {
	r.r.Buffer = buffer
	return r
}

//Header
func (r *builder) Header(header http.Header) *builder {
	r.r.Header = header
	return r
}

//Cookies
func (r *builder) Cookies(cookies []*http.Cookie) *builder {
	r.r.Cookies = cookies
	return r
}

//Code
func (r *builder) Code(code int) *builder {
	r.r.Code = code
	return r
}

//ContentType
func (r *builder) ContentType(contentType string) *builder {
	r.r.ContentType = contentType
	return r
}
