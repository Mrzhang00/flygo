package context

import (
	"bytes"
	"github.com/billcoding/calls"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var templateCaches = make(map[string]string, 0)
var camu = &sync.Mutex{}

func (c *Context) AddFunc(name string, tfunc interface{}) *Context {
	calls.True(name != "" && tfunc != nil, func() {
		c.funcMap[name] = tfunc
	})
	return c
}

func (c *Context) AddFuncMap(funcMap template.FuncMap) *Context {
	calls.True(funcMap != nil && len(funcMap) > 0, func() {
		for k, v := range funcMap {
			c.AddFunc(k, v)
		}
	})
	return c
}

func (c *Context) Template(prefix string, data map[string]interface{}) {
	calls.False(c.templateConfig.Enable, func() {
		c.logger.Warn("[Template]disabled")
	})
	calls.True(c.templateConfig.Enable, func() {
		camu.Lock()
		defer camu.Unlock()
		fileName := prefix + c.templateConfig.Suffix
		realPath := filepath.Join(c.templateConfig.Root, fileName)
		tpl, have := templateCaches[fileName]
		if !have {
			bytes2, err := ioutil.ReadFile(realPath)
			calls.NNil(err, func() {
				c.logger.Error("[Template]%v", err)
			})
			calls.Nil(err, func() {
				tpl = string(bytes2)
				calls.True(c.templateConfig.Cache, func() {
					templateCaches[fileName] = tpl
				})
			})
		}

		calls.NNil(data, func() {
			c.SetDataMap(data)
		})

		t, err := template.New("HTML").Parse(tpl)
		calls.NNil(err, func() {
			c.logger.Error("[Template]%v", err)
		})
		calls.Nil(err, func() {
			t.Funcs(c.funcMap)
			var w bytes.Buffer
			err := t.Execute(&w, c.dataMap)
			calls.NNil(err, func() {
				c.logger.Error("[Template]%v", err)
			})
			calls.Nil(err, func() {
				c.HTML(w.String())
			})
		})
	})
}
