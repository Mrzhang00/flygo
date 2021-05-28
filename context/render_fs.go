package context

import (
	"embed"
	"github.com/billcoding/flygo/mime"
)

// TextFSFile render text FS file
func (ctx *Context) TextFSFile(fs embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[TextFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.TEXT).Build())
}

// JSONFSFile render JSON FS file
func (ctx *Context) JSONFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[JSONFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.JSON).Build())
}

// XMLFSFile render XML FS file
func (ctx *Context) XMLFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[XMLFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.XML).Build())
}

// ImageFSFile render image FS file
func (ctx *Context) ImageFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[ImageFile]%v", err)
		panic(err)
	}
	ctx.Image(bytes)
}

// ICOFSFile render ico FS file
func (ctx *Context) ICOFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[ICOFile]%v", err)
		panic(err)
	}
	ctx.ICO(bytes)
}

// BMPFSFile render bmp FS file
func (ctx *Context) BMPFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[BMPFile]%v", err)
		panic(err)
	}
	ctx.BMP(bytes)
}

// JPGFSFile render jpg FS file
func (ctx *Context) JPGFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[JPGFile]%v", err)
		panic(err)
	}
	ctx.JPG(bytes)
}

// JPEGFSFile render jpeg FS file
func (ctx *Context) JPEGFSFile(fs *embed.FS, name string) {
	ctx.JPGFSFile(fs, name)
}

// PNGJSFile render png FS file
func (ctx *Context) PNGJSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[PNGFile]%v", err)
		panic(err)
	}
	ctx.PNG(bytes)
}

// GIFFSFile render gif FS file
func (ctx *Context) GIFFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[GIFFile]%v", err)
		panic(err)
	}
	ctx.GIF(bytes)
}

// HTMLFSFile render html FS file
func (ctx *Context) HTMLFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[HTMLFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.HTML).Build())
}

// CSSFSFile render css FS file
func (ctx *Context) CSSFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[CSSFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.CSS).Build())
}

// JSFSFile render js FS file
func (ctx *Context) JSFSFile(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[JSFile]%v", err)
		panic(err)
	}
	ctx.Render(RenderBuilder().Buffer(bytes).ContentType(mime.JS).Build())
}

// FileFS render file
func (ctx *Context) FileFS(fs *embed.FS, name string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		ctx.Logger.Errorf("[File]%v", err)
		panic(err)
	}
	ctx.Binary(bytes)
}
