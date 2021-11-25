package main

import (
	"errors"
	"fmt"
	"math"
)



func GaussPartial(a0 [][]float64, b0 []float64) ([]float64, error) {
	// make augmented matrix
	m := len(b0)
	a := make([][]float64, m)

	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}

	// WP algorithm from Gaussian elimination page
	// produces row-eschelon form
	for k := range a {
		iMax := 0
		max := -1.
		for i := k; i < m; i++ {
			row := a[i]
			// compute scale factor s = max abs in row
			s := -1.
			for j := k; j < m; j++ {
				x := math.Abs(row[j])
				if x > s {
					s = x
				}
			}
			// scale the abs used to pick the pivot.
			if abs := math.Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		fmt.Println(a)
		if a[iMax][k] == 0 {
			return nil, errors.New("singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
		fmt.Println(a)
	} // end of WP algorithm.

	// now back substitute to get result.
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}
