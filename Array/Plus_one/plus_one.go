package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func plusOneBest(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]

		digits[i] = (d + 1) % 10

		if (d+1)/10 == 0 {
			return digits
		}

	}

	digits = append([]int{1}, digits...)
	return digits
}
func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))

	fmt.Println(plusOneBest([]int{1, 2, 3}))
}
