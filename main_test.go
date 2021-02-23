package main

import "testing"

func TestPort(t *testing.T) {
	p := getPort()
	if p != 8080 {
		t.Error("expected port to be 8080")
	}
}
