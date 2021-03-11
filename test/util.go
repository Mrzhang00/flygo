package test

import (
	"fmt"
	"github.com/billcoding/flygo"
	"github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func req(url, method string) string {
	req, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		panic(err)
	}
	req.BasicAuth()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

var port = int32(60000)

func nextPort() int {
	atomic.AddInt32(&port, 1)
	return int(port)
}

func setAppPort() {
	flygo.GetApp().Config.Flygo.Server.Port = nextPort()
}

func testServe(t *testing.T, port int, method, expected string) {
	testServeWithHost(t, "", port, method, expected)
}

func testServeWithHost(t *testing.T, host string, port int, method, expected string) {
	go func() {
		app := flygo.NewApp()
		app.Config.Flygo.Server.Host = host
		app.Config.Flygo.Server.Port = port
		app.Route(strings.ToUpper(method), "/", func(ctx *context.Context) {
			c.Text(expected)
		}).Run()
	}()
	<-time.After(time.Millisecond)
	if host == "" {
		host = "localhost"
	}
	result := req(fmt.Sprintf("http://%s:%d", host, port), method)
	if result != expected {
		t.Fatalf("Unexpected result : '%s', expected result : '%s'\n", result, expected)
	}
}
