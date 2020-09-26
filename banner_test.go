package flygo

import "testing"

//Test banner
func TestBanner(t *testing.T) {

	//Disable banner
	NewApp().Banner(false).Run()

}
