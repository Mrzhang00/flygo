package context

import (
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// MultipartFile struct
type MultipartFile struct {
	// Logger MultipartFile logger
	Logger log.Logger
	// Filename file name
	Filename string
	// ContentType Content type
	ContentType string
	// Size file's total size
	Size int64
	// FileHeader file headers
	FileHeader *multipart.FileHeader
	// Headers for Request
	Headers http.Header
}

// Copy source file in multiple
func (file *MultipartFile) Copy(distName string) error {
	var f multipart.File
	var dist *os.File
	var err error
	f, err = file.FileHeader.Open()
	if err != nil {
		file.Logger.Error("%v", err)
		return err
	}

	dist, err = os.Create(distName)
	if err != nil {
		file.Logger.Error("%v", err)
		return err
	}

	_, err = io.Copy(dist, f)
	if err != nil {
		file.Logger.Error("%v", err)
		return err
	}
	_ = f.Close()
	_ = dist.Close()
	return nil
}

// ParseMultipart parse multiple Request
func (ctx *Context) ParseMultipart(maxMemory int64) error {
	var err error
	err = ctx.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		ctx.logger.Error("[ParseMultipart]%v", err)
		return err
	}
	paramMap := make(map[string][]string, 0)
	for name, values := range ctx.Request.MultipartForm.Value {
		paramMap[name] = values
	}
	ctx.SetParamMap(paramMap)
	for name, header := range ctx.Request.MultipartForm.File {
		mfs := make([]*MultipartFile, 0)
		for _, fileHeader := range header {
			mf := &MultipartFile{
				Logger:      log.New("[MultipartFile]"),
				Filename:    fileHeader.Filename,
				ContentType: fileHeader.Header.Get(headers.MIME),
				Size:        fileHeader.Size,
				FileHeader:  fileHeader,
				Headers:     make(map[string][]string, 0),
			}
			for k, v := range fileHeader.Header {
				mf.Headers[k] = v
			}
			mfs = append(mfs, mf)
		}
		ctx.MultipartMap[name] = mfs
	}
	return nil
}

// MultipartFile get multiple file
func (ctx *Context) MultipartFile(name string) *MultipartFile {
	files := ctx.MultipartFiles(name)
	if files != nil && len(files) > 0 {
		return files[0]
	}
	return nil
}

// MultipartFiles get multiple files
func (ctx *Context) MultipartFiles(name string) []*MultipartFile {
	files, have := ctx.MultipartMap[name]
	if have {
		return files
	}
	return nil
}
