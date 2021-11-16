package main

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCoundown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	got := buffer.String()
	want := "3\n2\n1\nGo!"
	if got != want {
		t.Errorf("Got %q, want \n%q", got, want)
	}
}
