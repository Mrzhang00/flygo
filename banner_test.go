package flygo

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//Test banner text
func TestBannerText(t *testing.T) {

	//Disable banner
	app := NewApp()
	app.Config.Banner.Type = "text"
	app.Config.Banner.Text = `

THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!
THIS IS FLYGO BANNER TEXT!!!!

`
	app.Run()
}

//Test banner
func TestBannerFile(t *testing.T) {

	ioutil.WriteFile("banner.txt", []byte(`

THIS IS FLYGO BANNER FROM FILE!!!!

`), 0760)

	//Disable banner
	app := NewApp()
	app.Config.Banner.Type = "file"

	go func() {
		time.AfterFunc(time.Second, func() {
			os.Remove("banner.txt")
		})
	}()

	app.Run()
}
