package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		s = s + fmt.Sprintf("[%d] ", l.Val)
		l = l.Next
	}
	return s
}

func NewList(arr []int) *ListNode {
	var result *ListNode = nil
	for i := len(arr) - 1; i >= 0; i-- {
		result = &ListNode{
			Val:  arr[i],
			Next: result,
		}
	}
	return result
}

func hasCycle(head *ListNode) bool {
	dict := make(map[*ListNode]bool)
	for head != nil {
		head = head.Next
		if _, ok := dict[head]; ok {
			return true
		} else {
			dict[head] = true
		}
	}
	return false
}

func main() {
	head := NewList([]int{1, 1, 2, 1})
	fmt.Println(head)
	fmt.Println(hasCycle(head))

}
