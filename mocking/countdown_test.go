//package mocking
package main

import (
	"bytes"
	"testing"
)

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep(_ int) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	m := &MockSleeper{}
	Countdown(buffer, m)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if m.Calls != 4 {
		t.Errorf("slept %d instead of %d times", m.Calls, 4)
	}
}
