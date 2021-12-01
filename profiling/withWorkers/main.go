package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

const (
	output     = "output.png"
	width      = 2048
	height     = 2048
	complexity = 4
	numWorkers = 8
)

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	f, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	img := createWorkers(width, height)

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func createWorkers(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	c := make(chan int, width)

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for i := range c {
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		c <- i
	}

	close(c)
	w.Wait()
	return m
}

func pixel(i, j, width, height int) color.Color {
	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
