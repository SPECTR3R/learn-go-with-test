package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()

	http.Get(a)
	aDuration := time.Since(startA).Seconds()

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB).Seconds()

	if aDuration < bDuration {
		return a
	}
	return b
}
