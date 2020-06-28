package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	res := root

	if root == nil {
		return root
	}

	curr := root
	pre :=curr
	for curr.Val != key {
		pre = curr
		if curr.Val < key {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

}
