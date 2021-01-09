package test

import (
	"github.com/billcoding/flygo"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// TestBannerText test
func TestBannerText(t *testing.T) {
	go func() {
		app := flygo.GetApp()
		app.Config.Flygo.Banner.Type = "text"
		app.Config.Flygo.Banner.Text = `THIS IS FLYGO CUSTOM BANNER TEXT!!!!`
		setAppPort()
		app.Run()
	}()
	<-time.After(time.Second)
}

// TestBannerFile test
func TestBannerFile(t *testing.T) {
	go func() {
		_ = ioutil.WriteFile("banner.txt", []byte(`THIS IS FLYGO BANNER FROM FILE!!!!`), 0760)
		app := flygo.GetApp()
		app.Config.Flygo.Banner.Type = "file"
		setAppPort()
		app.Run()
	}()
	<-time.After(time.Second)
	_ = os.Remove("banner.txt")
}
