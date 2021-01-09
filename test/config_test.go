package test

import (
	"testing"
)

func TestServeHost(t *testing.T) {
	testServeWithHost(t, "127.0.0.1", 8090, "GET", "TestServeHost")
}

func TestServePort(t *testing.T) {
	testServe(t, nextPort(), "GET", "TestServePort")
}
