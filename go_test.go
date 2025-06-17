package main

import "testing"

func TestScope(t *testing.T) {
	var a bool
	(func() {
		a := true
		t.Log(a)
	})()
	t.Log(a)
}
