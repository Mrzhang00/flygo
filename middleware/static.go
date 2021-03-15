package middleware

import (
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/billcoding/flygo/mime"
	"github.com/billcoding/flygo/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type static struct {
	cache   bool
	caches  map[string][]byte
	root    string
	mimes   map[string]string
	logger  log.Logger
	handler func(ctx *context.Context)
}

// Static new static
func Static(cache bool, root string, handlers ...func(ctx *context.Context)) *static {
	rot := root
	if rot == "" || rot == "." || rot == "./" {
		executeDir, _ := os.Executable()
		rot = filepath.Dir(executeDir)
	}
	if strings.TrimPrefix(rot, "./") != "" {
		executeDir, _ := os.Executable()
		rot = filepath.Join(filepath.Dir(executeDir), strings.TrimPrefix(rot, "./"))
	}
	st := &static{
		cache:  cache,
		caches: make(map[string][]byte, 0),
		root:   rot,
		mimes:  defaultMimes(),
		logger: log.New("[Static]"),
	}
	if len(handlers) > 0 && handlers[0] != nil {
		st.handler = handlers[0]
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
	return TypeBefore
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

// Type implements
func (s *static) Handler() func(ctx *context.Context) {
	if s.handler != nil {
		return s.handler
	}
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
		buffer, have := s.caches[urlPath]
		if !have {
			mm, extHave := s.mimes[ext]
			if !extHave {
				mm = mime.BINARY
			}
			ctx.Render(context.RenderBuilder().Buffer(buffer).ContentType(mm).Build())
		} else {
			bytes, err := ioutil.ReadFile(realPath)
			if err != nil {
				s.logger.Warn("[Handler]%v", err)
				ctx.Chain()
			} else {
				buffer = bytes
				if s.cache {
					s.caches[urlPath] = buffer
				}
				mm, extHave := s.mimes[ext]
				if !extHave {
					mm = mime.BINARY
				}
				ctx.Render(context.RenderBuilder().Buffer(buffer).ContentType(mm).Build())
			}
		}
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
