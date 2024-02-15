package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Santi")
	want := "Hello Santi!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
