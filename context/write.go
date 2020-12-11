package context

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"io/ioutil"
	"net/url"
)

//Text
func (c *Context) Text(text string) {
	c.Render([]byte(text), mime.TEXT)
}

//TextFile
func (c *Context) TextFile(textFile string) {
	bytes, err := c.readFile(textFile)
	if err != nil {
		c.logger.Error("[TextFile]%v", err)
		return
	}
	c.Text(string(bytes))
}

//JSON
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.logger.Error("[JSON]%v", err.Error())
		return
	}
	c.Render(jsonData, mime.JSON)
}

//JSONText
func (c *Context) JSONText(json string) {
	c.Render([]byte(json), mime.JSON)
}

//JSONFile
func (c *Context) JSONFile(jsonFile string) {
	bytes, err := c.readFile(jsonFile)
	if err != nil {
		c.logger.Error("[JSONFile]%v", err)
		return
	}
	c.Render(bytes, mime.JSON)
}

//XML
func (c *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		c.logger.Error("[XML]%v", err)
		return
	}
	c.Render(xmlData, mime.XML)
}

//XMLText
func (c *Context) XMLText(xml string) {
	c.Render([]byte(xml), mime.XML)
}

//XMLFile
func (c *Context) XMLFile(xmlFile string) {
	bytes, err := c.readFile(xmlFile)
	if err != nil {
		c.logger.Error("[XMLFile]%v", err)
		return
	}
	c.Render(bytes, mime.XML)
}

//Image
func (c *Context) Image(buffer []byte) {
	c.Render(buffer, mime.JPG)
}

//ImageFile
func (c *Context) ImageFile(imageFile string) {
	bytes, err := c.readFile(imageFile)
	if err != nil {
		c.logger.Error("[ImageFile]%v", err)
		return
	}
	c.Image(bytes)
}

//ICO
func (c *Context) ICO(buffer []byte) {
	c.Render(buffer, mime.ICO)
}

//ICOFile
func (c *Context) ICOFile(icoFile string) {
	bytes, err := c.readFile(icoFile)
	if err != nil {
		c.logger.Error("[ICOFile]%v", err)
		return
	}
	c.ICO(bytes)
}

//BMP
func (c *Context) BMP(buffer []byte) {
	c.Render(buffer, mime.BMP)
}

//BMPFile
func (c *Context) BMPFile(bmpFile string) {
	bytes, err := c.readFile(bmpFile)
	if err != nil {
		c.logger.Error("[BMPFile]%v", err)
		return
	}
	c.BMP(bytes)
}

//JPG
func (c *Context) JPG(buffer []byte) {
	c.Render(buffer, mime.JPG)
}

//JPGFile
func (c *Context) JPGFile(jpgFile string) {
	bytes, err := c.readFile(jpgFile)
	if err != nil {
		c.logger.Error("[JPGFile]%v", err)
		return
	}
	c.JPG(bytes)
}

//JPEG
func (c *Context) JPEG(buffer []byte) {
	c.Render(buffer, mime.JPG)
}

//JPEGFile
func (c *Context) JPEGFile(jpegFile string) {
	c.JPGFile(jpegFile)
}

//PNG
func (c *Context) PNG(buffer []byte) {
	c.Render(buffer, mime.PNG)
}

//PNGFile
func (c *Context) PNGFile(pngFile string) {
	bytes, err := c.readFile(pngFile)
	if err != nil {
		c.logger.Error("[PNGFile]%v", err)
		return
	}
	c.PNG(bytes)
}

//GIF
func (c *Context) GIF(buffer []byte) {
	c.Render(buffer, mime.GIF)
}

//GIFFile
func (c *Context) GIFFile(gifFile string) {
	bytes, err := c.readFile(gifFile)
	if err != nil {
		c.logger.Error("[GIFFile]%v", err)
		return
	}
	c.GIF(bytes)
}

//HTML
func (c *Context) HTML(html string) {
	c.Render([]byte(html), mime.HTML)
}

//HTMLFile
func (c *Context) HTMLFile(htmlFile string) {
	bytes, err := c.readFile(htmlFile)
	if err != nil {
		c.logger.Error("[HTMLFile]%v", err)
		return
	}
	c.Render(bytes, mime.HTML)
}

//CSS
func (c *Context) CSS(css string) {
	c.Render([]byte(css), mime.CSS)
}

//CSSFile
func (c *Context) CSSFile(cssFile string) {
	bytes, err := c.readFile(cssFile)
	if err != nil {
		c.logger.Error("[CSSFile]%v", err)
		return
	}
	c.Render(bytes, mime.CSS)
}

//JS
func (c *Context) JS(js string) {
	c.Render([]byte(js), mime.JS)
}

//JSFile
func (c *Context) JSFile(jsFile string) {
	bytes, err := c.readFile(jsFile)
	if err != nil {
		c.logger.Error("[JSFile]%v", err)
		return
	}
	c.Render(bytes, mime.JS)
}

//Binary
func (c *Context) Binary(buffer []byte) {
	c.Render(buffer, mime.BINARY)
}

//File
func (c *Context) File(file string) {
	bytes, err := c.readFile(file)
	if err != nil {
		c.logger.Error("[File]%v", err)
		return
	}
	c.Binary(bytes)
}

//Download
func (c *Context) Download(file, fileName string) {
	bytes, err := c.readFile(file)
	if err != nil {
		c.logger.Error("[Download]%v", err)
		return
	}
	fn := url.PathEscape(fileName)
	c.Header().Set(headers.ContentDisposition, fmt.Sprintf("attachment; filename=\"%s\"", fn))
	c.Binary(bytes)
}

//readFile
func (c *Context) readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

//Render
func (c *Context) Render(buffer []byte, contentType string) {
	c.Header().Set(headers.MIME, contentType)
	c.Write(buffer)
}
