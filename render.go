package flygo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

//Response text to client
func (c *Context) Text(text string) {
	c.Render([]byte(text), contentTypeText)
}

//Response text file to client
func (c *Context) TextFile(textFile string) {
	bytes, err := readRealPath(textFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.Text(string(bytes))
}

//Response json to client
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		app.Error(err.Error())
		return
	}
	c.Render(jsonData, contentTypeJson)
}

//Response json file to client
func (c *Context) JSONFile(jsonFile string) {
	bytes, err := readRealPath(jsonFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeJson)
}

//Response xml to client
func (c *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		app.Error(err.Error())
		return
	}
	c.Render(xmlData, contentTypeXml)
}

//Response xml file to client
func (c *Context) XMLFile(xmlFile string) {
	bytes, err := readRealPath(xmlFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeXml)
}

//Response image to client
func (c *Context) Image(buffer []byte) {
	c.Render(buffer, contentTypeImage)
}

//Response image file to client
func (c *Context) ImageFile(imageFile string) {
	bytes, err := readRealPath(imageFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.Image(bytes)
}

//Response bmp image to client
func (c *Context) BMP(buffer []byte) {
	c.Render(buffer, contentTypeBmp)
}

//Response bmp image file to client
func (c *Context) BMPFile(bmpFile string) {
	bytes, err := readRealPath(bmpFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.BMP(bytes)
}

//Response jpg image to client
func (c *Context) JPG(buffer []byte) {
	c.Render(buffer, contentTypeJpg)
}

//Response jpg image file to client
func (c *Context) JPGFile(jpgFile string) {
	bytes, err := readRealPath(jpgFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.JPG(bytes)
}

//Response jpeg image to client
func (c *Context) JPEG(buffer []byte) {
	c.Render(buffer, contentTypeJpg)
}

//Response jpeg image file to client
func (c *Context) JPEGFile(jpegFile string) {
	c.JPGFile(jpegFile)
}

//Response png image to client
func (c *Context) PNG(buffer []byte) {
	c.Render(buffer, contentTypePng)
}

//Response png image file to client
func (c *Context) PNGFile(pngFile string) {
	bytes, err := readRealPath(pngFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.PNG(bytes)
}

//Response gif image to client
func (c *Context) GIF(buffer []byte) {
	c.Render(buffer, contentTypeGif)
}

//Response gif image file to client
func (c *Context) GIFFile(gifFile string) {
	bytes, err := readRealPath(gifFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.GIF(bytes)
}

//Response html to client
func (c *Context) HTML(buffer []byte) {
	c.Render(buffer, contentTypeHtml)
}

//Response html file to client
func (c *Context) HTMLFile(htmlFile string) {
	bytes, err := readRealPath(htmlFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.HTML(bytes)
}

//Response css to client
func (c *Context) CSS(buffer []byte) {
	c.Render(buffer, contentTypeCSS)
}

//Response css file to client
func (c *Context) CSSFile(cssFile string) {
	bytes, err := readRealPath(cssFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.CSS(bytes)
}

//Response js to client
func (c *Context) JS(buffer []byte) {
	c.Render(buffer, contentTypeJS)
}

//Response js file to client
func (c *Context) JSFile(jsFile string) {
	bytes, err := readRealPath(jsFile)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.JS(bytes)
}

//Response file to client
func (c *Context) Binary(buffer []byte) {
	c.Render(buffer, contentTypeBinary)
}

//Response file to client
func (c *Context) File(file string) {
	bytes, err := readRealPath(file)
	if err != nil {
		app.Error("%v", err)
		return
	}
	c.Binary(bytes)
}

//Response download file to client
func (c *Context) Download(file, fileName string) {
	bytes, err := readRealPath(file)
	if err != nil {
		app.Error("%v", err)
		return
	}
	fn := url.PathEscape(fileName)
	c.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fn))
	c.Binary(bytes)
}

//Base Response
func (c *Context) Render(buffer []byte, contentType string) {
	if !c.Response.done {
		c.Response.data = buffer
		c.Response.contentType = contentType
		c.Response.done = true
	}
}

func getRealPath(fileName string) string {
	return filepath.Join(app.Config.Flygo.Server.WebRoot, fileName)
}

func readRealPath(fileName string) ([]byte, error) {
	rp := getRealPath(fileName)
	return ioutil.ReadFile(rp)
}
