package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return depth(root) != -1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}
