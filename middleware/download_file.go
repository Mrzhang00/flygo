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

type downloadFile struct {
	root    string
	dateDir bool
}

// DownloadFile return new downloadFile
func DownloadFile() *downloadFile {
	return &downloadFile{
		root:    os.TempDir(),
		dateDir: true,
	}
}

// Name implements
func (df *downloadFile) Name() string {
	return "DownloadFile"
}

// Type implements
func (df *downloadFile) Type() *Type {
	return TypeHandler
}

// Method implements
func (df *downloadFile) Method() Method {
	return MethodGet
}

// Pattern implements
func (df *downloadFile) Pattern() Pattern {
	return "/download/file"
}

// Handler implements
func (df *downloadFile) Handler() func(ctx *context.Context) {
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
func (df *downloadFile) Root(root string) *downloadFile {
	df.root = root
	return df
}

// DateDir enable
func (df *downloadFile) DateDir(dateDir bool) *downloadFile {
	df.dateDir = dateDir
	return df
}
