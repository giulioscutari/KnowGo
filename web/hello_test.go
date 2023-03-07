package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Olly")
	want := "Hello, Olly!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}