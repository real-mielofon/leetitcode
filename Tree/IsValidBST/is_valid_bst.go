package main

import (
	"fmt"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("[l:%s v:%d: r:%s]", t.Left, t.Val, t.Right)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isValidBSTAndLess(root *TreeNode) (result bool, min int, max int) {
	max = root.Val
	min = root.Val
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false, 0, 0
		}
		if res, minLeft, maxLeft := isValidBSTAndLess(root.Left); res {
			if maxLeft >= root.Val {
				return false, 0, 0
			}
			min = MinInt(min, minLeft)
			max = MaxInt(max, maxLeft)
		} else {
			return false, 0, 0
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false, 0, 0
		}
		if res, minRight, maxRight := isValidBSTAndLess(root.Right); res {
			if minRight <= root.Val {
				return false, 0, 0
			}
			min = MinInt(min, minRight)
			max = MaxInt(max, maxRight)
		} else {
			return false, 0, 0
		}
	}
	return true, min, max
}

func isValidBST(root *TreeNode) bool {
	res, _, _ := isValidBSTAndLess(root)
	return res
}

func NewTree(treeArr []interface{}) *TreeNode {
	if len(treeArr) == 0 || treeArr[0] == nil {
		return nil
	}
	var tree TreeNode
	treeNodes := make([]*TreeNode, len(treeArr))
	treeNodes[0] = &tree

	first := 0
	countNodes := 1
	level := 0
	for {
		if first >= len(treeArr) {
			break // за конец прыгнули
		}
		last := first + countNodes - 1
		for j := first; j <= last; j++ {
			if j >= len(treeArr) {
				break // за конец прыгнули
			}
			if val, ok := treeArr[j].(int); ok {
				treeNodes[j].Val = val
				left, right := last+(j-first)*2+1, last+(j-first)*2+2
				if left < len(treeArr) {
					if _, ok := treeArr[left].(int); ok {
						treeNodes[left] = &TreeNode{}
						treeNodes[j].Left = treeNodes[left]
					}
				}
				if right < len(treeArr) {
					if _, ok := treeArr[right].(int); ok {
						treeNodes[right] = &TreeNode{}
						treeNodes[j].Right = treeNodes[right]
					}
				}
			}
		}
		first = last + 1
		countNodes *= 2
		level++
	}

	return &tree
}

/*
				10
		7				12
	6		8		11		13
*/
func main() {
	//treeArr := []interface{}{5, 4, 6, nil, nil, 3, 7}
	//treeArr := []interface{}{2, 1, 3}
	//treeArr := []interface{}{3, 1, 5, 0, 2, 4, 6, nil, nil, nil, 3}
	treeArr := []interface{}{10, 7, 12, 6, 8, 11, 13}

	tree := NewTree(treeArr)

	fmt.Printf("%v %v\n", treeArr, tree)

	fmt.Println(isValidBST(tree))
}
