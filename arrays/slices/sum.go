package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }
