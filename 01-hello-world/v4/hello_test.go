package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("should return name with hello", func(t *testing.T) {
		got := Hello("world")
		want := "Hello, world"

		assertMessage(t, got, want)
	})

	t.Run("should return name as World when name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
}
