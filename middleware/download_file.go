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

type DownloadFile struct {
	root    string
	dateDir bool
}

func NewDownloadFile() *DownloadFile {
	return &DownloadFile{root: os.TempDir(), dateDir: true}
}

// Name implements
func (df *DownloadFile) Name() string {
	return "DownloadFile"
}

// Type implements
func (df *DownloadFile) Type() *Type {
	return TypeHandler
}

// Method implements
func (df *DownloadFile) Method() Method {
	return MethodGet
}

// Pattern implements
func (df *DownloadFile) Pattern() Pattern {
	return "/download/file"
}

// Handler implements
func (df *DownloadFile) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		fileName := ctx.Param("file")
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
func (df *DownloadFile) Root(root string) *DownloadFile {
	df.root = root
	return df
}

func (df *DownloadFile) DateDir(dateDir bool) *DownloadFile {
	df.dateDir = dateDir
	return df
}

func (df *DownloadFile) Path(file string) string {
	dd := ""
	if df.dateDir {
		dd = time.Now().Format("20060102")
	}
	return filepath.Join(df.root, dd, file)
}

func (df *DownloadFile) File(file string) (*os.File, error) {
	return os.Open(df.Path(file))
}

func (df *DownloadFile) Buf(file string) ([]byte, error) {
	return ioutil.ReadFile(df.Path(file))
}
