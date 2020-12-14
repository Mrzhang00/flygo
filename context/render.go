package context

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Define render struct
type Render struct {
	rended      bool
	Buffer      []byte
	Header      http.Header
	Cookies     []*http.Cookie
	Code        int
	ContentType string
}

//Render
func (c *Context) Render() *Render {
	return c.render
}

//Rended
func (r *Render) Rended() bool {
	return r.rended
}

//Text
func (c *Context) Text(text string) {
	c.render = RenderBuilder().Buffer([]byte(text)).ContentType(mime.TEXT).Build()
}

//TextFile
func (c *Context) TextFile(textFile string) {
	bytes, err := readFile(textFile)
	if err != nil {
		c.logger.Error("[TextFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer([]byte(string(bytes))).ContentType(mime.TEXT).Build()
}

//JSON
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.logger.Error("[JSON]%v", err.Error())
		return
	}
	c.render = RenderBuilder().Buffer(jsonData).ContentType(mime.JSON).Build()
}

//JSONText
func (c *Context) JSONText(json string) {
	c.render = RenderBuilder().Buffer([]byte(json)).ContentType(mime.JSON).Build()
}

//JSONFile
func (c *Context) JSONFile(jsonFile string) {
	bytes, err := readFile(jsonFile)
	if err != nil {
		c.logger.Error("[JSONFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.JSON).Build()
}

//XML
func (c *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		c.logger.Error("[XML]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(xmlData).ContentType(mime.XML).Build()
}

//XMLText
func (c *Context) XMLText(xml string) {
	c.render = RenderBuilder().Buffer([]byte(xml)).ContentType(mime.XML).Build()
}

//XMLFile
func (c *Context) XMLFile(xmlFile string) {
	bytes, err := readFile(xmlFile)
	if err != nil {
		c.logger.Error("[XMLFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.XML).Build()
}

//Image
func (c *Context) Image(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build()
}

//ImageFile
func (c *Context) ImageFile(imageFile string) {
	bytes, err := readFile(imageFile)
	if err != nil {
		c.logger.Error("[ImageFile]%v", err)
		return
	}
	c.Image(bytes)
}

//ICO
func (c *Context) ICO(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.ICO).Build()
}

//ICOFile
func (c *Context) ICOFile(icoFile string) {
	bytes, err := readFile(icoFile)
	if err != nil {
		c.logger.Error("[ICOFile]%v", err)
		return
	}
	c.ICO(bytes)
}

//BMP
func (c *Context) BMP(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.BMP).Build()
}

//BMPFile
func (c *Context) BMPFile(bmpFile string) {
	bytes, err := readFile(bmpFile)
	if err != nil {
		c.logger.Error("[BMPFile]%v", err)
		return
	}
	c.BMP(bytes)
}

//JPG
func (c *Context) JPG(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build()
}

//JPGFile
func (c *Context) JPGFile(jpgFile string) {
	bytes, err := readFile(jpgFile)
	if err != nil {
		c.logger.Error("[JPGFile]%v", err)
		return
	}
	c.JPG(bytes)
}

//JPEG
func (c *Context) JPEG(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build()
}

//JPEGFile
func (c *Context) JPEGFile(jpegFile string) {
	c.JPGFile(jpegFile)
}

//PNG
func (c *Context) PNG(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.PNG).Build()
}

//PNGFile
func (c *Context) PNGFile(pngFile string) {
	bytes, err := readFile(pngFile)
	if err != nil {
		c.logger.Error("[PNGFile]%v", err)
		return
	}
	c.PNG(bytes)
}

//GIF
func (c *Context) GIF(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.GIF).Build()
}

//GIFFile
func (c *Context) GIFFile(gifFile string) {
	bytes, err := readFile(gifFile)
	if err != nil {
		c.logger.Error("[GIFFile]%v", err)
		return
	}
	c.GIF(bytes)
}

//HTML
func (c *Context) HTML(html string) {
	c.render = RenderBuilder().Buffer([]byte(html)).ContentType(mime.HTML).Build()
}

//HTMLFile
func (c *Context) HTMLFile(htmlFile string) {
	bytes, err := readFile(htmlFile)
	if err != nil {
		c.logger.Error("[HTMLFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.HTML).Build()
}

//CSS
func (c *Context) CSS(css string) {
	c.render = RenderBuilder().Buffer([]byte(css)).ContentType(mime.CSS).Build()
}

//CSSFile
func (c *Context) CSSFile(cssFile string) {
	bytes, err := readFile(cssFile)
	if err != nil {
		c.logger.Error("[CSSFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.CSS).Build()
}

//JS
func (c *Context) JS(js string) {
	c.render = RenderBuilder().Buffer([]byte(js)).ContentType(mime.JS).Build()
}

//JSFile
func (c *Context) JSFile(jsFile string) {
	bytes, err := readFile(jsFile)
	if err != nil {
		c.logger.Error("[JSFile]%v", err)
		return
	}
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.JS).Build()
}

//Binary
func (c *Context) Binary(buffer []byte) {
	c.render = RenderBuilder().Buffer(buffer).ContentType(mime.BINARY).Build()
}

//File
func (c *Context) File(file string) {
	bytes, err := readFile(file)
	if err != nil {
		c.logger.Error("[File]%v", err)
		return
	}
	c.Binary(bytes)
}

//Download
func (c *Context) Download(file, fileName string) {
	bytes, err := readFile(file)
	if err != nil {
		c.logger.Error("[Download]%v", err)
		return
	}
	fn := url.PathEscape(fileName)
	c.render = RenderBuilder().Buffer(bytes).ContentType(mime.BINARY).Header(http.Header{
		headers.ContentDisposition: []string{fmt.Sprintf("attachment; filename=\"%s\"", fn)},
	}).Build()
}

//readFile
func readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
