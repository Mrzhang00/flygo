package test

import (
	"testing"
)

// TestServeGet test
func TestServeGet(t *testing.T) {
	testServe(t, nextPort(), "GET", "TestServeGet")
}

// TestServePost test
func TestServePost(t *testing.T) {
	testServe(t, nextPort(), "POST", "TestServePost")
}

// TestServeDelete test
func TestServeDelete(t *testing.T) {
	testServe(t, nextPort(), "DELETE", "TestServeDelete")
}

// TestServePut test
func TestServePut(t *testing.T) {
	testServe(t, nextPort(), "PUT", "TestServePut")
}

// TestServePatch test
func TestServePatch(t *testing.T) {
	testServe(t, nextPort(), "PATCH", "TestServePatch")
}

// TestServeOptions test
func TestServeOptions(t *testing.T) {
	testServe(t, nextPort(), "OPTIONS", "TestServeOptions")
}
