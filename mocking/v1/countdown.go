package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout)
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)

	}
	fmt.Fprint(out, finalWord)
}
