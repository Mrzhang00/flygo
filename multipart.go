package flygo

import (
	"io"
	"mime/multipart"
	"os"
)

//Define MultipartFile struct
type MultipartFile struct {
	logger      *log                  //bundle logger
	filename    string                //original file name
	contentType string                //mime type
	size        int64                 //file length in byte
	fileHeader  *multipart.FileHeader //file ptr
	headers     headerMap             //file headers
}

//Copy file to dist name
func (file *MultipartFile) Copy(distName string) error {
	var f multipart.File
	var dist *os.File
	var err error
	f, err = file.fileHeader.Open()
	defer f.Close()
	if err != nil {
		file.logger.Error("%v", err)
		return err
	}
	dist, err = os.Create(distName)
	defer dist.Close()
	if err != nil {
		file.logger.Error("%v", err)
		return err
	}
	_, err = io.Copy(dist, f)
	if err != nil {
		file.logger.Error("%v", err)
		return err
	}
	return nil
}

//Get multipart file content type
func (file *MultipartFile) ContentType() string {
	return file.contentType
}

//Get multipart file size in byte
func (file *MultipartFile) Size() int64 {
	return file.size
}

//Get multipart file name
func (file *MultipartFile) Filename() string {
	return file.filename
}

//Get multipart header
func (file *MultipartFile) Header(name string) string {
	v := file.headers[name]
	if v != nil {
		return v[0]
	}
	return ""
}

//Get multipart headers
func (file *MultipartFile) Headers() map[string][]string {
	return file.headers
}

//Parse Multipart
func (c *Context) ParseMultipart(maxMemory int64) error {
	var err error
	err = c.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		return err
	}
	c.MultipartParsed = true
	for name, values := range c.Request.MultipartForm.Value {
		c.ParamMap[name] = values
	}
	for name, header := range c.Request.MultipartForm.File {
		mfs := make([]*MultipartFile, 0)
		for _, fileHeader := range header {
			mf := &MultipartFile{
				logger:      c.app.Logger,
				filename:    fileHeader.Filename,
				contentType: fileHeader.Header.Get(headerContentType),
				size:        fileHeader.Size,
				fileHeader:  fileHeader,
				headers:     make(map[string][]string, 0),
			}
			for k, v := range fileHeader.Header {
				mf.headers[k] = v
			}
			mfs = append(mfs, mf)
		}
		c.Multipart[name] = mfs
	}
	return err
}

//Get multipart file
func (c *Context) MultipartFile(name string) *MultipartFile {
	files := c.MultipartFiles(name)
	if len(files) <= 0 {
		return nil
	}
	return files[0]
}

//Get request parameter
func (c *Context) MultipartFiles(name string) []*MultipartFile {
	if !c.MultipartParsed {
		c.app.Logger.Warn("Multipart is not parsed")
		return nil
	}
	return c.Multipart[name]
}
