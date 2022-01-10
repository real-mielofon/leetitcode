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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var other *ListNode
	if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
		result = list1
		other = list2
	} else {
		result = list2
		other = list1
	}
	if result == nil {
		return result
	}
	current := result
	for current.Next != nil && other != nil {
		// skip current
		for current.Next != nil && current.Next.Val <= other.Val {
			current = current.Next
		}
		current.Next, other = other, current.Next
	}
	if other != nil {
		current.Next = other
	}
	return result
}

func main() {

	//Input: list1 = [1,2,4], list2 = [1,3,4]
	//Output: [1,1,2,3,4,4]
	//
	//Example 2:
	//
	//Input: list1 = [], list2 = []
	//Output: []
	//
	//Example 3:
	//
	//Input: list1 = [], list2 = [0]
	//Output: [0]

	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(list1)
	fmt.Println(list2)
	mergedList := mergeTwoLists(list1, list2)
	fmt.Println(mergedList)

	list1 = nil
	list2 = nil
	fmt.Println(list1)
	fmt.Println(list2)
	mergedList = mergeTwoLists(list1, list2)
	fmt.Println(mergedList)

	list1 = nil
	list2 = &ListNode{0, nil}
	fmt.Println(list1)
	fmt.Println(list2)
	mergedList = mergeTwoLists(list1, list2)
	fmt.Println(mergedList)

	list1 = &ListNode{2, nil}
	list2 = &ListNode{1, nil}
	fmt.Println(list1)
	fmt.Println(list2)
	mergedList = mergeTwoLists(list1, list2)
	fmt.Println(mergedList)
}
