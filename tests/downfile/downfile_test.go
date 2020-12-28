package main

import (
	"github.com/billcoding/flygo"
	"testing"
)

func TestDownfile(t *testing.T) {
	downFile := New().Root("/Users/local/tmp2")
	flygo.NewApp().Use(downFile).Run()
}
