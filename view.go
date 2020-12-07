package flygo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//try parse view
func (c *Context) parseView(name string) (string, error) {
	data := c.app.viewCaches[name]
	if data == "" {
		webRoot := c.app.Config.Flygo.Server.WebRoot
		viewPrefix := c.app.Config.Flygo.View.Prefix
		viewSuffix := c.app.Config.Flygo.View.Suffix
		realPath := strings.Join([]string{filepath.Join(webRoot, viewPrefix, name), viewSuffix}, ".")
		buffer, err := ioutil.ReadFile(realPath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("View not found : %s", name))
		}
		data = string(buffer)
	}
	if c.app.Config.Flygo.View.Cache {
		c.app.viewCaches[name] = data
	}
	return data, nil
}
