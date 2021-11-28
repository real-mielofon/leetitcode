package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	offsets := make([]int, 0)
	for i := 0; i < len(haystack); i++ {
		r := haystack[i]
		for j := 0; j < len(offsets); j++ {
			offset := offsets[j]
			if r == needle[offset] {
				if offset == len(needle)-1 {
					return i - offset
				}
				offset++
				offsets[j] = offset
			} else {
				// удаляем
				offsets = append(offsets[0:j], offsets[j+1:]...)
				j-- // откат из-за удалённого
			}
		}

		if r == needle[0] {
			offsets = append(offsets[:], 1)
			if len(needle) == 1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("aa", "aaa"))
}
