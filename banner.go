package flygo

import (
	"fmt"
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

//printBanner
func (a *App) printBanner() {
	if !a.Config.Banner.Enable {
		return
	}
	switch a.Config.Banner.Type {
	case "default":
		fmt.Println(strings.Join(banners, "\n"))
		break
	case "text":
		fmt.Println(a.Config.Banner.Text)
		break
	case "file":
		bytes, err := ioutil.ReadFile(a.Config.Banner.File)
		if err != nil {
			a.Logger.Error("[printBanner]%v", err)
			return
		}
		fmt.Printf("%v\n", string(bytes))
		break
	}
	fmt.Println()
}
