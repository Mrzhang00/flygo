package flygo

import (
	"fmt"
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
	if a.banner {
		fmt.Println(strings.Join(banners, "\n"))
	}
	fmt.Println()
}
