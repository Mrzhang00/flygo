package middleware

import (
	"fmt"
	"github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type downFile struct {
	root    string
	dateDir bool
}

// DownFile return new downFile
func DownFile() *downFile {
	return &downFile{
		root:    os.TempDir(),
		dateDir: true,
	}
}

// Name implements
func (df *downFile) Name() string {
	return "DownFile"
}

// Type implements
func (df *downFile) Type() *Type {
	return TypeHandler
}

// Method implements
func (df *downFile) Method() Method {
	return MethodGet
}

// Pattern implements
func (df *downFile) Pattern() Pattern {
	return "/download/file"
}

// Handler implements
func (df *downFile) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		fileName := ctx.ParamWith("file", "")
		if fileName == "" {
			ctx.JSONText(`{"msg":"file is empty","code":1}`)
			return
		}
		root := df.root
		dateDir := ""
		if df.dateDir {
			dateDir = time.Now().Format("20060102")
		}
		filePath := filepath.Join(root, dateDir, fileName)
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			ctx.JSON(map[string]interface{}{"code": 1, "msg": "file is not exists"})
			return
		}
		ctx.Render(context.RenderBuilder().Header(map[string][]string{
			headers.ContentDisposition: {"attachment;filename=" + url.QueryEscape(fileName)},
		}).ContentType(mime.BINARY).Buffer(bytes).Build())
	}
}

// Root path
func (df *downFile) Root(root string) *downFile {
	df.root = root
	return df
}

// DateDir enable
func (df *downFile) DateDir(dateDir bool) *downFile {
	df.dateDir = dateDir
	return df
}
