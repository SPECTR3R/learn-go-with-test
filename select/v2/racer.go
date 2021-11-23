package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measure(a)
	bDuration := measure(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measure(url string) (duration time.Duration) {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
