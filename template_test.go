package flygo

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	app := NewApp()
	app.TemplateFuncs = map[string]interface{}{
		"add": func(a, b int) int {
			return a + b
		},
	}
	app.Get("/v", func(c *Context) {
		c.AddViewFuncMap("br", br)
		c.ViewWithData("index", map[string]interface{}{
			"aaa":  "dsfdsfds",
			"list": []string{"sfds", "234343"},
		})
	})
	app.Run()
}

func br(str string) string {
	return str + " -- " + str
}
