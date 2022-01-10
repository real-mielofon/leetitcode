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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	node := head
	for node != nil {
		node.Next, prev, node = prev, node, node.Next
	}
	return prev
}

func copyList(head *ListNode) *ListNode {
	var first *ListNode = nil
	var prev *ListNode = nil
	for head != nil {
		current := &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		if first == nil {
			first = current
		}
		if prev != nil {
			prev.Next = current
		}
		prev = current
		head = head.Next
	}
	return first
}

func isPalindrome(head *ListNode) bool {
	reversed := reverseList(copyList(head))
	for head != nil {
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}
	return true
}

func main() {
	head := NewList([]int{1, 1, 2, 1})
	fmt.Println(head)
	fmt.Println(isPalindrome(head))
}
