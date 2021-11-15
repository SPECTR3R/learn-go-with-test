package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// t. Run ("collection of any size", func(t *testing.T) {
// numbers := []int{1, 2, 3}
// got := Sum (numbers)
// want := 6
// if got != want {
// t. Errorf ("got %d want %d given, %v", got, want, numbers)
// })}
