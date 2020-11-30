package flygo

import (
	"fmt"
	"testing"
)

//Test multipart
func TestMultipart(t *testing.T) {

	app := NewApp()

	app.Post("/upload", func(c *Context) {
		//First parse
		c.ParseMultipart(1024 * 1024)

		//Get file
		file := c.MultipartFile("file")

		if file != nil {
			filename := file.Filename()
			size := file.Size()
			contentType := file.ContentType()
			file.Copy("D:\\tmpfile")
			fmt.Printf("filename = %v,size = %v,contentType = %v\n", filename, size, contentType)
			c.Text("success")
		}

		c.Text("fail")
	})

	app.Run()

}
