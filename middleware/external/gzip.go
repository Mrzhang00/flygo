package external

import (
	"bytes"
	gz "compress/gzip"
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/mime"
	"strings"
)

type gzip struct {
	contentType []string
	minSize     int
	level       int
}

func NewGzip() *gzip {
	return &gzip{
		contentType: []string{
			"application/javascript",
			"application/json",
			"application/xml",
			"text/javascript",
			"text/json",
			"text/xml",
			"text/plain",
			"text/xml",
			"html/css",
		},
		minSize: 2 << 9, //1KB
		level:   gz.BestSpeed,
	}
}

func (g *gzip) Name() string {
	return "Gzip"
}

func (g *gzip) Type() *middleware.Type {
	return middleware.TypeAfter
}

func (g *gzip) Method() middleware.Method {
	return middleware.MethodAny
}

func (g *gzip) Pattern() middleware.Pattern {
	return middleware.PatternAny
}

func (g *gzip) accept(ctx *context.Context) bool {
	acceptEncoding := ctx.Request.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "gzip")
}

func (g *gzip) Handler() func(c *context.Context) {
	return func(ctx *context.Context) {
		rendered := ctx.Rendered()
		if g.accept(ctx) && rendered != nil {
			odata := rendered.Buffer
			if nil != odata && len(odata) >= g.minSize {
				ct := rendered.ContentType
				if strings.Index(ct, ";") != -1 {
					ct = strings.TrimSpace(strings.Split(ct, ";")[0])
				}
				if ct == "" {
					ct = mime.BINARY
				}
				ctx.Header().Set("Vary", "Content-Encoding")
				ctx.Header().Set("Content-Encoding", "gzip")
				var buffers bytes.Buffer
				gw, err := gz.NewWriterLevel(&buffers, g.level)
				defer gw.Close()
				if err != nil {
					panic(err)
				}
				_, err = gw.Write(odata)
				if err != nil {
					panic(err)
				}
				gw.Flush()
				ctx.Write(buffers.Bytes())
			}
		}
		ctx.Chain()
	}
}

func (g *gzip) ContentType(contentType ...string) *gzip {
	g.contentType = contentType
	return g
}

func (g *gzip) MinSize(minSize int) *gzip {
	g.minSize = minSize
	return g
}

func (g *gzip) Level(level int) *gzip {
	g.level = level
	return g
}
