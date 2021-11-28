package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		s := countAndSay(n - 1)
		sout := ""
		start := 0
		for i, c := range s {
			if c != []rune(s)[start] {
				sout = sout + fmt.Sprintf("%d%c", i-start, []rune(s)[start])
				start = i
			}
		}
		n := len([]rune(s))
		n = len([]rune(s)) - start
		sym := []rune(s)[start]
		sout = sout + fmt.Sprintf("%d%c", n, sym)
		return sout
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(countAndSay(i))
	}
}
