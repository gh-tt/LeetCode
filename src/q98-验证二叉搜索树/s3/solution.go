package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int) bool
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

/*func dfs(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
}*/
