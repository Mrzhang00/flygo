package flygo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//try parse view
func parseView(name string) (string, error) {
	data := app.viewCaches[name]
	if data == "" {
		webRoot := app.Config.Flygo.Server.WebRoot
		viewPrefix := app.Config.Flygo.View.Prefix
		viewSuffix := app.Config.Flygo.View.Suffix
		realPath := strings.Join([]string{filepath.Join(webRoot, viewPrefix, name), viewSuffix}, ".")
		buffer, err := ioutil.ReadFile(realPath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("View not found : %s", name))
		}
		data = string(buffer)
	}
	if app.Config.Flygo.View.Cache {
		app.viewCaches[name] = data
	}
	return data, nil
}
