package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	root *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := []pair{{root: root, sum: root.Val}}
	for len(stack) > 0 {
		node, nodeSum := stack[len(stack)-1].root, stack[len(stack)-1].sum
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil && nodeSum == sum {
			return true
		}
		if node.Right != nil {
			stack = append(stack, pair{node.Right, nodeSum + node.Right.Val})
		}
		if node.Left != nil {
			stack = append(stack, pair{node.Left, nodeSum + node.Left.Val})
		}
	}
	return false
}
