package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}
	i1 := 0
	i2 := 0
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == nums2[i2] {
			result = append(result, nums1[i1])
			i1++
			i2++
		} else if nums1[i1] < nums2[i2] {
			i1++
		} else {
			i2++
		}
	}
	return result
}

func intersect_best(nums1 []int, nums2 []int) []int {

	m := map[int]int{}
	for _, n := range nums1 {
		m[n]++
	}

	result := []int{}
	for _, n := range nums2 {
		if v := m[n]; v > 0 {
			result = append(result, n)
			m[n]--
		}
	}
	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

	fmt.Println(intersect_best([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

}
