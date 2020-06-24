package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	curr := root
	for curr != nil {
		if curr.Val > val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				break
			} else {
				curr = curr.Left
			}
		} else {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				break
			} else {
				curr = curr.Right
			}
		}
	}
	return root
}
