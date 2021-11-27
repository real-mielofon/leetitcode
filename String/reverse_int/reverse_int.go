package main

import (
	"fmt"
)

func reverse(x int) int {
	var maxInt32 = 1 << 31
	if (x < -maxInt32) || x > (maxInt32-1) {
		return 0
	}
	n := 0
	for x != 0 {
		n *= 10
		n += x % 10
		x = x / 10
	}
	return n
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-123))
}
