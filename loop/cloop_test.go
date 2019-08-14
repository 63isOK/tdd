package loop

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	t.Run("repeat 5 char", func(t *testing.T) {
		got := repeat("c", 5)
		want := "ccccc"
		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat 7 char", func(t *testing.T) {
		got := repeat("a", 7)
		want := "aaaaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("q", 3)
	}
}
