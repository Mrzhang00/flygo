package context

import (
	"github.com/billcoding/flygo/mime"
	"net/http"
)

// Builder struct
type Builder struct {
	r *Render
}

// Build render
func (r *Builder) Build() *Render {
	return r.r
}

// RenderBuilder render
func RenderBuilder() *Builder {
	return &Builder{r: &Render{Code: 200, ContentType: mime.TEXT}}
}

// DefaultBuild render
func (r *Builder) DefaultBuild() *Render {
	return r.Header(http.Header{}).Cookies(make([]*http.Cookie, 0)).Code(200).ContentType(mime.TEXT).Build()
}

// Buffer render
func (r *Builder) Buffer(buffer []byte) *Builder {
	r.r.Buffer = buffer
	return r
}

// Header render
func (r *Builder) Header(header http.Header) *Builder {
	r.r.Header = header
	return r
}

// Cookies render
func (r *Builder) Cookies(cookies []*http.Cookie) *Builder {
	r.r.Cookies = cookies
	return r
}

// Code render
func (r *Builder) Code(code int) *Builder {
	r.r.Code = code
	return r
}

// ContentType render
func (r *Builder) ContentType(contentType string) *Builder {
	r.r.ContentType = contentType
	return r
}
