package middleware

import (
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type static struct {
	cache  bool
	caches map[string][]byte
	root   string
	mimes  map[string]string
	logger log.Logger
}

func Static(cache bool, root string) *static {
	rot := root

	calls.True(rot == "" || rot == "." || rot == "./", func() {
		execdir, _ := os.Executable()
		rot = filepath.Dir(execdir)
	})
	calls.NEmpty(strings.TrimPrefix(rot, "./"), func() {
		execdir, _ := os.Executable()
		rot = filepath.Join(filepath.Dir(execdir), strings.TrimPrefix(rot, "./"))
	})
	return &static{
		cache:  cache,
		caches: make(map[string][]byte, 0),
		root:   rot,
		mimes:  defaultMimes(),
		logger: log.New("[Static]"),
	}
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

func (s *static) Type() *Type {
	return TypeBefore
}

func (s *static) Name() string {
	return "Static"
}

func (s *static) Method() Method {
	return MethodGet
}

func (s *static) Pattern() Pattern {
	return "/static/*"
}

func (s *static) Handler() func(c *c.Context) {
	return func(ctx *c.Context) {
		if strings.HasSuffix(ctx.Request.URL.Path, "/") {
			ctx.Chain()
			return
		}

		vpath := strings.TrimPrefix(ctx.Request.URL.Path, "/static")
		rpath := filepath.Join(s.root, vpath)
		ext := ""
		extPos := strings.LastIndexByte(vpath, '.')

		calls.True(extPos != -1, func() {
			ext = vpath[extPos+1:]
		})

		buffer, have := s.caches[vpath]
		calls.False(have, func() {
			bytes, err := ioutil.ReadFile(rpath)
			calls.NNil(err, func() {
				s.logger.Warn("[Handler]%v", err)
				ctx.Chain()
			})
			calls.Nil(err, func() {
				buffer = bytes
				calls.True(s.cache, func() {
					s.caches[vpath] = buffer
				})
				mm, have := s.mimes[ext]
				calls.False(have, func() {
					mm = mime.BINARY
				})
				ctx.Rende(c.RenderBuilder().Buffer(buffer).ContentType(mm).Build())
			})
		})

		calls.True(have, func() {
			mm, have := s.mimes[ext]
			calls.False(have, func() {
				mm = mime.BINARY
			})
			ctx.Rende(c.RenderBuilder().Buffer(buffer).ContentType(mm).Build())
		})
	}
}

func (s *static) Add(ext, mime string) *static {
	s.mimes[ext] = mime
	return s
}

func (s *static) Adds(m map[string]string) *static {
	calls.NNil(m, func() {
		for k, v := range m {
			s.Add(k, v)
		}
	})
	return s
}
