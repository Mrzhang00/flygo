package context

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var templateCaches = make(map[string]string, 0)
var mutex = &sync.Mutex{}

// AddFunc add func
func (c *Context) AddFunc(name string, tfunc interface{}) *Context {
	if name != "" && tfunc != nil {
		c.funcMap[name] = tfunc
	}
	return c
}

// AddFuncMap add funcMap
func (c *Context) AddFuncMap(funcMap template.FuncMap) *Context {
	if funcMap != nil && len(funcMap) > 0 {
		for k, v := range funcMap {
			c.AddFunc(k, v)
		}
	}
	return c
}

// Template render template
func (c *Context) Template(prefix string, data map[string]interface{}) {
	if !c.templateConfig.Enable {
		c.logger.Warn("[Template]disabled")
	} else {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := prefix + c.templateConfig.Suffix
		realPath := filepath.Join(c.templateConfig.Root, fileName)
		tpl, have := templateCaches[fileName]
		if !have {
			bytes2, err := ioutil.ReadFile(realPath)
			if err != nil {
				c.logger.Error("[Template]%v", err)
			} else {
				tpl = string(bytes2)
				if c.templateConfig.Cache {
					templateCaches[fileName] = tpl
				}
			}
		}
		if data != nil {
			c.SetDataMap(data)
		}
		t, err := template.New("HTML").Parse(tpl)
		if err != nil {
			c.logger.Error("[Template]%v", err)
		} else {
			t.Funcs(c.funcMap)
			var w bytes.Buffer
			err := t.Execute(&w, c.dataMap)
			if err != nil {
				c.logger.Error("[Template]%v", err)
			} else {
				c.HTML(w.String())
			}
		}
	}
}
