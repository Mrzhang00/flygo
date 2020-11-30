package flygo

import (
	"strings"
	"testing"
)

//Test field preset default val
func TestPresetDefaultVal(t *testing.T) {
	app := NewApp()

	idField := NewField("id").DefaultVal("1000")

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field preset split val
func TestPresetSplit(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Split(true).DefaultVal("1,2,3,4,5,6")
	app.Get("/", func(c *Context) {
		ids := c.Params("id")
		c.Text("id = " + strings.Join(ids, " - "))
	}, idField)

	app.Run()
}

//Test field preset concat vals
func TestPresetConcat(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Concat(true).DefaultVal("1000")
	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}
