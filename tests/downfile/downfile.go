package main

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	mw "github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

//Define DownFile struct
type DownFile struct {
	logger *log.Logger
	//The root dir
	root string
	//enable date dir
	dateDir bool
}

//Return new DownFile
func New() *DownFile {
	return &DownFile{
		logger:  log.New(os.Stdout, "[DownFile]", log.LstdFlags),
		root:    os.TempDir(),
		dateDir: true,
	}
}

//Name
func (*DownFile) Name() string {
	return "DownFile"
}

//Type
func (df *DownFile) Type() *mw.Type {
	return mw.TypeHandler
}

//Method
func (df *DownFile) Method() mw.Method {
	return mw.MethodGet
}

//Pattern
func (df *DownFile) Pattern() mw.Pattern {
	return "/download/downfile"
}

//Handler
func (df *DownFile) Handler() func(c *c.Context) {
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
func (df *DownFile) Root(root string) *DownFile {
	df.root = root
	return df
}

//DateDir
func (df *DownFile) DateDir(dateDir bool) *DownFile {
	df.dateDir = dateDir
	return df
}
