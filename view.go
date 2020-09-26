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
		realPath := strings.Join([]string{app.GetWebRoot(), app.GetViewPrefix(), name}, string(filepath.Separator)) + "." + app.GetViewSuffix()
		buffer, err := ioutil.ReadFile(realPath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("View not found : %s", name))
		}
		data = string(buffer)
	}
	if app.GetViewCache() {
		app.viewCaches[name] = data
	}
	return data, nil
}
