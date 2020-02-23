package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
