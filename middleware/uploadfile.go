package middleware

import (
	"fmt"
	"github.com/billcoding/flygo/context"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type uploadFile struct {
	logger     *log.Logger
	root       string
	size       int
	mimes      []string
	extensions []string
	domain     string
	prefix     string
	dateDir    bool
}

// UploadFile return new uploadFile
func UploadFile() *uploadFile {
	return &uploadFile{
		logger:     log.New(os.Stdout, "[uploadFile]", log.LstdFlags),
		root:       os.TempDir(),
		size:       100 * 1024 * 1024,
		mimes:      []string{"text/plain", "image/jpeg", "image/jpg", "image/png", "image/gif", "application/octet-stream"},
		extensions: []string{".txt", ".jpg", ".png", ".gif", ".xlsx"},
		domain:     "http://localhost/",
		prefix:     "/download/file",
		dateDir:    true,
	}
}

// Type implements
func (uf *uploadFile) Type() *Type {
	return TypeHandler
}

// Name implements
func (*uploadFile) Name() string {
	return "UploadFile"
}

// Method implements
func (uf *uploadFile) Method() Method {
	return MethodPost
}

// Pattern implements
func (uf *uploadFile) Pattern() Pattern {
	return "/upload/file"
}

// Handler implements
func (uf *uploadFile) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		defer removeTmpFiles(ctx)
		type jd struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		}
		type jdd struct {
			Msg  string                 `json:"msg"`
			Code int                    `json:"code"`
			Data map[string]interface{} `json:"data"`
		}
		getJson := func(msg string, code int) *jd {
			return &jd{
				Msg:  msg,
				Code: code,
			}
		}
		getJsonData := func(data map[string]interface{}) *jdd {
			return &jdd{
				Msg:  "success",
				Code: 0,
				Data: data,
			}
		}

		err := ctx.ParseMultipart(int64(uf.size))
		if err != nil {
			uf.logger.Println(err)
			ctx.JSON(getJson("parse file error", 1))
			return
		}

		files := ctx.MultipartFiles("file")
		for _, file := range files {
			err := verifyFile(file, uf)
			if err != nil {
				uf.logger.Println(err)
				ctx.JSON(getJson(err.Error(), 1))
				return
			}
		}
		uploadFiles := make([]UFile, 0)
		saveFiles := make([]string, 0)
		dateDir := ""
		if uf.dateDir {
			dateDir = time.Now().Format("20060102")
		}
		parentPath := filepath.Join(uf.root, dateDir)
		_, err = ioutil.ReadDir(parentPath)
		if err != nil {
			err := os.MkdirAll(parentPath, os.ModePerm)
			if err != nil {
				uf.logger.Println(err)
				ctx.JSON(getJson(err.Error(), 1))
				return
			}
		}

		for _, file := range files {
			extIndex := strings.LastIndexByte(file.Filename, '.')
			ext := file.Filename[extIndex+1:]

			saveFile := fmt.Sprintf("%x.%s", rand.Int63(), ext)
			saveFilePath := filepath.Join(parentPath, saveFile)

			fileUrl := fmt.Sprintf("%s%s?file=%s", uf.domain, uf.prefix, saveFile)
			uploadFiles = append(uploadFiles, UFile{
				File: saveFile,
				Url:  fileUrl,
			})

			err := file.Copy(saveFilePath)
			if err != nil {
				uf.logger.Println(err)
				_ = os.Chmod(saveFilePath, os.ModePerm)
				removeSaveFiles(saveFiles)
				ctx.JSON(getJson(err.Error(), 1))
				return
			}

			saveFiles = append(saveFiles, saveFilePath)
		}

		ctx.JSON(getJsonData(map[string]interface{}{
			"total": len(files),
			"files": uploadFiles,
		}))

	}
}

// Root set
func (uf *uploadFile) Root(root string) *uploadFile {
	uf.root = root
	return uf
}

// Size set
func (uf *uploadFile) Size(size int) *uploadFile {
	uf.size = size
	return uf
}

// Extensions set
func (uf *uploadFile) Extensions(extensions []string) *uploadFile {
	uf.extensions = extensions
	return uf
}

// AddExtension add extensions
func (uf *uploadFile) AddExtension(extensions ...string) *uploadFile {
	uf.extensions = append(uf.extensions, extensions...)
	return uf
}

// Mimes set
func (uf *uploadFile) Mimes(mimes []string) *uploadFile {
	uf.mimes = mimes
	return uf
}

// AddMime add mimes
func (uf *uploadFile) AddMime(mimes ...string) *uploadFile {
	uf.mimes = append(uf.mimes, mimes...)
	return uf
}

// Domain set
func (uf *uploadFile) Domain(domain string) *uploadFile {
	uf.domain = domain
	return uf
}

// Prefix set
func (uf *uploadFile) Prefix(prefix string) *uploadFile {
	uf.prefix = prefix
	return uf
}

// DateDir set
func (uf *uploadFile) DateDir(dateDir bool) *uploadFile {
	uf.dateDir = dateDir
	return uf
}

// UFile struct
type UFile struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func verifyFile(file *context.MultipartFile, uf *uploadFile) error {

	err := verifySize(file, uf)
	if err != nil {
		return err
	}

	err = verifyMime(file, uf)
	if err != nil {
		return err
	}

	err = verifyExt(file, uf)
	if err != nil {
		return err
	}

	return nil
}

func verifySize(file *context.MultipartFile, uf *uploadFile) error {
	if file.Size > int64(uf.size) {
		return fmt.Errorf("the file size exceed limit[max:%d, current:%d]", uf.size, file.Size)
	}
	return nil
}

func verifyMime(file *context.MultipartFile, uf *uploadFile) error {
	ms := strings.Join(uf.mimes, "|")
	mimeAll := fmt.Sprintf("|%s|", ms)
	if !strings.Contains(mimeAll, fmt.Sprintf("|%s|", file.ContentType)) {
		return fmt.Errorf("the file mime type not supported[supports:%s, current:%s]", ms, file.ContentType)
	}
	return nil
}

func verifyExt(file *context.MultipartFile, uf *uploadFile) error {
	es := strings.Join(uf.extensions, "|")
	extAll := fmt.Sprintf("|%s|", es)
	extIndex := strings.LastIndexByte(file.Filename, '.')
	if extIndex == -1 {
		return fmt.Errorf("the file ext not supported[supports:%s, current:%s]", es, "")
	}
	ext := file.Filename[extIndex:]
	if !strings.Contains(extAll, fmt.Sprintf("|%s|", ext)) {
		return fmt.Errorf("the file ext not supported[supports:%s, current:%s]", es, ext)
	}
	return nil
}

func removeTmpFiles(ctx *context.Context) {
	defer func() {
		if re := recover(); re != nil {
		}
	}()
	form := ctx.Request.MultipartForm
	if form != nil {
		_ = form.RemoveAll()
	}
}

func removeSaveFiles(saveFiles []string) {
	for _, f := range saveFiles {
		_ = os.Remove(f)
	}
}
