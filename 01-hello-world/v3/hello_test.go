package main

import "testing"

const englishHelloPrefixTest = "Hello, "

func TestHello(t *testing.T) {
	got := Hello("world")
	want := englishHelloPrefixTest + "world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
