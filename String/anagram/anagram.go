package main

import (
	"fmt"
	"sort"
)

type runes []rune

func (r *runes) Len() int {
	return len(*r)
}

func (r *runes) Less(i, j int) bool {
	return (*r)[i] < (*r)[j]
}

func (r *runes) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func isAnagram(s string, t string) bool {
	sBytes := runes(s)
	tBytes := runes(t)

	sort.Sort(&sBytes)
	sort.Sort(&tBytes)

	for i := 0; i < sBytes.Len(); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("", ""))
}
