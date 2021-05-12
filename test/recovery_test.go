package test

import (
	"github.com/billcoding/flygo"
	"github.com/billcoding/flygo/context"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestRecovery(t *testing.T) {
	go func() {
		app := flygo.GetApp()
		app.GET("", func(ctx *context.Context) {
			panic(`This is panic message`)
		})
		app.UseRecovery()
		app.RecoveryConfig("code_val", 5000, "message")
		app.Run()
	}()
	<-time.After(time.Second)
	resp, _ := http.Get("http://localhost")
	bytes, _ := io.ReadAll(resp.Body)
	t.Log(string(bytes))
}
