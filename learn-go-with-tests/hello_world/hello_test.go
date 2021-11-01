package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("SW")
	want := "Hello, SW"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
