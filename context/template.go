package context

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

var templateCaches = make(map[string]string, 0)

// AddFunc add func
func (ctx *Context) AddFunc(name string, funcMap interface{}) *Context {
	if name != "" && funcMap != nil {
		ctx.funcMap[name] = funcMap
	}
	return ctx
}

// AddFuncMap add funcMap
func (ctx *Context) AddFuncMap(funcMap template.FuncMap) *Context {
	if funcMap != nil && len(funcMap) > 0 {
		for k, v := range funcMap {
			ctx.AddFunc(k, v)
		}
	}
	return ctx
}

// Template render template
func (ctx *Context) Template(prefix string, data map[string]interface{}) {
	if !ctx.templateConfig.Enable {
		ctx.Logger.Warnf("template: disabled")
	} else {
		fileName := prefix + ctx.templateConfig.Suffix
		realPath := filepath.Join(ctx.templateConfig.Root, fileName)
		tpl, have := templateCaches[fileName]
		if !have {
			bytes2, err := ioutil.ReadFile(realPath)
			if err != nil {
				panic(err)
			} else {
				tpl = string(bytes2)
				if ctx.templateConfig.Cache {
					templateCaches[fileName] = tpl
				}
			}
		}
		if data != nil {
			ctx.SetDataMap(data)
		}
		t, err := template.New("HTML").Parse(tpl)
		if err != nil {
			panic(err)
		} else {
			t.Funcs(ctx.funcMap)
			var w bytes.Buffer
			err := t.Execute(&w, ctx.dataMap)
			if err != nil {
				panic(err)
			} else {
				ctx.HTML(w.String())
			}
		}
	}
}
