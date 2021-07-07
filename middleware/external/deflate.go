package external

import (
	"bytes"
	"compress/flate"
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/mime"
	"strings"
)

type deflate struct {
	contentType []string
	minSize     int
	level       int
}

func NewDeflate() *deflate {
	return &deflate{
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
		level:   flate.BestSpeed,
	}
}

//Type
func (d *deflate) Type() *middleware.Type {
	return middleware.TypeAfter
}

func (d *deflate) Name() string {
	return "Deflate"
}

func (d *deflate) Method() middleware.Method {
	return middleware.MethodAny
}

func (d *deflate) Pattern() middleware.Pattern {
	return middleware.PatternAny
}

func (d *deflate) accept(ctx *context.Context) bool {
	acceptEncoding := ctx.Request.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "deflate")
}

func (d *deflate) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		rendered := ctx.Rendered()
		if d.accept(ctx) && rendered != nil {
			odata := rendered.Buffer
			if nil != odata && len(odata) >= d.minSize {
				ct := rendered.ContentType
				if strings.Index(ct, ";") != -1 {
					ct = strings.TrimSpace(strings.Split(ct, ";")[0])
				}
				if ct == "" {
					ct = mime.BINARY
				}
				ctx.Header().Set("Vary", "Content-Encoding")
				ctx.Header().Set("Content-Encoding", "deflate")
				var buffers bytes.Buffer
				fw, err := flate.NewWriter(&buffers, d.level)
				defer fw.Close()
				if err != nil {
					panic(err)
				}
				{
					_, err := fw.Write(odata)
					if err != nil {
						panic(err)
					}
				}
				fw.Flush()
				ctx.Write(buffers.Bytes())
			}
		}
		ctx.Chain()
	}
}

func (d *deflate) ContentType(contentType ...string) *deflate {
	d.contentType = contentType
	return d
}

func (d *deflate) MinSize(minSize int) *deflate {
	d.minSize = minSize
	return d
}

func (d *deflate) Level(level int) *deflate {
	d.level = level
	return d
}
