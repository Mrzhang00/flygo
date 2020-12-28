package middleware

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

//Define downFile struct
type downFile struct {
	logger *log.Logger
	//The root dir
	root string
	//enable date dir
	dateDir bool
}

//downFile
func DownFile() *downFile {
	return &downFile{
		logger:  log.New(os.Stdout, "[downFile]", log.LstdFlags),
		root:    os.TempDir(),
		dateDir: true,
	}
}

//Name
func (df *downFile) Name() string {
	return "DownFile"
}

//Type
func (df *downFile) Type() *Type {
	return TypeHandler
}

//Method
func (df *downFile) Method() Method {
	return MethodGet
}

//Pattern
func (df *downFile) Pattern() Pattern {
	return "/download/downfile"
}

//Handler
func (df *downFile) Handler() func(c *c.Context) {
	return func(ctx *c.Context) {
		type req struct {
			File string `binding:"name(file)" validate:"required(T) minlength(1) message(file is empty)"`
		}
		fs := req{}
		ctx.BindAndValidate(&fs, func() {
			root := df.root
			dateDir := ""
			if df.dateDir {
				dateDir = time.Now().Format("20060102")
			}
			filePath := filepath.Join(root, dateDir, fs.File)
			bytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				df.logger.Println(err)
				ctx.JSON(map[string]interface{}{"code": 1, "msg": "file is not exists"})
				return
			}
			ctx.Rende(c.RenderBuilder().Header(map[string][]string{
				headers.ContentDisposition: {"attachement;filename=" + url.QueryEscape(fs.File)},
			}).ContentType(mime.BINARY).Buffer(bytes).Build())
		})
	}
}

//Root
func (df *downFile) Root(root string) *downFile {
	df.root = root
	return df
}

//DateDir
func (df *downFile) DateDir(dateDir bool) *downFile {
	df.dateDir = dateDir
	return df
}
