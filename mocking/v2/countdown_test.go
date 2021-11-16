package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCoundown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := "3\n2\n1\nGo!"
	if got != want {
		t.Errorf("Got %q, want \n%q", got, want)
	}
	if spySleeper.Calls != countdownStart {
		t.Errorf("not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.Calls)
	}
}
