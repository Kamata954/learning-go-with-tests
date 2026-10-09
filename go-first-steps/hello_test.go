package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kimani")
	want := "Hello, Kimani"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
