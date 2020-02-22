package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("should return a hello as default to person", func(t *testing.T) {
		got := Hello("Goki", "")
		want := "Hello, Goki"

		assertMessage(t, got, want)
	})

	t.Run("should return a hello as spanish to person", func(t *testing.T) {
		got := Hello("Goki", "es")
		want := "Hola, Goki"

		assertMessage(t, got, want)
	})

	t.Run("should return a hello as french to person", func(t *testing.T) {
		got := Hello("Goki", "fr")
		want := "Bonjour, Goki"

		assertMessage(t, got, want)
	})

	t.Run("should return a hello as default to default when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
}
