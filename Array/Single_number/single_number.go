package main

import "fmt"

func singleNumber(nums []int) int {
	saveNum := make(map[int]bool)
	for _, num := range nums {
		if _, ok := saveNum[num]; !ok {
			saveNum[num] = true
		} else {
			delete(saveNum, num)
		}
	}
	for k := range saveNum {
		return k
	}
	return -1
}

func singleNumber2(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	result := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(result)
	result = singleNumber2([]int{4, 1, 2, 1, 2})
	fmt.Println(result)
}
