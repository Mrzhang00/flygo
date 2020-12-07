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
	bytes, err := c.readRealPath(textFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Text(string(bytes))
}

//Response json to client
func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.app.Logger.Error(err.Error())
		return
	}
	c.Render(jsonData, contentTypeJson)
}

//Response json text to client
func (c *Context) JSONText(json string) {
	c.Render([]byte(json), contentTypeJson)
}

//Response json file to client
func (c *Context) JSONFile(jsonFile string) {
	bytes, err := c.readRealPath(jsonFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeJson)
}

//Response xml to client
func (c *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		c.app.Logger.Error(err.Error())
		return
	}
	c.Render(xmlData, contentTypeXml)
}

//Response xml text to client
func (c *Context) XMLText(xml string) {
	c.Render([]byte(xml), contentTypeXml)
}

//Response xml file to client
func (c *Context) XMLFile(xmlFile string) {
	bytes, err := c.readRealPath(xmlFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
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
	bytes, err := c.readRealPath(imageFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
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
	bytes, err := c.readRealPath(bmpFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
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
	bytes, err := c.readRealPath(jpgFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
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
	bytes, err := c.readRealPath(pngFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
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
	bytes, err := c.readRealPath(gifFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.GIF(bytes)
}

//Response html to client
func (c *Context) HTML(html string) {
	c.Render([]byte(html), contentTypeHtml)
}

//Response html file to client
func (c *Context) HTMLFile(htmlFile string) {
	bytes, err := c.readRealPath(htmlFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeHtml)
}

//Response css to client
func (c *Context) CSS(css string) {
	c.Render([]byte(css), contentTypeCSS)
}

//Response css file to client
func (c *Context) CSSFile(cssFile string) {
	bytes, err := c.readRealPath(cssFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeCSS)
}

//Response js to client
func (c *Context) JS(js string) {
	c.Render([]byte(js), contentTypeJS)
}

//Response js file to client
func (c *Context) JSFile(jsFile string) {
	bytes, err := c.readRealPath(jsFile)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Render(bytes, contentTypeJS)
}

//Response file to client
func (c *Context) Binary(buffer []byte) {
	c.Render(buffer, contentTypeBinary)
}

//Response file to client
func (c *Context) File(file string) {
	bytes, err := c.readRealPath(file)
	if err != nil {
		c.app.Logger.Error("%v", err)
		return
	}
	c.Binary(bytes)
}

//Response download file to client
func (c *Context) Download(file, fileName string) {
	bytes, err := c.readRealPath(file)
	if err != nil {
		c.app.Logger.Error("%v", err)
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

func (c *Context) getRealPath(fileName string) string {
	return filepath.Join(c.app.Config.Flygo.Server.WebRoot, fileName)
}

func (c *Context) readRealPath(fileName string) ([]byte, error) {
	rp := c.getRealPath(fileName)
	return ioutil.ReadFile(rp)
}
