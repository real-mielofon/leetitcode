package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func deleteNextNode(node *ListNode) {
	if node.Next != nil {
		node.Next = node.Next.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := make([]*ListNode, 0)
	node := head
	for node != nil {
		if len(arr) < n+1 {
			arr = append(arr[:], node)
		} else {
			arr = append(arr[1:], node)
		}
		node = node.Next
	}
	if len(arr) < n+1 {
		if len(arr) == n {
			return arr[0].Next
		}
		return nil
	} else {
		deleteNextNode(arr[0])
		return head
	}
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(head)
	head = removeNthFromEnd(head, 1)
	fmt.Println(head)

	head = &ListNode{1, nil}
	fmt.Println(head)
	head = removeNthFromEnd(head, 1)
	fmt.Println(head)

	head = &ListNode{1, &ListNode{2, nil}}
	fmt.Println(head)
	head = removeNthFromEnd(head, 2)
	fmt.Println(head)

}
