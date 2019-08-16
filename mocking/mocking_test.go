package macking

import (
	"bytes"
	"reflect"
	"testing"
)

type spy struct {
	Calls int
}

func (s *spy) Sleep() {
	s.Calls++
}

type sleepOrder struct {
	call []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *sleepOrder) Sleep() {
	s.call = append(s.call, sleep)
}

func (s *sleepOrder) Write(p []byte) (n int, err error) {
	s.call = append(s.call, write)
	return
}

func TestDo(t *testing.T) {

	t.Run("check result", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		s := &spy{}
		Do(buffer, s)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}

		if s.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", s.Calls)
		}
	})

	t.Run("check order", func(t *testing.T) {
		s := &sleepOrder{}
		Do(s, s)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, s.call) {
			t.Errorf("want %v got %v", want, s.call)
		}
	})
}
