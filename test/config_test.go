package test

import (
	"testing"
)

// TestBannerFile test
func TestServeHost(t *testing.T) {
	testServeWithHost(t, "127.0.0.1", 8090, "GET", "TestServeHost")
}

// TestServePort test
func TestServePort(t *testing.T) {
	testServe(t, nextPort(), "GET", "TestServePort")
}
