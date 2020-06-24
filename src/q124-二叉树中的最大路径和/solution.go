package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		tmp := node.Val + left + right
		maxSum = max(maxSum, tmp)
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
