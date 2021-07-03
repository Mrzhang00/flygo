package flygo

import (
	"io/ioutil"
	"path/filepath"
)

func (a *App) parseYml() {
	file, err := filepath.Abs(a.ConfigFile)
	if err != nil {
		a.Logger.Debugln(err)
		return
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		a.Logger.Debugf("parseYml: %v", err)
	} else {
		err = a.Config.Unmarshal(bytes)
		if err != nil {
			if err != nil {
				a.Logger.Debugf("parseYml: %v", err)
			}
		}
	}
}
