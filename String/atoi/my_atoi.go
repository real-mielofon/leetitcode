package main

import (
	"fmt"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	const maxInt = 2 << 30

	s = strings.TrimSpace(s)
	sign := 1
	res := 0
	inNumber := false

outFor:
	for _, r := range s {
		switch {
		case r == '-':
			if !inNumber {
				sign = -1
				inNumber = true
			} else {
				break outFor
			}
		case r == '+':
			if !inNumber {
				inNumber = true
			} else {
				break outFor
			}
		case unicode.IsDigit(r):
			n := int(r) - int('0')
			res *= 10
			res += n
			if res*sign < -maxInt {
				return -maxInt
			}
			if res*sign > maxInt-1 {
				return maxInt - 1
			}

			inNumber = true
		default:
			break outFor
		}
	}
	return sign * res
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))

}
