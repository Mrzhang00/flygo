package flygo

import (
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"syscall"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	var a = GetApp()
	sig := make(chan os.Signal, 1)
	go func() {
		get := func(c *Context) {
			c.Text("get")
		}
		a.Get("/", get).Run()
	}()

	go func() {
		time.AfterFunc(time.Second, func() {

			resp, _ := http.Get("http://localhost")
			bytes, _ := ioutil.ReadAll(resp.Body)
			if !reflect.DeepEqual(string(bytes), "get") {
				t.Fail()
			}

			sig <- syscall.SIGHUP
		})
	}()
	<-sig
}
