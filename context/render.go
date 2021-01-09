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

type Render struct {
	rended      bool
	Buffer      []byte
	Header      http.Header
	Cookies     []*http.Cookie
	Code        int
	ContentType string
}

func (c *Context) Render() *Render {
	return c.render
}

func (r *Render) Rended() bool {
	return r.rended
}

func (c *Context) Rende(r *Render) {
	c.render.rended = true
	if r.Buffer != nil {
		c.render.Buffer = r.Buffer
	}
	if r.ContentType != "" {
		c.render.ContentType = r.ContentType
	}
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			c.render.Header[k] = v
		}
	}
	if len(r.Cookies) > 0 {
		for _, cookie := range r.Cookies {
			c.render.Cookies = append(c.render.Cookies, cookie)
		}
	}
	if r.Code > 0 {
		c.render.Code = r.Code
	}
}

func (c *Context) Text(text string) {
	c.Rende(RenderBuilder().Buffer([]byte(text)).ContentType(mime.TEXT).Build())
}

func (c *Context) TextFile(textFile string) {
	bytes, err := readFile(textFile)
	if err != nil {
		c.logger.Error("[TextFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer([]byte(string(bytes))).ContentType(mime.TEXT).Build())
}

func (c *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.logger.Error("[JSON]%v", err.Error())
		return
	}
	c.Rende(RenderBuilder().Buffer(jsonData).ContentType(mime.JSON).Build())
}

func (c *Context) JSONText(json string) {
	c.Rende(RenderBuilder().Buffer([]byte(json)).ContentType(mime.JSON).Build())
}

func (c *Context) JSONFile(jsonFile string) {
	bytes, err := readFile(jsonFile)
	if err != nil {
		c.logger.Error("[JSONFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.JSON).Build())
}

func (c *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		c.logger.Error("[XML]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(xmlData).ContentType(mime.XML).Build())
}

func (c *Context) XMLText(xml string) {
	c.Rende(RenderBuilder().Buffer([]byte(xml)).ContentType(mime.XML).Build())
}

func (c *Context) XMLFile(xmlFile string) {
	bytes, err := readFile(xmlFile)
	if err != nil {
		c.logger.Error("[XMLFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.XML).Build())
}

func (c *Context) Image(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

func (c *Context) ImageFile(imageFile string) {
	bytes, err := readFile(imageFile)
	if err != nil {
		c.logger.Error("[ImageFile]%v", err)
		return
	}
	c.Image(bytes)
}

func (c *Context) ICO(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.ICO).Build())
}

func (c *Context) ICOFile(icoFile string) {
	bytes, err := readFile(icoFile)
	if err != nil {
		c.logger.Error("[ICOFile]%v", err)
		return
	}
	c.ICO(bytes)
}

func (c *Context) BMP(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.BMP).Build())
}

func (c *Context) BMPFile(bmpFile string) {
	bytes, err := readFile(bmpFile)
	if err != nil {
		c.logger.Error("[BMPFile]%v", err)
		return
	}
	c.BMP(bytes)
}

func (c *Context) JPG(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

func (c *Context) JPGFile(jpgFile string) {
	bytes, err := readFile(jpgFile)
	if err != nil {
		c.logger.Error("[JPGFile]%v", err)
		return
	}
	c.JPG(bytes)
}

func (c *Context) JPEG(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

func (c *Context) JPEGFile(jpegFile string) {
	c.JPGFile(jpegFile)
}

func (c *Context) PNG(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.PNG).Build())
}

func (c *Context) PNGFile(pngFile string) {
	bytes, err := readFile(pngFile)
	if err != nil {
		c.logger.Error("[PNGFile]%v", err)
		return
	}
	c.PNG(bytes)
}

func (c *Context) GIF(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.GIF).Build())
}

func (c *Context) GIFFile(gifFile string) {
	bytes, err := readFile(gifFile)
	if err != nil {
		c.logger.Error("[GIFFile]%v", err)
		return
	}
	c.GIF(bytes)
}

func (c *Context) HTML(html string) {
	c.Rende(RenderBuilder().Buffer([]byte(html)).ContentType(mime.HTML).Build())
}

func (c *Context) HTMLFile(htmlFile string) {
	bytes, err := readFile(htmlFile)
	if err != nil {
		c.logger.Error("[HTMLFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.HTML).Build())
}

func (c *Context) CSS(css string) {
	c.Rende(RenderBuilder().Buffer([]byte(css)).ContentType(mime.CSS).Build())
}

func (c *Context) CSSFile(cssFile string) {
	bytes, err := readFile(cssFile)
	if err != nil {
		c.logger.Error("[CSSFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.CSS).Build())
}

func (c *Context) JS(js string) {
	c.Rende(RenderBuilder().Buffer([]byte(js)).ContentType(mime.JS).Build())
}

func (c *Context) JSFile(jsFile string) {
	bytes, err := readFile(jsFile)
	if err != nil {
		c.logger.Error("[JSFile]%v", err)
		return
	}
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.JS).Build())
}

func (c *Context) Binary(buffer []byte) {
	c.Rende(RenderBuilder().Buffer(buffer).ContentType(mime.BINARY).Build())
}

func (c *Context) File(file string) {
	bytes, err := readFile(file)
	if err != nil {
		c.logger.Error("[File]%v", err)
		return
	}
	c.Binary(bytes)
}

func (c *Context) Download(file, fileName string) {
	bytes, err := readFile(file)
	if err != nil {
		c.logger.Error("[Download]%v", err)
		return
	}
	fn := url.PathEscape(fileName)
	c.Rende(RenderBuilder().Buffer(bytes).ContentType(mime.BINARY).Header(http.Header{
		headers.ContentDisposition: []string{fmt.Sprintf("attachment; filename=\"%s\"", fn)},
	}).Build())
}

func readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
