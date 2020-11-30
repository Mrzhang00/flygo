package flygo

import (
	"testing"
)

//Test banner
func TestBanner(t *testing.T) {

	//Disable banner
	app := NewApp()
	app.ConfigFile = "test_ymls/test_banner.yml"

	app.Run()
}
