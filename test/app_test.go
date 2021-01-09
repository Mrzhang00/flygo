package test

import (
	"testing"
)

func TestServeGet(t *testing.T) {
	testServe(t, nextPort(), "GET", "TestServeGet")
}

func TestServePost(t *testing.T) {
	testServe(t, nextPort(), "POST", "TestServePost")
}

func TestServeDelete(t *testing.T) {
	testServe(t, nextPort(), "DELETE", "TestServeDelete")
}

func TestServePut(t *testing.T) {
	testServe(t, nextPort(), "PUT", "TestServePut")
}

func TestServePatch(t *testing.T) {
	testServe(t, nextPort(), "PATCH", "TestServePatch")
}

func TestServeOptions(t *testing.T) {
	testServe(t, nextPort(), "OPTIONS", "TestServeOptions")
}
