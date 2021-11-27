package main

import "fmt"

func moveZeroes(nums []int) {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zero] = nums[i]
			zero++
		}
	}
	for i := zero; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
	fmt.Println(s)

	s = []int{0}
	moveZeroes(s)
	fmt.Println(s)
}
