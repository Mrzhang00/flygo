package context

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/billcoding/flygo/headers"
	"github.com/billcoding/flygo/mime"
	"net/http"
	"net/url"
	"os"
)

// Render struct
type Render struct {
	Buffer      []byte
	Header      http.Header
	Cookies     []*http.Cookie
	Code        int
	ContentType string
}

// Rendered panic(err) the render
func (ctx *Context) Rendered() *Render {
	return ctx.render
}

// Render from r
func (ctx *Context) Render(r *Render) {
	if r.Buffer != nil {
		ctx.render.Buffer = r.Buffer
	}
	if r.ContentType != "" {
		ctx.render.ContentType = r.ContentType
	}
	if r.Header != nil && len(r.Header) > 0 {
		for k, v := range r.Header {
			ctx.render.Header[k] = v
		}
	}
	if r.Cookies != nil && len(r.Cookies) > 0 {
		for _, cookie := range r.Cookies {
			ctx.render.Cookies = append(ctx.render.Cookies, cookie)
		}
	}
	if r.Code > 0 {
		ctx.render.Code = r.Code
	}
}

// Text render text
func (ctx *Context) Text(text string) {
	ctx.Render(RenderBuilder().Buffer([]byte(text)).ContentType(mime.TEXT).Build())
}

// TextFile render text file
func (ctx *Context) TextFile(textFile string) {
	bytes, err := readFile(textFile)
	if err != nil {
		ctx.Logger.Errorf("[TextFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.TEXT).Build())
}

// JSON render JSON
func (ctx *Context) JSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		ctx.Logger.Errorf("[JSON]%v", err.Error())
	}
	ctx.Render(RenderBuilder().Buffer(jsonData).ContentType(mime.JSON).Build())
}

// JSONText render JSON text
func (ctx *Context) JSONText(json string) {
	ctx.Render(RenderBuilder().Buffer([]byte(json)).ContentType(mime.JSON).Build())
}

// JSONFile render JSON file
func (ctx *Context) JSONFile(jsonFile string) {
	bytes, err := readFile(jsonFile)
	if err != nil {
		ctx.Logger.Errorf("[JSONFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.JSON).Build())
}

// XML render XML
func (ctx *Context) XML(data interface{}) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		ctx.Logger.Errorf("[XML]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(xmlData).ContentType(mime.XML).Build())
}

// XMLText render XML text
func (ctx *Context) XMLText(xml string) {
	ctx.Render(RenderBuilder().Buffer([]byte(xml)).ContentType(mime.XML).Build())
}

// XMLFile render XML file
func (ctx *Context) XMLFile(xmlFile string) {
	bytes, err := readFile(xmlFile)
	if err != nil {
		ctx.Logger.Errorf("[XMLFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.XML).Build())
}

// Image render image
func (ctx *Context) Image(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

// ImageFile render image file
func (ctx *Context) ImageFile(imageFile string) {
	bytes, err := readFile(imageFile)
	if err != nil {
		ctx.Logger.Errorf("[ImageFile]%v", err)
		panic(err)
	}
	ctx.Image(bytes)
}

// ICO render ico
func (ctx *Context) ICO(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.ICO).Build())
}

// ICOFile render ico file
func (ctx *Context) ICOFile(icoFile string) {
	bytes, err := readFile(icoFile)
	if err != nil {
		ctx.Logger.Errorf("[ICOFile]%v", err)
		panic(err)
	}
	ctx.ICO(bytes)
}

// BMP render bmp
func (ctx *Context) BMP(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.BMP).Build())
}

// BMPFile render bmp file
func (ctx *Context) BMPFile(bmpFile string) {
	bytes, err := readFile(bmpFile)
	if err != nil {
		ctx.Logger.Errorf("[BMPFile]%v", err)
		panic(err)
	}
	ctx.BMP(bytes)
}

// JPG render jpg
func (ctx *Context) JPG(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

// JPGFile render jpg file
func (ctx *Context) JPGFile(jpgFile string) {
	bytes, err := readFile(jpgFile)
	if err != nil {
		ctx.Logger.Errorf("[JPGFile]%v", err)
		panic(err)
	}
	ctx.JPG(bytes)
}

// JPEG render jpeg
func (ctx *Context) JPEG(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.JPG).Build())
}

// JPEGFile render jpeg file
func (ctx *Context) JPEGFile(jpegFile string) {
	ctx.JPGFile(jpegFile)
}

// PNG render png
func (ctx *Context) PNG(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.PNG).Build())
}

// PNGFile render png file
func (ctx *Context) PNGFile(pngFile string) {
	bytes, err := readFile(pngFile)
	if err != nil {
		ctx.Logger.Errorf("[PNGFile]%v", err)
		panic(err)
	}
	ctx.PNG(bytes)
}

// GIF render gif
func (ctx *Context) GIF(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.GIF).Build())
}

// GIFFile render gif file
func (ctx *Context) GIFFile(gifFile string) {
	bytes, err := readFile(gifFile)
	if err != nil {
		ctx.Logger.Errorf("[GIFFile]%v", err)
		panic(err)
	}
	ctx.GIF(bytes)
}

// HTML render html
func (ctx *Context) HTML(html string) {
	ctx.Render(RenderBuilder().Buffer([]byte(html)).ContentType(mime.HTML).Build())
}

// HTMLFile render html file
func (ctx *Context) HTMLFile(htmlFile string) {
	bytes, err := readFile(htmlFile)
	if err != nil {
		ctx.Logger.Errorf("[HTMLFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.HTML).Build())
}

// CSS render css
func (ctx *Context) CSS(css string) {
	ctx.Render(RenderBuilder().Buffer([]byte(css)).ContentType(mime.CSS).Build())
}

// CSSFile render css file
func (ctx *Context) CSSFile(cssFile string) {
	bytes, err := readFile(cssFile)
	if err != nil {
		ctx.Logger.Errorf("[CSSFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.CSS).Build())
}

// JS render js
func (ctx *Context) JS(js string) {
	ctx.Render(RenderBuilder().Buffer([]byte(js)).ContentType(mime.JS).Build())
}

// JSFile render js file
func (ctx *Context) JSFile(jsFile string) {
	bytes, err := readFile(jsFile)
	if err != nil {
		ctx.Logger.Errorf("[JSFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.JS).Build())
}

// Binary render bin
func (ctx *Context) Binary(buffer []byte) {
	ctx.Render(RenderBuilder().Buffer(buffer).ContentType(mime.BINARY).Build())
}

// File render file
func (ctx *Context) File(file string) {
	bytes, err := readFile(file)
	if err != nil {
		ctx.Logger.Errorf("[File]%v", err)
		panic(err)
	}
	ctx.Binary(bytes)
}

// Download render download
func (ctx *Context) Download(file, fileName string) {
	bytes, err := readFile(file)
	if err != nil {
		ctx.Logger.Errorf("[Download]%v", err)
		panic(err)
	}
	fn := url.PathEscape(fileName)
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.BINARY).Header(http.Header{
		headers.ContentDisposition: []string{fmt.Sprintf("attachment; filename=\"%s\"", fn)},
	}).Build())
}

func readFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
