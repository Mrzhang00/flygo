package context

import (
	"github.com/billcoding/flygo/mime"
	"net/http"
)

//Define Builder struct
type Builder struct {
	r *Render
}

//Build
func (r *Builder) Build() *Render {
	if r.r.Header == nil {
		r.r.Header = make(map[string][]string, 0)
	}
	if r.r.Cookies == nil {
		r.r.Cookies = make([]*http.Cookie, 0)
	}
	return r.r
}

//RenderBuilder
func RenderBuilder() *Builder {
	return &Builder{r: &Render{Code: 200, ContentType: mime.TEXT}}
}

//DefaultBuild
func (r *Builder) DefaultBuild() *Render {
	return r.Header(http.Header{}).Cookies(make([]*http.Cookie, 0)).Code(200).ContentType(mime.TEXT).Build()
}

//Buffer
func (r *Builder) Buffer(buffer []byte) *Builder {
	r.r.Buffer = buffer
	return r
}

//Header
func (r *Builder) Header(header http.Header) *Builder {
	r.r.Header = header
	return r
}

//Cookies
func (r *Builder) Cookies(cookies []*http.Cookie) *Builder {
	r.r.Cookies = cookies
	return r
}

//Code
func (r *Builder) Code(code int) *Builder {
	r.r.Code = code
	return r
}

//ContentType
func (r *Builder) ContentType(contentType string) *Builder {
	r.r.ContentType = contentType
	return r
}
