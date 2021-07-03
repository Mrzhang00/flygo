package middleware

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/mime"
	"github.com/billcoding/flygo/util"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type static struct {
	cache      bool
	caches     map[string][]byte
	mimeCaches map[string]string
	root       string
	mimes      map[string]string
}

// Static new static
func Static(cache bool, root string) *static {
	root, err := filepath.Abs(root)
	if err != nil {
		panic(err)
	}
	st := &static{
		cache:      cache,
		caches:     make(map[string][]byte, 0),
		mimeCaches: make(map[string]string, 0),
		root:       root,
		mimes:      defaultMimes(),
	}
	return st
}

func defaultMimes() map[string]string {
	return map[string]string{

		"jpg":  mime.JPG,
		"jpeg": mime.JPG,
		"png":  mime.PNG,
		"gif":  mime.GIF,
		"bmp":  mime.BMP,
		"ico":  mime.ICO,

		"txt":  mime.TEXT,
		"css":  mime.CSS,
		"js":   mime.JS,
		"json": mime.JSON,
		"xml":  mime.XML,
		"html": mime.HTML,

		"zip": mime.ZIP,
		"7z":  mime.ZIP7,
		"tar": mime.TAR,
		"gz":  mime.GZIP,
		"tgz": mime.TGZ,
		"rar": mime.RAR,

		"xls":  mime.XLS,
		"xlsx": mime.XLSX,
		"doc":  mime.DOC,
		"docx": mime.DOCX,
		"ppt":  mime.PPT,
		"pptx": mime.PPTX,

		"": mime.BINARY,
	}
}

// Type implements
func (s *static) Type() *Type {
	return TypeHandler
}

// Name implements
func (s *static) Name() string {
	return "Static"
}

// Method implements
func (s *static) Method() Method {
	return MethodGet
}

// Pattern implements
func (s *static) Pattern() Pattern {
	return "/static/*"
}

// Handler implements
func (s *static) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		if strings.HasSuffix(util.TrimLeftAndRight(ctx.Request.URL.Path), "/") {
			ctx.Chain()
			return
		}
		urlPath := strings.TrimPrefix(util.TrimLeftAndRight(ctx.Request.URL.Path), "/static")
		realPath := filepath.Join(s.root, urlPath)
		ext := ""
		extPos := strings.LastIndexByte(urlPath, '.')
		if extPos != -1 {
			ext = urlPath[extPos+1:]
		}
		var (
			buffer   []byte
			have     bool
			mimeType string
		)
		if buffer, have = s.caches[urlPath]; !have {
			bytes, err := ioutil.ReadFile(realPath)
			if err != nil {
				ctx.Chain()
				return
			} else {
				buffer = bytes
				if mm, extHave := s.mimes[ext]; !extHave {
					mimeType = mime.BINARY
				} else {
					mimeType = mm
				}
				if s.cache {
					s.caches[urlPath] = buffer
					s.mimeCaches[urlPath] = mimeType
				}
			}
		} else {
			mimeType = s.mimeCaches[urlPath]
		}
		ctx.Render(context.RenderBuilder().Buffer(buffer).ContentType(mimeType).Build())
	}
}

// Add mime
func (s *static) Add(ext, mime string) *static {
	s.mimes[ext] = mime
	return s
}

// Adds mimes
func (s *static) Adds(m map[string]string) *static {
	if m != nil {
		for k, v := range m {
			s.Add(k, v)
		}
	}
	return s
}
