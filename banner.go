package flygo

import (
	"io/ioutil"
	"strings"
)

var banners []string

func init() {
	banners = append(banners, ` _______  __      ____    ____  _______   ______`)
	banners = append(banners, `|   ____||  |     \   \  /   / /  _____| /  __  \`)
	banners = append(banners, `|  |__   |  |      \   \/   / |  |  __  |  |  |  |`)
	banners = append(banners, `|   __|  |  |       \_    _/  |  | |_ | |  |  |  |`)
	banners = append(banners, "|  |     |  `----.    |  |    |  |__| | |  `--'  |")
	banners = append(banners, `|__|     |_______|    |__|     \______|  \______/`)
}

//Print the banner
func (a *App) printBanner() {
	if !a.Config.Flygo.Banner.Enable {
		return
	}
	switch a.Config.Flygo.Banner.Type {
	case "default":
		a.Info(strings.Join(banners, "\n"))
		break
	case "text":
		a.Info(a.Config.Flygo.Banner.Text)
		break
	case "file":
		bytes, err := ioutil.ReadFile(a.Config.Flygo.Banner.File)
		if err != nil {
			a.Error("%v", err)
			return
		}
		a.Info("%v\n", string(bytes))
		break
	}
	a.Info("\n")
}
