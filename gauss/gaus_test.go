package main

import (
	"fmt"
	"log"
	"math"
	"testing"
)

type testCase struct {
	Name string
	a    [][]float64
	b    []float64
	x    []float64
}

// result from above test case turns out to be correct to this tolerance.
const ε = 1e-14

func TestWalk(t *testing.T) {
	cases := []testCase{

		{
			Name: "6x6",
			a: [][]float64{
				{1.00, 0.00, 0.00, 0.00, 0.00, 0.00},
				{1.00, 0.63, 0.39, 0.25, 0.16, 0.10},
				{1.00, 1.26, 1.58, 1.98, 2.49, 3.13},
				{1.00, 1.88, 3.55, 6.70, 12.62, 23.80},
				{1.00, 2.51, 6.32, 15.88, 39.90, 100.28},
				{1.00, 3.14, 9.87, 31.01, 97.41, 306.02},
			},
			b: []float64{-0.01, 0.61, 0.91, 0.99, 0.60, 0.02},
			x: []float64{
				-0.01, 1.602790394502114, -1.6132030599055613,
				1.2454941213714368, -0.4909897195846576, 0.065760696175232,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			x, err := GaussPartial(tc.a, tc.b)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(x)
			for i, xi := range x {
				if math.Abs(tc.x[i]-xi) > ε {
					log.Println("out of tolerance")
					t.Errorf("for %v, want %v", x[i], tc.x)
				}
			}
		})
	}
}
