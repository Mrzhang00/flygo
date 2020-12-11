package context

import (
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

//Define MultipartFile struct
type MultipartFile struct {
	Logger      log.Logger
	Filename    string
	ContentType string
	Size        int64
	FileHeader  *multipart.FileHeader
	Headers     http.Header
}

//Copy
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

//ParseMultipart
func (c *Context) ParseMultipart(maxMemory int64) {
	var err error
	err = c.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		c.logger.Error("[ParseMultipart]%v", err)
		return
	}
	paramMap := make(map[string][]string, 0)
	for name, values := range c.Request.MultipartForm.Value {
		paramMap[name] = values
	}
	c.SetParamMap(paramMap)
	for name, header := range c.Request.MultipartForm.File {
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
		c.MultipartMap[name] = mfs
	}
}

//MultipartFile
func (c *Context) MultipartFile(name string) *MultipartFile {
	files := c.MultipartFiles(name)
	if files != nil && len(files) > 0 {
		return files[0]
	}
	return nil
}

//MultipartFiles
func (c *Context) MultipartFiles(name string) []*MultipartFile {
	files, have := c.MultipartMap[name]
	if have {
		return files
	}
	return nil
}
