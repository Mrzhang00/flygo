package flygo

import (
	"testing"
)

//Test field validate min number
func TestValidateMin(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Min(10)

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate max number
func TestValidateMax(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Max(10)

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate fixed length string
func TestValidateLength(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Length(10)

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate fixed string
func TestValidateFixed(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Fixed("123")

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate min length string
func TestValidateMinLength(t *testing.T) {
	app := NewApp()

	idField := NewField("id").MinLength(5)

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate max length string
func TestValidateMaxLength(t *testing.T) {
	app := NewApp()

	idField := NewField("id").MaxLength(10)

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate optional val
func TestValidateOptional(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Enums("10", "20")

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}

//Test field validate regex
func TestValidateRegex(t *testing.T) {
	app := NewApp()

	idField := NewField("id").Regex("^\\d{5}$")

	app.Get("/", func(c *Context) {
		c.Text("id = " + c.Param("id"))
	}, idField)

	app.Run()
}
